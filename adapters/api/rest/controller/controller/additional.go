package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shei99/agrator/adapters/api/rest/dto"
	"github.com/shei99/agrator/domain/service"
)

type AdditionalDataController struct {
	service service.AdditionalDataService
}

func NewAdditionalDataController(
	batriumService service.AdditionalDataService,
) *AdditionalDataController {
	return &AdditionalDataController{
		service: batriumService,
	}
}

func (controller *AdditionalDataController) PostAdditional(c *gin.Context) {
	var additionalDataDto dto.AdditionalData

	if err := c.BindJSON(&additionalDataDto); err != nil {
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	fmt.Println(additionalDataDto)

	batriumId := model.NewAdditionalDataIdentifier(
		additionalDataDto.HubId,
		additionalDataDto.SystemId,
	)

	additionalData := model.NewAdditionalData()
	additionalData.AdditionalDataId = batriumId
	additionalData.Type = additionalDataDto.Type
	additionalData.Data = additionalDataDto.Data

	controller.service.AddAdditionalData(*additionalData)
}
