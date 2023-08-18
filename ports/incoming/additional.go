package incoming

import (
	"github.com/gin-gonic/gin"
	"github.com/shei99/agrator/adapters/api/rest/controller/controller"
)

func NewAdditionalDataRouter(router *gin.Engine, controller controller.AdditionalDataController) {
	group := router.Group("/batrium")

	group.POST("/additionalData", controller.PostAdditional)
}
