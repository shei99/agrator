package repository

import (
	"strconv"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/shei99/agrator/domain/model"
)

type CellRepository struct {
	client       influxdb2.Client
	organisation string
	bucket       string
}

func NewCellRepository(
	connectionUrl string,
	token string,
	organisation string,
	bucket string,
) *CellRepository {
	return &CellRepository{
		client:       influxdb2.NewClient(connectionUrl, token),
		organisation: organisation,
		bucket:       bucket,
	}
}

func (repository *CellRepository) SaveCellnode(cellnode model.Cellnode) {
	for _, cell := range cellnode.Cells {
		p := influxdb2.NewPointWithMeasurement("cells").
			AddTag("batriumId", cellnode.BatriumId.Id).
			AddTag("cellId", strconv.Itoa(int(cell.Id))).
			AddField("cellVolt", cell.MinCellVolt).
			// AddField("maxCellVolt", cell.MaxCellVolt).
			AddField("cellTemp", cell.CellTemp).
			AddField("status", int(cell.Status))
		// SetTime(time.Now())
		repository.save(p)
	}
}

func (repository *CellRepository) save(point *write.Point) {
	repository.client.WriteAPI(repository.organisation, repository.bucket).
		WritePoint(point)
	// repository.client.Close()
}
