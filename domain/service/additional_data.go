package service

import "github.com/shei99/agrator/domain/model"

type AdditionalDataService interface {
	AddAdditionalData(additionalData model.AdditionalData)
}
