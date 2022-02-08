package room_service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"hostelManagementSystem/dtos"
	"hostelManagementSystem/models"
	"hostelManagementSystem/service/mocks"
	"testing"
)
var roomServiceImpl RoomServiceImpl
var saveRoomDto dtos.SaveRoomDto
var roomMock mocks.RoomMock
var room models.Room
var roomErrorMock mocks.RoomMockError
var err error
func init()  {
	roomServiceImpl = RoomServiceImpl{}
	saveRoomDto = dtos.SaveRoomDto{}
	roomMock = mocks.RoomMock{}
	room= models.Room{}
	roomErrorMock = mocks.RoomMockError{}
	roomErrorMock.Err = err
}
func TestValidateDtoInputsWhenRequiredValuesAreNotPassedIn(t *testing.T){
	dtoToValidate := dtos.SaveRoomDto{}

	err:=validateDtoInputs(dtoToValidate)

	assert.NotNil(t, err)
}
func TestValidateDtoInputsWhenOnlyTheRoomNameWasPassed(t *testing.T)  {
	dtoToValidate := dtos.SaveRoomDto{}
	dtoToValidate.RoomName = "Kale"

	err:=validateDtoInputs(dtoToValidate)

	assert.NotNil(t, err)
}

func TestValidateDtoInputsWhenOnlyTheRoomNumberWasGiven(t *testing.T){
	dtoToValidate := dtos.SaveRoomDto{}
	dtoToValidate.RoomNumber = 8

	err:=validateDtoInputs(dtoToValidate)

	assert.NotNil(t, err)
}

func TestValidateDtoInputsWhenOnlyTheNumberOfAvailableBedsWasGiven(t *testing.T){
	dtoToValidate := dtos.SaveRoomDto{}
	dtoToValidate.NumberOfAvailableBeds = 8

	err:=validateDtoInputs(dtoToValidate)

	assert.NotNil(t, err)
}

func TestValidateDtoInputsWhenTheValueOfRoomNumberIsEqualsZero(t *testing.T){
	dtoToValidate := dtos.SaveRoomDto{}
	dtoToValidate.RoomNumber = 0

	err:=validateDtoInputs(dtoToValidate)

	assert.NotNil(t, err)
}

func TestValidateDtoInputsWithAllValuesGivenCorrectly(t *testing.T){
	dtoToValidate := dtos.SaveRoomDto{}
	dtoToValidate.RoomName = "Jaja"
	dtoToValidate.NumberOfAvailableBeds = 100
	dtoToValidate.RoomNumber = 0

	err:=validateDtoInputs(dtoToValidate)

	assert.NotNil(t, err)
}

func TestMapDtoToRoom(t *testing.T)  {
	dtoToMap := dtos.SaveRoomDto{}
	dtoToMap.RoomName = "Jaja"
	dtoToMap.NumberOfAvailableBeds = 100
	dtoToMap.RoomNumber = 8

	room := models.Room{}

	room = mapDtoToRoom(dtoToMap, room)

	assert.NotNil(t, room)
	assert.Equal(t, "Jaja",room.RoomName)
	assert.Equal(t, uint64(100),room.NumberOfAvailableBeds)
	assert.Equal(t, uint64(8),room.RoomNumber)
}

func TestSaveRoom(t *testing.T)  {
	saveRoomDto.RoomName = "Jaja"
	saveRoomDto.RoomNumber = 6
	saveRoomDto.NumberOfAvailableBeds = 20

	roomMock.On("SaveRoom",saveRoomDto).Return(room,nil)

	returnedRoom,_ := roomMock.SaveRoom(saveRoomDto)

	assert.NotNil(t, returnedRoom)
}

func TestSaveRoomWhenDtoInPutsAreNotCorrectlyImputed(t *testing.T)  {
	saveRoomDto.RoomNumber = 30
	saveRoomDto.NumberOfAvailableBeds = 90
	_,err :=roomServiceImpl.SaveRoom(saveRoomDto)
	assert.NotNil(t, err)
}

func TestSaveRoomWhenRoomHasAlreadyBeenSavedBefore(t *testing.T){
	saveRoomDto.RoomName = "Jaja"
	saveRoomDto.RoomNumber = 6
	saveRoomDto.NumberOfAvailableBeds = 20

	roomErrorMock.Err = errors.New("room already exist ")

	roomErrorMock.On("SaveRoom",saveRoomDto).Return(nil,roomErrorMock.Err)

	_,err1 := roomErrorMock.SaveRoom(saveRoomDto)


	assert.NotNil(t, err1)

	assert.Error(t, roomErrorMock.Err,"room already exist ")
}