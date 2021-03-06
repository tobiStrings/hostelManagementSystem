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

func (roomService *RoomServiceImpl) SaveRoom(saveRoomDto dtos.SaveRoomDto)(models.Room,error) {
	err := validateDtoInputs(saveRoomDto)

	if err != nil{
		return models.Room{},errors.New(err.Error())
	}

	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)

	room := models.Room{}

	defer cancel()

	err =roomCollection.FindOne(ctx,bson.M{"roomname":saveRoomDto.RoomName}).Decode(&room)
	if err != nil{
		if saveRoomDto.RoomName == room.RoomName {
			return models.Room{},errors.New("Room has already been created ")
		}

		room = mapDtoToRoom(saveRoomDto,room)
		room.DateCreated, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		room.DateUpdated, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		room.Beds = make([]models.Bed,saveRoomDto.NumberOfAvailableBeds)
		_,err=roomCollection.InsertOne(ctx,room)
		if err != nil {
			return models.Room{},errors.New(err.Error())
		}
		return room,nil
	}

	return models.Room{}, nil
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

func mapDtoToRoom(dto dtos.SaveRoomDto, room models.Room) models.Room {
	room.RoomName = dto.RoomName
	room.RoomNumber = dto.RoomNumber
	room.NumberOfAvailableBeds = dto.NumberOfAvailableBeds
	return room
}

func (roomServiceImpl *RoomServiceImpl)DeleteRoomByRoomId(roomId string)error{
	return nil
}


func (roomServiceImpl *RoomServiceImpl)FindRoomById(roomId string)(models.Room,error){
	return models.Room{},nil
}

func (roomServiceImpl *RoomServiceImpl)FindRoomByRoomName(roomName string)(models.Room,error){
	return models.Room{},nil
}

func (roomServiceImpl *RoomServiceImpl)GetNumberOfAvailableBedsByRoomName(roomName string)uint64{
	return 0
}