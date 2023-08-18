package service

import "github.com/shei99/agrator/domain/model"

type ShuntService interface {
	AnalyseShuntData(shunt model.ShuntValue)
}
