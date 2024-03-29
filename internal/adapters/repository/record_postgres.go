package repository

import (
	"errors"

	"example.com/internal/adapters/repository/query"
	"example.com/internal/core/domain"
	E "github.com/IBM/fp-go/ioeither"
)

func (m *PostgresRepository) AddRecord(data domain.Record) E.IOEither[error, domain.Record] {
	var payment_due float64
	var record_id int

	// Begin a transaction
	tx := m.db.Begin()

	// Rollback the transaction if any error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Add Product
	err := tx.Raw(query.AddProduct, data.ProductDetails).Scan(&data.ProductDetails.ProductId).Error
	if err != nil {
		tx.Rollback()
		return E.Left[domain.Record](errors.New("data not added"))
	}

	// Calculate Payment Due
	payment_due = data.ProductDetails.ProductPrice - data.CreditAmount

	// Add Payment Record
	err = tx.Raw(query.AddPaymentRecord, payment_due, data.UserId, data.ProductDetails.ProductId).Scan(&record_id).Error
	if err != nil {
		tx.Rollback()
		return E.Left[domain.Record](errors.New("data not added"))
	}

	// Add Payment Logs
	err = tx.Raw(query.AddPaymentLogs, data.CreditAmount, payment_due, record_id).Scan(&record_id).Error
	if err != nil {
		tx.Rollback()
		return E.Left[domain.Record](errors.New("data not added"))
	}

	// Commit the transaction if all operations are successful
	if err := tx.Commit().Error; err != nil {
		return E.Left[domain.Record](errors.New("failed to commit transaction"))
	}

	return E.Right[error](data)
}
