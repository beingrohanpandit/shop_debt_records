package ports

import (
	"example.com/internal/core/domain"
	E "github.com/IBM/fp-go/ioeither"
)

type UserRepository interface {
	AddUser(data domain.User) E.IOEither[error, domain.User]
	GetUser(id string) E.IOEither[error, *domain.GetUser]
	GetUsers() E.IOEither[error, []*domain.GetUser]
	UpdateUser(data domain.User) E.IOEither[error, domain.User]
	DeleteUser(id string) E.IOEither[error, bool]
}
