package registry

import (
	"github.com/gin-gonic/gin"
	"wan-api-kol-event/Controllers"
)

func RegisterRoutes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")
	{
		// base API
		apiV1.GET("/kols", Controllers.GetKolsController)
		// insert Dummy data
		apiV1.POST("/kols/insert-dummy-data", Controllers.InsertKolsController)
	}
}
