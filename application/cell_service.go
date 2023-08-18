package application

import (
	"context"

	"github.com/reactivex/rxgo/v2"

	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/domain/service"
	"github.com/shei99/agrator/ports/outgoing"
)

type CellService struct {
	windowService  service.WindowService
	cellRepository outgoing.CellRepository

	cellnodeChannel  chan rxgo.Item
	cellAnalggration service.Analggration
}

func NewCellService(
	windowService service.WindowService,
	cellRepository outgoing.CellRepository,
	cellAnalggration service.Analggration,
) *CellService {
	cellService := &CellService{
		windowService:    windowService,
		cellRepository:   cellRepository,
		cellAnalggration: cellAnalggration,
	}

	cellService.Setup()
	return cellService
}

func (service *CellService) Setup() {
	service.initCellEvaluation()
}

func (service *CellService) initCellEvaluation() {
	service.cellnodeChannel = make(chan rxgo.Item)
	observable := rxgo.FromChannel(service.cellnodeChannel, rxgo.WithPublishStrategy())
	service.cellAnalggration.InnerWindowEvaluation(observable)
	service.cellAnalggration.CriticalWindowEvaluation(observable)
	observable.Connect(context.Background())
}

func (service *CellService) AnalyseCellData(
	cellnode model.Cellnode,
) {
	service.produceCellnode(cellnode)
}

func (service *CellService) produceCellnode(cellnode model.Cellnode) {
	service.cellnodeChannel <- rxgo.Item{V: service.enrichCellnodeWithViolation(cellnode)}
}

func (service *CellService) enrichCellnodeWithViolation(
	cellnode model.Cellnode,
) model.CellnodeWithViolation {
	enrichedCellnode := model.NewCellNodeWithViolation(cellnode)

	batriumState := service.windowService.State(cellnode.BatriumId)
	enrichedCellnode.InnerWindowViolation = service.checkCellnodeForViolation(
		cellnode,
		batriumState.CellInnerWindow,
	)
	enrichedCellnode.CriticalWindowViolation = service.checkCellnodeForViolation(
		cellnode,
		batriumState.CellCriticalWindow,
	)

	return *enrichedCellnode
}

func (service *CellService) checkCellnodeForViolation(
	cellnode model.Cellnode,
	window model.Window,
) bool {
	if window.CellTempHi == 0 && window.CellTempLo == 0 && window.CellVoltHi == 0 &&
		window.CellVoltLo == 0 {
		return false
	}

	for _, cell := range cellnode.Cells {
		if cell.CellTemp < window.CellTempLo || cell.CellTemp > window.CellTempHi {
			return true
		}
		if cell.CellVolt < window.CellVoltLo || cell.CellVolt > window.CellVoltHi {
			return true
		}
	}
	return false
}
