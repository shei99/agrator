package outgoing

import "github.com/shei99/agrator/domain/model"

type WindowRepository interface {
	SaveWindow(
		windowType string,
		measurementType string,
		batriumId model.BatriumIdentifier,
		window model.Window,
	)
	GetLatestBatriumState() map[model.BatriumIdentifier]model.BatriumState
}
