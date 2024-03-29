package repository

import (
	"errors"

	"example.com/internal/adapters/repository/query"
	"example.com/internal/core/domain"
	E "github.com/IBM/fp-go/ioeither"
)

func (m *PostgresRepository) GetUser(id string) E.IOEither[error, *domain.GetUser] {
	data := &domain.GetUser{}
	err := m.db.Raw(query.GetUser, id).Scan(data).Error
	if err != nil {
		return E.Left[*domain.GetUser](errors.New("data not found"))
	}
	return E.Right[error](data)
}

func (m *PostgresRepository) GetUsers() E.IOEither[error, *[]domain.GetUser] {
	data := &[]domain.GetUser{}
	err := m.db.Raw(query.GetUsers).Scan(data).Error
	if err != nil {
		return E.Left[*[]domain.GetUser](errors.New("data not found"))
	}
	return E.Right[error](data)
}

func (m *PostgresRepository) AddUser(data domain.User) E.IOEither[error, domain.User] {
	err := m.db.Exec(query.AddUser, data.FirstName, data.LastName, data.PhoneNo, data.Address).Error
	if err != nil {
		return E.Left[domain.User](errors.New("data not added"))
	}
	return E.Right[error](data)
}

func (m *PostgresRepository) UpdateUser(data domain.User) E.IOEither[error, domain.User] {
	err := m.db.Exec(query.UpdateUser, data.FirstName, data.LastName, data.PhoneNo, data.Address).Error
	if err != nil {
		return E.Left[domain.User](errors.New("data not updated"))
	}
	return E.Right[error](data)
}

func (m *PostgresRepository) Delete(id string) E.IOEither[error, bool] {
	err := m.db.Exec(query.DeleteUser, id).Error
	if err != nil {
		return E.Left[bool](errors.New("data not deleted"))
	}
	return E.Right[error](true)
}
