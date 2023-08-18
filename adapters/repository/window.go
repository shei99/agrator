package repository

import (
	"context"
	"encoding/json"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	"github.com/shei99/agrator/domain/model"
)

type WindowRepository struct {
	client       influxdb2.Client
	organisation string
	bucket       string
}

func NewWindowRepository(
	connectionUrl string,
	token string,
	organisation string,
	bucket string,
) *WindowRepository {
	return &WindowRepository{
		client:       influxdb2.NewClient(connectionUrl, token),
		organisation: organisation,
		bucket:       bucket,
	}
}

func (repository *WindowRepository) SaveWindow(
	windowType string, measurementType string,
	batriumId model.BatriumIdentifier,
	window model.Window,
) {
	p := influxdb2.NewPointWithMeasurement(measurementType+"-window").
		AddTag("batriumId", batriumId.Id).
		AddTag("type", windowType).
		AddField("voltLo", window.CellVoltLo).
		AddField("voltHi", window.CellVoltHi).
		AddField("tempLo", window.CellTempLo).
		AddField("tempHi", window.CellTempHi)

	repository.save(p)
}

func (repository *WindowRepository) GetLatestBatriumState() map[model.BatriumIdentifier]model.BatriumState {
	batriumIds := repository.getAllBatriumIds()
	fmt.Println(batriumIds)
	batriumState := make(map[model.BatriumIdentifier]model.BatriumState)
	for _, batriumId := range batriumIds {
		fmt.Println(batriumId)
		singleBatriumState, ok := batriumState[batriumId]

		if !ok {
			batriumState[batriumId] = *model.NewBatriumState()
		}

		singleBatriumState.CellInnerWindow = repository.getLatestBatriumWindow(
			batriumId,
			"inner",
			"cell",
		)
		singleBatriumState.CellCriticalWindow = repository.getLatestBatriumWindow(
			batriumId,
			"critical",
			"cell",
		)
		singleBatriumState.ShuntInnerWindow = repository.getLatestBatriumWindow(
			batriumId,
			"inner",
			"shunt",
		)
		singleBatriumState.ShuntCriticalWindow = repository.getLatestBatriumWindow(
			batriumId,
			"critical",
			"shunt",
		)
		batriumState[batriumId] = singleBatriumState
	}

	return batriumState
}

func (repository *WindowRepository) getAllBatriumIds() []model.BatriumIdentifier {
	queryAPI := repository.client.QueryAPI(repository.organisation)

	query := fmt.Sprintf(
		`from(bucket:"%s")|> range(start: -4d)|> filter(fn: (r) => r["_measurement"] == "cell-window")
  |> filter(fn: (r) => r["type"] == "inner")
  |> last()`,
		repository.bucket,
	)
	// get QueryTableResult
	result, err := queryAPI.Query(
		context.Background(),
		query,
	)
	batriumIds := []model.BatriumIdentifier{}
	if err == nil {
		// Iterate over query response
		for result.Next() {
			batriumId := model.BatriumIdentifier{
				Id: fmt.Sprint(result.Record().ValueByKey("batriumId")),
			}
			if len(batriumIds) == 0 || batriumIds[len(batriumIds)-1] != batriumId {
				batriumIds = append(batriumIds, batriumId)
			}
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	repository.client.Close()
	return batriumIds
}

func (repository *WindowRepository) getLatestBatriumWindow(
	batriumId model.BatriumIdentifier,
	windowType string,
	measurementType string,
) model.Window {
	queryAPI := repository.client.QueryAPI(repository.organisation)

	query := fmt.Sprintf(
		`from(bucket:"%s")|> range(start: -4d)|> filter(fn: (r) => r["_measurement"] == "%s-window")
  |> filter(fn: (r) => r["type"] == "%s")
	|> filter(fn: (r) => r["batriumId"] == "%s")
  |> last()`,
		repository.bucket,
		measurementType,
		windowType,
		batriumId.Id,
	)
	// get QueryTableResult
	result, err := queryAPI.Query(
		context.Background(),
		query,
	)

	windowData := make(map[string]interface{})
	if err == nil {
		// Iterate over query response
		for result.Next() {
			if windowData == nil {
				windowData = make(map[string]interface{})
			}

			windowData[result.Record().Field()] = result.Record().Value()
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}
	// Ensures background processes finishes
	repository.client.Close()

	var window model.Window
	jsonbody, err := json.Marshal(windowData)
	if err != nil {
		// do error check
		fmt.Println(err)
		// return
	}

	if err := json.Unmarshal(jsonbody, &window); err != nil {
		// do error check
		fmt.Println(err)
		// return
	}

	fmt.Println(window)
	return window
}

func (repository *WindowRepository) save(point *write.Point) {
	repository.client.WriteAPI(repository.organisation, repository.bucket).
		WritePoint(point)
	// repository.client.Close()
}
