package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
	"hostelManagementSystem/repositories"
	"time"
)

func MapDtoToStudent(dto dtos.StudentDto,student models.Student) models.Student{
	student.FirstName = dto.FirstName
	student.LastName = dto.LastName
	student.Age = dto.Age
	student.MatricNumber = dto.MaricNumber
	return student
}

type StudentService interface {
	SaveStudent(dto dtos.StudentDto) (models.Student,error)
	GetStudentById(studentId string) (models.Student,error)
}

type StudentServiceImp struct {
	studentService StudentService
}

var validate = validator.New()

var studentCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"student")


func (studentService *StudentServiceImp) SaveStudent(dto dtos.StudentDto) (models.Student, error){
	var ctx,cancel =context.WithTimeout(context.Background(),100*time.Second)

	foundStudent := models.Student{}

	matric_number :=dto.MaricNumber

	defer cancel()
		st := studentCollection.FindOne(ctx, bson.M{"matric_number": matric_number}).Decode(&foundStudent)
		if st != nil{
			fmt.Println("Error")
			fmt.Println(st.Error())
			return models.Student{}, errors.New(st.Error())
		}
		if &foundStudent != nil{
			return foundStudent,errors.New("user already exist")
		}


	var newStudent models.Student

	newStudent = MapDtoToStudent(dto,newStudent)


	validationErr := validate.Struct(newStudent)

	if validationErr != nil{
		return models.Student{},errors.New("user not created because validation cannot be done")
	}

	newStudent.DateCreated,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	newStudent.DateUpdated,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
	newStudent.Id =primitive.NewObjectID()
	newStudent.StudentId = newStudent.Id.Hex()

	_,insertError := studentCollection.InsertOne(ctx,newStudent)

	if insertError != nil {
		return models.Student{},errors.New("user not created because student cannot be inserted")
	}

	defer cancel()

	return newStudent,nil
}

func (studentService *StudentServiceImp) GetStudentById(studentId string) (models.Student,error) {
	var ctx,cancel =context.WithTimeout(context.Background(),10*time.Second)

	var studentToBeReturned models.Student

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(studentId)

	err := studentCollection.FindOne(ctx,bson.M{"id": objId}).Decode(&studentToBeReturned)

	if err != nil{
		return models.Student{},errors.New("student not found")
	}

	return studentToBeReturned,nil
}