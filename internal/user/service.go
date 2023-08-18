package user

import (
	"context"
	"fmt"
	"golang.org/x/exp/rand"
)

type Service interface {
	CreateUser(ctx context.Context, dto *CreateUserDto) (string, error)
	GetUserById(ctx context.Context, id int) (*Model, error)
	GetAllUsers(ctx context.Context) (*[]Model, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, id int, user *UpdateUserDto) (string, error)
	GenerateRandomUser(ctx context.Context) (*Model, error)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) CreateUser(ctx context.Context, dto *CreateUserDto) (string, error) {
	err := s.storage.CreateOne(dto)

	if err != nil {
		return "Failed to create user", err
	}

	return "User successful create", nil
}

func (s *service) GetUserById(ctx context.Context, id int) (*Model, error) {
	user, err := s.storage.GetOne(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetAllUsers(ctx context.Context) (*[]Model, error) {
	user, err := s.storage.GetAll(1, 1)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	return s.storage.DeleteByID(id)
}

func (s *service) UpdateUser(ctx context.Context, id int, user *UpdateUserDto) (string, error) {
	err := s.storage.UpdateOne(id, user)

	if err != nil {
		return "Failed to update user", err
	}

	return fmt.Sprintf("Successful update user with name: %s", user.FirstName), nil

}

func (s *service) GenerateRandomUser(ctx context.Context) (*Model, error) {
	count, err := s.storage.Count()

	id := rand.Intn(int(count)) + 1

	user, userErr := s.GetUserById(ctx, id)

	if err != nil || userErr != nil {
		return nil, err
	}

	return user, nil

}
