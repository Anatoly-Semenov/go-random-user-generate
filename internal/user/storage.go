package user

type Storage interface {
	GetOne(id int) (*Model, error)
	GetAll(limit int, offset int) (*[]Model, error)
	DeleteByID(id int) error
	UpdateOne(id int, user *UpdateUserDto) error
	CreateOne(userDto *CreateUserDto) error
	Count() (int64, error)
}
