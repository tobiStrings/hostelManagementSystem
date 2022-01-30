package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
	"hostelManagementSystem/service/mocks"
	"testing"
)

var student models.Student

var studentDto dtos.StudentDto

var studentService StudentServiceImp

var studentErrorMock *mocks.StudentErrorTest

var studentMock *mocks.StudentMock
func init()  {
	studentDto = dtos.NewStudentDto("Titobiloluwa","Ligali",23,"154760")
	studentMock = new(mocks.StudentMock)
	studentErrorMock = new(mocks.StudentErrorTest)
	studentErrorMock.Err = errors.New("")
}



func TestMapDtoToStudent(t *testing.T) {

	studentDto  = dtos.NewStudentDto("Titobiloluwa","Ligali",23,"154760")

	studentMapped :=MapDtoToStudent(studentDto,student)

	assert.NotNil(t, studentMapped,"student not nil")

	assert.Equal(t, studentMapped.FirstName,"Titobiloluwa")
}

func TestSaveStudent(t *testing.T)  {

	studentMock.On("SaveStudent",studentDto).Return(student,nil)


	returnedStudent, _ := studentMock.SaveStudent(studentDto)


	studentService = StudentServiceImp{studentMock}

	assert.NotNil(t, returnedStudent)
}

func TestSaveWhenStructValidationThrowsAnError(t *testing.T){

	studentErrorMock.On("SaveStudent",studentDto).Return(student,studentErrorMock.Err)

	studentService = StudentServiceImp{studentErrorMock}

	_, err := studentErrorMock.SaveStudent(studentDto)

	assert.NotNil(t, err)
}

func TestThatStudentSavedCanBeFoundWithAnId(t *testing.T)  {
	studentMock.On("")
}