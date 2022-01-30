package routes

import (
	"github.com/gin-gonic/gin"
	"hostelManagementSystem/controller"
)

func StudentRoute(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("/student",controller.SaveStudent())
	incomingRoutes.GET("findStudent/:id",controller.GetStudentById())
}