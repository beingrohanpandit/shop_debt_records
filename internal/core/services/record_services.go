package services

import (
	"example.com/internal/core/domain"
	"example.com/internal/core/ports"
	E "github.com/IBM/fp-go/ioeither"
)

type RecordService struct {
	repo ports.RecordRepository
}

func NewRecordService(repo ports.RecordRepository) *RecordService {
	return &RecordService{
		repo: repo,
	}
}

func (m *RecordService) AddRecord(data domain.Record) E.IOEither[error, domain.Record] {
	return m.repo.AddRecord(data)
}
