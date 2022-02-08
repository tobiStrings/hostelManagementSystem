package routes

import (
	"github.com/gin-gonic/gin"
	"hostelManagementSystem/controller"
)

func RoomRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("/room",controller.CreateRoom())
}
