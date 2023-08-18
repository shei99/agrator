package service

import "github.com/shei99/agrator/domain/model"

type WindowService interface {
	State(batriumId model.BatriumIdentifier) model.BatriumState
	UpdateCellInnerWindow(batriumId model.BatriumIdentifier, window model.Window)
	UpdateCellCritical(batriumId model.BatriumIdentifier, window model.Window)
	UpdateShuntInnerWindow(batriumId model.BatriumIdentifier, window model.Window)
	UpdateShuntCritical(batriumId model.BatriumIdentifier, window model.Window)
}
