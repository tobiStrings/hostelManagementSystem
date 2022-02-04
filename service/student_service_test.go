package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
	"hostelManagementSystem/service/mocks"
	"hostelManagementSystem/service/student_service"
	"testing"
)

var student models.Student

var studentDto dtos.StudentDto

var studentService student_service.StudentServiceImp

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

	studentMapped := student_service.MapDtoToStudent(studentDto,student)

	assert.NotNil(t, studentMapped,"student not nil")

	assert.Equal(t, studentMapped.FirstName,"Titobiloluwa")
}

func TestSaveStudent(t *testing.T)  {

	studentMock.On("SaveStudent",studentDto).Return(student,nil)


	returnedStudent, _ := studentMock.SaveStudent(studentDto)


	studentService = student_service.StudentServiceImp{}

	assert.NotNil(t, returnedStudent)
}

func TestSaveWhenStructValidationThrowsAnError(t *testing.T){

	studentErrorMock.On("SaveStudent",studentDto).Return(student,studentErrorMock.Err)

	studentService = student_service.StudentServiceImp{}

	_, err := studentErrorMock.SaveStudent(studentDto)

	assert.NotNil(t, err)
}
func TestSaveStudentWhenStudentHasAlreadyBeenSaved(t *testing.T)  {
	studentMock.On("SaveStudent",studentDto).Return(student,nil)


	returnedStudent, _ := studentMock.SaveStudent(studentDto)


	studentService = student_service.StudentServiceImp{}

	assert.NotNil(t, returnedStudent)

	//Saving the same student the second time

	studentErrorMock.On("SaveStudent",studentDto).Return(models.Student{},studentErrorMock.Err)

	studentService = student_service.StudentServiceImp{}

	_, err := studentErrorMock.SaveStudent(studentDto)

	assert.NotNil(t, err)
}
func TestThatStudentSavedCanBeFoundWithAnId(t *testing.T)  {

	studentMock.On("GetStudentById","1236674").Return(student,nil)

	studentService = student_service.StudentServiceImp{}

	student,_ := studentMock.GetStudentById("1236674")

	assert.NotNil(t, student)

	assert.Equal(t, student,student)
}

func TestThatNonExistingStudentCannotBeFound(t *testing.T)  {
	emptyStudent := models.Student{}

	studentErrorMock.On("GetStudentById","rr555353").Return(emptyStudent,studentErrorMock.Err)

	studentService = student_service.StudentServiceImp{}

	_, err := studentErrorMock.GetStudentById("rr555353")

	assert.NotNil(t, err)

	assert.Error(t, studentErrorMock.Err,"")
}