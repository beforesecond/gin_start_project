package routers

import (
	"test-gin/controllers"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("/home", controllers.MainHome)
}
