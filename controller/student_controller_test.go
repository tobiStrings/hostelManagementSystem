package controller

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"hostelManagementSystem/models"
	"testing"
)

type Mock struct {
	mock.Mock
}

func TestSaveStudent(t *testing.T) {
	studentToBeReturned := new(models.Student)
	var mockTest Mock
	mockTest.On("SaveStudent",nil).Return(studentToBeReturned)
	mockTest.Called(1,true,"SaveStudent")

	assert.Equal(t, studentToBeReturned,new(models.Student))

}