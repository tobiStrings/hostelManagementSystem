package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/repositories"
	"hostelManagementSystem/service"
	"log"
	"net/http"
)


var studentCollection *mongo.Collection = repositories.OpenCollection(repositories.Client,"student")
var validate = validator.New()
var studentService = new(service.StudentServiceImp)

func SaveStudent() gin.HandlerFunc{
	return func(c *gin.Context) {
		var studentDto dtos.StudentDto

		err := c.BindJSON(&studentDto)

		validatedStruct := validate.Struct(studentDto)

		if validatedStruct != nil{
			log.Fatalln("error validating dto"+ validatedStruct.Error())
		}

		newStudent,err :=studentService.SaveStudent(studentDto)
		var msg string
		if err != nil{
			msg = err.Error()
		}

		if msg =="user not created because validation cannot be done"{
			c.JSON(http.StatusBadRequest,gin.H{"error":msg})
			return
		}else
		if msg == "user not created because student cannot be inserted"{
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}else
		if msg == "student already exist"{
			c.JSON(http.StatusBadRequest,gin.H{"error":msg})
			return
		}
		if msg == "mongo: no documents in result"{
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}
		c.JSON(http.StatusOK,newStudent)
	}
}

func GetStudentById() gin.HandlerFunc{
	return func(c *gin.Context) {

		studentId := c.Param("id")

		student, err := studentService.GetStudentById(studentId)

		if err != nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})

			return
		}

		c.JSON(http.StatusOK,student)
	}
}
