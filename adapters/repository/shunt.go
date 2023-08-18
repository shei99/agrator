package repository

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/shei99/agrator/domain/model"
)

type ShuntRepository struct {
	client       influxdb2.Client
	organisation string
	bucket       string
}

func NewShuntRepository(
	connectionUrl string,
	token string,
	organisation string,
	bucket string,
) *ShuntRepository {
	return &ShuntRepository{
		client:       influxdb2.NewClient(connectionUrl, token),
		organisation: organisation,
		bucket:       bucket,
	}
}

func (repository *ShuntRepository) SaveShunt(shunt model.ShuntValue) {
	p := influxdb2.NewPointWithMeasurement("shunt").
		AddTag("batriumId", shunt.BatriumId.Id).
		AddField("shuntVolt", shunt.ShuntVoltage).
		AddField("shuntTemp", shunt.ShuntTemp)
	repository.save(p)
}

func (repository *ShuntRepository) save(point *write.Point) {
	repository.client.WriteAPI(repository.organisation, repository.bucket).
		WritePoint(point)
	// repository.client.Close()
}
