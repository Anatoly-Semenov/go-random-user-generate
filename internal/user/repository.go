package user

import (
	"errors"
	"gorm.io/gorm"
)

type repository struct {
	Db *gorm.DB
}

func NewRepository(Db *gorm.DB) *repository {
	return &repository{Db: Db}
}

func (r *repository) GetAll(limit int, offset int) (*[]Model, error) {
	users := []Model{}

	err := r.Db.Model(&Model{}).Find(&users).Error

	if err != nil {
		return &[]Model{}, err
	}

	return &users, err
}

func (r *repository) GetOne(id int) (*Model, error) {
	var user Model

	err := r.Db.Model(Model{}).Where("id = ?", id).Take(&user).Error

	if err != nil {
		return &Model{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &Model{}, errors.New("User Not Found")
	}

	return &user, err
}

func (r *repository) DeleteByID(id int) error {
	var user Model
	err := r.Db.Model(Model{}).Where("id = ?", id).Delete(&user).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateOne(id int, userDto *UpdateUserDto) error {
	request := r.Db.Model(Model{}).Where("id = ?", id).Updates(&userDto)

	err := request.Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) CreateOne(userDto *CreateUserDto) error {
	request := r.Db.Model(Model{}).Create(&userDto)

	err := request.Error

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Count() (int64, error) {
	var count int64

	request := r.Db.Model(Model{}).Count(&count)

	err := request.Error

	if err != nil {
		return count, err
	}

	return count, nil
}
