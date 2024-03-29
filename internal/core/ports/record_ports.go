package ports

import (
	"example.com/internal/core/domain"
	E "github.com/IBM/fp-go/ioeither"
)

type RecordRepository interface {
	AddRecord(data domain.Record) E.IOEither[error, domain.Record]
}
