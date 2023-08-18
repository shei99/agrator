package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shei99/agrator/adapters/api/rest/dto"
	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/domain/service"
)

type WindowController struct {
	service service.WindowService
}

func NewWindowController(batriumService service.WindowService) *WindowController {
	return &WindowController{
		service: batriumService,
	}
}

func (controller *WindowController) PostCellInnerWindow(c *gin.Context) {
	var cellInnerWindowDto dto.CellInnerWindow

	if err := c.BindJSON(&cellInnerWindowDto); err != nil {
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	// fmt.Println(cellInnerWindowDto)

	window := model.NewWindow()
	window.CellVoltLo = cellInnerWindowDto.ControlDischargeCellVoltLo
	window.CellVoltHi = cellInnerWindowDto.ControlChargeCellVoltHi
	window.CellTempLo = cellInnerWindowDto.ControlChargeCellTempLo
	window.CellTempHi = cellInnerWindowDto.ControlChargeCellTempHi

	batriumId := model.NewBatriumIdentifier(cellInnerWindowDto.HubId, cellInnerWindowDto.SystemId)

	controller.service.UpdateCellInnerWindow(batriumId, *window)
}

func (controller *WindowController) PostShuntInnerWindow(c *gin.Context) {
	var shuntInnerWindowDto dto.ShuntInnerWindow

	if err := c.BindJSON(&shuntInnerWindowDto); err != nil {
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	// fmt.Println(shuntInnerWindowDto)

	window := model.NewWindow()
	window.CellVoltLo = shuntInnerWindowDto.ControlDischargeTargetLimpVolt
	window.CellVoltHi = shuntInnerWindowDto.ControlChargeTargetLimpVolt
	window.CellTempLo = 15
	window.CellTempHi = 40

	batriumId := model.NewBatriumIdentifier(shuntInnerWindowDto.HubId, shuntInnerWindowDto.SystemId)

	controller.service.UpdateShuntInnerWindow(batriumId, *window)
}

func (controller *WindowController) PostCellCritical(c *gin.Context) {
	var cellCriticalWindowDto dto.CellCriticalWindow

	if err := c.BindJSON(&cellCriticalWindowDto); err != nil {
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	// fmt.Println(cellCriticalWindowDto)

	window := model.NewWindow()
	window.CellVoltLo = cellCriticalWindowDto.ControlCriticalCellVoltLo
	window.CellVoltHi = cellCriticalWindowDto.ControlCriticalCellVoltHi
	window.CellTempLo = cellCriticalWindowDto.ControlCriticalCellTempLo
	window.CellTempHi = cellCriticalWindowDto.ControlCriticalCellTempHi

	batriumId := model.NewBatriumIdentifier(
		cellCriticalWindowDto.HubId,
		cellCriticalWindowDto.SystemId,
	)

	controller.service.UpdateCellCritical(batriumId, *window)
}

func (controller *WindowController) PostShuntCritical(c *gin.Context) {
	var shuntCriticalWindowDto dto.ShuntCriticalWindow

	if err := c.BindJSON(&shuntCriticalWindowDto); err != nil {
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	// fmt.Println(shuntCriticalWindowDto)

	window := model.NewWindow()
	window.CellVoltLo = shuntCriticalWindowDto.ControlCriticalShuntVoltLo
	window.CellVoltHi = shuntCriticalWindowDto.ControlCriticalShuntVoltHi
	window.CellTempLo = 10
	window.CellTempHi = 50

	batriumId := model.NewBatriumIdentifier(
		shuntCriticalWindowDto.HubId,
		shuntCriticalWindowDto.SystemId,
	)

	controller.service.UpdateShuntCritical(batriumId, *window)
}
