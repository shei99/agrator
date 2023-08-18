package outgoing

import "github.com/shei99/agrator/domain/model"

type ShuntRepository interface {
	SaveCellnode(cellnode model.Cellnode)
	SaveShunt(shunt model.ShuntValue)
	SaveAdditional(additionalData model.AdditionalData)
	SaveWindow(
		windowType string,
		measurementType string,
		batriumId model.BatriumIdentifier,
		window model.Window,
	)
	GetLatestBatriumState() map[model.BatriumIdentifier]model.BatriumState
}
