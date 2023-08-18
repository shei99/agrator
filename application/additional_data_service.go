package application

import (
	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/ports/outgoing"
)

type AdditionalDataService struct {
	additionalDataRepository outgoing.AdditionalDataRepository
}

func NewAdditionalDataService(
	additionalDataRepository outgoing.AdditionalDataRepository,
) *AdditionalDataService {
	return &AdditionalDataService{
		additionalDataRepository: additionalDataRepository,
	}
}

func (service *AdditionalDataService) AddAdditionalData(additionalData model.AdditionalData) {
	service.additionalDataRepository.SaveAdditional(additionalData)
}
