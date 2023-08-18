package incoming

import (
	"github.com/gin-gonic/gin"
	"github.com/shei99/agrator/adapters/api/rest/controller/controller"
)

func NewShuntRouter(router *gin.Engine, controller controller.ShuntController) {
	group := router.Group("/batrium")

	group.POST("/shunt/current", controller.PostShuntCurrent)
}
