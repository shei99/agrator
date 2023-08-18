package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shei99/agrator/adapters/api/rest/dto"
	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/domain/service"
)

type ShuntController struct {
	service service.ShuntService
}

func NewShuntController(batriumService service.ShuntService) *ShuntController {
	return &ShuntController{
		service: batriumService,
	}
}

func (controller *ShuntController) PostShuntCurrent(c *gin.Context) {
	var shuntDto dto.CurrentShunt

	if err := c.BindJSON(&shuntDto); err != nil {
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	// fmt.Println(shuntDto)

	shunt := model.NewShuntValue()
	shunt.BatriumId = model.NewBatriumIdentifier(shuntDto.HubId, shuntDto.SystemId)

	shunt.ShuntVoltage = shuntDto.ShuntVoltage
	shunt.ShuntTemp = shuntDto.ShuntTemp

	controller.service.AnalyseShuntData(*shunt)
}
