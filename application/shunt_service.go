package application

import (
	"context"

	"github.com/reactivex/rxgo/v2"

	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/domain/service"
	"github.com/shei99/agrator/ports/outgoing"
)

type ShuntService struct {
	windowService     service.WindowService
	batriumRepository outgoing.ShuntRepository
	shuntChannel      chan rxgo.Item
	shuntAnalggration service.Analggration
}

func NewShuntService(
	windowService service.WindowService,
	batriumRepo outgoing.ShuntRepository,
	shuntAnalggration service.Analggration,
) *ShuntService {
	shuntService := &ShuntService{
		windowService:     windowService,
		batriumRepository: batriumRepo,
		shuntAnalggration: shuntAnalggration,
	}

	shuntService.Setup()
	return shuntService
}

func (service *ShuntService) Setup() {
	service.initShuntEvaluation()
}

func (service *ShuntService) initShuntEvaluation() {
	service.shuntChannel = make(chan rxgo.Item)
	observable := rxgo.FromChannel(service.shuntChannel, rxgo.WithPublishStrategy())
	service.shuntAnalggration.InnerWindowEvaluation(observable)
	service.shuntAnalggration.CriticalWindowEvaluation(observable)
	observable.Connect(context.Background())
}

func (service *ShuntService) AnalyseShuntData(
	shunt model.ShuntValue,
) {
	service.produceShunt(shunt)
}

func (service *ShuntService) produceShunt(shunt model.ShuntValue) {
	service.shuntChannel <- rxgo.Item{V: service.enrichShuntWithViolation(shunt)}
}

func (service *ShuntService) enrichShuntWithViolation(
	shunt model.ShuntValue,
) model.ShuntWithViolation {
	enrichedShunt := model.NewShuntWithViolation(shunt)

	batriumState := service.windowService.State(shunt.BatriumId)
	enrichedShunt.InnerWindowViolation = service.checkShuntForViolation(
		shunt,
		batriumState.ShuntInnerWindow,
	)
	enrichedShunt.CriticalWindowViolation = service.checkShuntForViolation(
		shunt,
		batriumState.ShuntCriticalWindow,
	)

	return *enrichedShunt
}

func (service *ShuntService) checkShuntForViolation(
	shunt model.ShuntValue,
	window model.Window,
) bool {
	if window.CellTempHi == 0 && window.CellTempLo == 0 && window.CellVoltHi == 0 &&
		window.CellVoltLo == 0 {
		return false
	}

	if shunt.ShuntTemp < window.CellTempLo || shunt.ShuntTemp > window.CellTempHi {
		return true
	}
	if shunt.ShuntVoltage < window.CellVoltLo || shunt.ShuntVoltage > window.CellVoltHi {
		return true
	}
	return false
}
