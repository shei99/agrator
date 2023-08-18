package repository

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/shei99/agrator/domain/model"
)

type AdditionalDataRepository struct {
	client       influxdb2.Client
	organisation string
	bucket       string
}

func NewAdditionalDataRepository(
	connectionUrl string,
	token string,
	organisation string,
	bucket string,
) *AdditionalDataRepository {
	return &AdditionalDataRepository{
		client:       influxdb2.NewClient(connectionUrl, token),
		organisation: organisation,
		bucket:       bucket,
	}
}

func (repository *AdditionalDataRepository) SaveAdditional(additionalData model.AdditionalData) {
	p := influxdb2.NewPointWithMeasurement(additionalData.Type).
		AddTag("batriumId", additionalData.BatriumId.Id)
	for k, v := range additionalData.Data {
		p.AddField(k, v)
	}

	repository.save(p)
}

func (repository *AdditionalDataRepository) save(point *write.Point) {
	repository.client.WriteAPI(repository.organisation, repository.bucket).
		WritePoint(point)
	// repository.client.Close()
}
