package controller

import (
	"github.com/gin-gonic/gin"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/service/room_service"
	"log"
	"net/http"
)

//var validate = validator.New()
var roomService = new(room_service.RoomServiceImpl)

func CreateRoom() gin.HandlerFunc {
	return func(c *gin.Context) {
		var saveRoomDto dtos.SaveRoomDto

		err:=c.BindJSON(&saveRoomDto)

		validateStruct := validate.Struct(saveRoomDto)

		if validateStruct != nil{
			log.Fatalln("error validating dto"+ validateStruct.Error())
		}

		room,err := roomService.SaveRoom(saveRoomDto)

		msg := err.Error()
		if err != nil{
			if msg == "Please enter room name "{
				c.JSON(http.StatusBadRequest,gin.H{"error":msg})
				return
			}

			if msg == "Please enter Room number "{
				c.JSON(http.StatusBadRequest,gin.H{"error":msg})
				return
			}

			if msg == "Please enter number of available beds "{
				c.JSON(http.StatusBadRequest,gin.H{"error":msg})
				return
			}

			if msg =="mongo: no documents in result"{
				c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
				return
			}

			if msg == "Room has already been created "{
				c.JSON(http.StatusBadRequest,gin.H{"error":msg})
				return
			}
		}
		c.JSON(http.StatusOK,room)
	}
}
