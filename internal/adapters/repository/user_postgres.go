package repository

import (
	"example.com/internal/core/domain"
	E "github.com/IBM/fp-go/ioeither"
)

func (m *PostgresRepository) GetUser(id string) E.IOEither[error, *domain.GetUser] {
	data := &domain.GetUser{}
	return E.Right[error](data)
}

func (m *PostgresRepository) GetUsers() E.IOEither[error, *[]domain.GetUser] {
	data := &[]domain.GetUser{}
	return E.Right[error](data)
}

func (m *PostgresRepository) AddUser(data domain.User) E.IOEither[error, domain.User] {
	return E.Right[error](data)
}

func (m *PostgresRepository) UpdateUser(data domain.User) E.IOEither[error, domain.User] {
	return E.Right[error](data)
}

func (m *PostgresRepository) Delete(id string) E.IOEither[error, bool] {
	return E.Right[error](true)
}
