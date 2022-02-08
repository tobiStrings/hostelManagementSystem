package dtos

type SaveRoomDto struct {
	RoomName string `json:"room_name" binding:"required,min=4,max=12"`
	RoomNumber uint64 `json:"room_number" binding:"required,min=1,max=8"`
	NumberOfAvailableBeds uint64  `json:"number_of_available_beds" binding:"required,min=0,max=15"`
}