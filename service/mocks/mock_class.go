package mocks

import (
	"github.com/stretchr/testify/mock"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
)

type StudentMock struct {
	mock.Mock

}
var student models.Student
func (s *StudentMock) SaveStudent(dto dtos.StudentDto) (models.Student, error) {
	args :=s.Called(dto)

	studentToReturn := models.Student{}

	var ret = args.Error(1)

	return  studentToReturn,ret
}

func (s *StudentMock) GetStudentById(studentId string) (models.Student, error) {
	args := s.Called(studentId)

	studentToReturn := models.Student{}

	return studentToReturn,args.Error(1)
}

type StudentErrorTest struct {
	mock.Mock
	Err error

}


func (s *StudentErrorTest) SaveStudent(dto dtos.StudentDto) (models.Student, error) {
	//args :=s.Called(dto)

	studentToReturn := models.Student{}

	return  studentToReturn, s.Err
}

func (s *StudentErrorTest) GetStudentById(studentId string) (models.Student, error) {

	studentToReturn := models.Student{}

	return studentToReturn,s.Err
}
