package routers

import (
	"project/app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	var ctrl controllers.UserController
	v1 := route.Group("/api/user")
	v1.POST("register/", ctrl.Register)
}
