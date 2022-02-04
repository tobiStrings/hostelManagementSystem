package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hostel struct {
	HostelId primitive.ObjectID `json:"hostel_id"`
	Rooms []Room `json:"rooms"`
}

type Bed struct {
	BedId primitive.ObjectID `json:"bed_id"`
	BedNumber uint64 `json:"bed_number" binding:"required"`
	//StudentAssigned Student `json:"student_assigned" binding:"required"`
}

type Room struct {
	RoomId primitive.ObjectID `json:"room_id"`
	Beds []Bed `json:"beds"`
	RoomName string `json:"room_name" binding:"required,min=4,max=12"`
	RoomNumber uint64 `json:"room_number" binding:"required,min=1,max=8"`
	NumberOfAvailableBeds uint64 `json:"number_of_available_beds" binding:"required,min=0,max=15"`
}