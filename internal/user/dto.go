package user

type CreateUserDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dob       string `json:"dob"`
	Age       int    `json:"age"`
	Sex       Sex    `json:"sex"`
	Avatar    string `json:"avatar"`
}

type UpdateUserDto struct {
	CreateUserDto
}
