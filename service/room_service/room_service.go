package room_service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
	"hostelManagementSystem/repositories"
	"time"
)
var roomCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"room")


type RoomService interface {
	SaveRoom(saveRoomDto dtos.SaveRoomDto)(models.Room,error)
	DeleteRoomByRoomId(roomId string)error
	FindRoomById(roomId string)(models.Room,error)
	FindRoomByRoomName(roomName string)(models.Room,error)
	GetNumberOfAvailableBedsByRoomName(roomName string)uint64
}

type RoomServiceImpl struct {

}

func (roomService RoomServiceImpl) SaveRoom(saveRoomDto dtos.SaveRoomDto)(models.Room,error) {
	err := validateDtoInputs(saveRoomDto)

	if err != nil{
		return models.Room{},errors.New(err.Error())
	}

	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)

	room := models.Room{}

	defer cancel()

	err =roomCollection.FindOne(ctx,bson.M{"roomname":saveRoomDto.RoomName}).Decode(&room)

	if err != nil{
		return models.Room{}, errors.New(err.Error())
	}

	if saveRoomDto.RoomName == room.RoomName {
		return models.Room{},errors.New("Room has already been created ")
	} else {
		mapDtoToRoom(saveRoomDto,room)
		room.DateCreated, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		room.DateUpdated, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		_,err=roomCollection.InsertOne(ctx,room)
		if err != nil {
			return models.Room{},errors.New(err.Error())
		}
		return room,nil
	}
}

func validateDtoInputs(dto dtos.SaveRoomDto)error{
	if dto.RoomName == ""{
		return errors.New("Please enter room name ")
	} else
	if dto.RoomNumber <= 0 {
		return errors.New("Please enter Room number ")
	}else
	if dto.NumberOfAvailableBeds <0{
		return errors.New("Please enter number of available beds ")
	}
	return nil
}

func mapDtoToRoom(dto dtos.SaveRoomDto, room models.Room)  {
	room.RoomName = dto.RoomName
	room.RoomNumber = dto.RoomNumber
	room.NumberOfAvailableBeds = dto.NumberOfAvailableBeds
}