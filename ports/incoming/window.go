package incoming

import (
	"github.com/gin-gonic/gin"
	"github.com/shei99/agrator/adapters/api/rest/controller/controller"
)

func NewWindowRouter(router *gin.Engine, controller controller.WindowController) {
	group := router.Group("/batrium/window")

	group.POST("critical/cell", controller.PostCellCritical)
	group.POST("/inner/cell", controller.PostCellInnerWindow)
	group.POST("/critical/shunt", controller.PostShuntCritical)
	group.POST("/inner/shunt", controller.PostShuntInnerWindow)
}
