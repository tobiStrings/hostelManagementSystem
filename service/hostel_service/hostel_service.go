package hostel_service

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
	"hostelManagementSystem/repositories"
	"hostelManagementSystem/service/student_service"
	"time"
)

var roomCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"room")
//var bedCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"bed")
//var hostelCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"hostel")


type HostelService interface {
	AssignBedsToStudents(dto dtos.GetBedSpaceDto) (models.Bed,error)
}

type HostelServiceImpl struct {
	studentServiceImpl student_service.StudentServiceImp
}

func (hostelServiceIpl *HostelServiceImpl) AssignBedsToStudents(dto dtos.GetBedSpaceDto) (models.Bed,error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	err:=verifyDtoValues(dto)

	room := models.Room{}

	defer cancel()

	msg:=err.Error()

	if err != nil{
		return models.Bed{},errors.New(msg)
	}

	student,errFindingStudent :=hostelServiceIpl.studentServiceImpl.GetStudentByMatricNumber(dto.StudentMatricNumber)

	msg = errFindingStudent.Error()

	if errFindingStudent != nil{
		return models.Bed{},errors.New(msg)
	}

	if &student != nil{

		err := roomCollection.FindOne(ctx,bson.M{"roomname":dto.RoomName}).Decode(&room)

		if err != nil {
			return models.Bed{},errors.New(err.Error())
		}else
		if  &room == nil {
			return models.Bed{},errors.New("room does not exist")
		}
		if room.NumberOfAvailableBeds != 0{

			student.Room = room

			student.Room.NumberOfAvailableBeds -= 1

		} else {return models.Bed{}, errors.New(" Room does not have an empty space")}

		roomBeds := student.Room.Beds

		for i :=0; i<= len(student.Room.Beds); i++{

			if &roomBeds[i] == nil{

				student.Bed = roomBeds[i]

				return student.Bed,nil
			}
		}
	}

	return models.Bed{},errors.New("student cannot be assigned a bed space")
}

func verifyDtoValues(dto dtos.GetBedSpaceDto) error{
	if dto.RoomName == "" {
		return errors.New("room name cannot be empty")
	}else if dto.StudentMatricNumber == "" {
		return errors.New("student name cannot be empty")
	}

	return nil
}