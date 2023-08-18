package application

import (
	"fmt"

	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/ports/outgoing"
)

type WindowService struct {
	repository outgoing.WindowRepository
	state      map[model.BatriumIdentifier]model.BatriumState
}

func NewWindowService(
	repository outgoing.WindowRepository,
) *WindowService {
	batriumAnalyseConfiguration := WindowService{
		repository: repository,
		state:      make(map[model.BatriumIdentifier]model.BatriumState),
	}
	batriumAnalyseConfiguration.initState()

	return &batriumAnalyseConfiguration
}

func (service *WindowService) initState() {
	service.state = service.repository.GetLatestBatriumState()
	fmt.Println(service.state)
}

func (service *WindowService) State(batriumId model.BatriumIdentifier) model.BatriumState {
	return service.state[batriumId]
}

func (service *WindowService) UpdateCellInnerWindow(
	batriumId model.BatriumIdentifier,
	window model.Window,
) {
	batriumState, ok := service.state[batriumId]

	if ok {
		batriumState.CellInnerWindow = window
		service.state[batriumId] = batriumState
	} else {
		service.state[batriumId] = *model.NewBatriumState()
		service.UpdateCellInnerWindow(batriumId, window)
	}

	service.repository.SaveWindow("inner", "cells", batriumId, window)
}

func (service *WindowService) UpdateCellCriticalWindow(
	batriumId model.BatriumIdentifier,
	window model.Window,
) {
	batriumState, ok := service.state[batriumId]

	if ok {
		batriumState.CellCriticalWindow = window
		service.state[batriumId] = batriumState
	} else {
		service.state[batriumId] = *model.NewBatriumState()
		service.UpdateCellCriticalWindow(batriumId, window)
	}

	service.repository.SaveWindow("critical", "cells", batriumId, window)
}

func (service *WindowService) UpdateShuntInnerWindow(
	batriumId model.BatriumIdentifier,
	window model.Window,
) {
	batriumState, ok := service.state[batriumId]

	if ok {
		batriumState.ShuntInnerWindow = window
		service.state[batriumId] = batriumState
	} else {
		service.state[batriumId] = *model.NewBatriumState()
		service.UpdateShuntInnerWindow(batriumId, window)
	}

	service.repository.SaveWindow("inner", "shunt", batriumId, window)
}

func (service *WindowService) UpdateShuntCriticalWindow(
	batriumId model.BatriumIdentifier,
	window model.Window,
) {
	batriumState, ok := service.state[batriumId]

	if ok {
		batriumState.ShuntCriticalWindow = window
		service.state[batriumId] = batriumState
	} else {
		service.state[batriumId] = *model.NewBatriumState()
		service.UpdateShuntCriticalWindow(batriumId, window)
	}

	service.repository.SaveWindow("critical", "shunt", batriumId, window)
}

func (service *WindowService) checkForChange(
	oldWindow model.Window,
	newWindow model.Window,
) bool {
	return oldWindow == newWindow
}

func (service *WindowService) saveChanges(window model.Window) {
	// TODO: save with repository
}
