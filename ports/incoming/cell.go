package incoming

import (
	"github.com/gin-gonic/gin"

	"github.com/shei99/agrator/adapters/api/rest/controller/controller"
)

func NewCellRouter(router *gin.Engine, controller controller.CellController) {
	group := router.Group("/batrium")

	group.POST("/cells/current", controller.PostCellCurrent)
}
