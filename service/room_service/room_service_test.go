package room_service

import (
	"github.com/stretchr/testify/assert"
	"hostelManagementSystem/dtos"
	"testing"
)

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