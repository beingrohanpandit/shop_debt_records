package services

import (
	"example.com/internal/core/domain"
	"example.com/internal/core/ports"
	E "github.com/IBM/fp-go/ioeither"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (m *UserService) GetUser(id string) E.IOEither[error, *domain.GetUser] {
	return m.repo.GetUser(id)
}

func (m *UserService) GetUsers() E.IOEither[error, []*domain.GetUser] {
	return m.repo.GetUsers()
}

func (m *UserService) AddUser(data domain.User) E.IOEither[error, domain.User] {
	return m.repo.AddUser(data)
}

func (m *UserService) UpdateUser(data domain.User) E.IOEither[error, domain.User] {
	return m.repo.UpdateUser(data)
}

func (m *UserService) DeleteUser(id string) E.IOEither[error, bool] {
	return m.repo.DeleteUser(id)
}
