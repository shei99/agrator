package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shei99/agrator/adapters/api/rest/dto"
	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/domain/service"
)

type CellController struct {
	service service.CellService
}

func NewCellController(batriumService service.CellService) *CellController {
	return &CellController{
		service: batriumService,
	}
}

func (controller *CellController) PostCellCurrent(c *gin.Context) {
	var cellnodeDto dto.CurrentCellnode

	if err := c.BindJSON(&cellnodeDto); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "Error occured")
		return
	}

	cellnode := model.NewCellnode()
	cellnode.BatriumId = model.NewBatriumIdentifier(cellnodeDto.HubId, cellnodeDto.SystemId)

	cells := []model.Cell{}
	for _, node := range cellnodeDto.Nodes {
		cell := model.NewCell()
		cell.Id = node.Id
		cell.CellVolt = node.MaxCellVolt
		cell.CellTemp = node.CellTemp
		cell.Status = node.Status
		cells = append(cells, *cell)
	}
	cellnode.Cells = cells

	// fmt.Println(cellnode)

	controller.service.AnalyseCellData(*cellnode)
}
