package dtos

type StudentDto struct {
	FirstName string `json:"first_name" binding:"required,min=4,max=30"`
	LastName string `json:"last_name" binding:"required,min=4,max=30"`
	Age      uint64 `json:"age" binding:"required"`
	MaricNumber string `json:"matric_number" binding:"required,min=6,max=6"`
}

func NewStudentDto(firstName string, lastName string, age uint64, maricNumber string) StudentDto {
	return StudentDto{FirstName: firstName, LastName: lastName, Age: age, MaricNumber: maricNumber}
}

