package service

import "github.com/shei99/agrator/domain/model"

type CellService interface {
	AnalyseCellData(cellnode model.Cellnode)
}
