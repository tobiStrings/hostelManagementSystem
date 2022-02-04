package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)
type Student struct {
	Id  primitive.ObjectID `bson:"id"`
	FirstName string `json:"first_name" binding:"required,min=2,max=30"`
	LastName string `json:"last_name" binding:"required,min=2,max=30"`
	Age          uint64    `json:"age" binding:"required"`
	MatricNumber string    `json:"matric_number" binding:"required,min=6,max=6"`
	DateCreated  time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Room Room `json:"room"`
	Bed Bed `json:"bed"`
	StudentId string `json:"student_id"`
}

func NewStudent(firstName string, lastName string, age uint64, maricNumber string) *Student {
	return &Student{FirstName: firstName, LastName: lastName, Age: age, MatricNumber: maricNumber}
}
