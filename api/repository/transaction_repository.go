package repository

import (
	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/jinzhu/gorm"
)

type TransactionRepository interface {
	Save(*models.Transaction) (*models.Transaction, error)
	Find(uint64) (*models.Transaction, error)
	FindByPayerId(uint64) (*models.Transaction, error)
	FindByPayeeId(uint64) (*models.Transaction, error)
	FindAll() ([]*models.Transaction, error)
	Delete(uint64) error
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepositoryImpl {
	return &transactionRepositoryImpl{db}
}

func (r *transactionRepositoryImpl) Save(transaction *models.Transaction) (*models.Transaction, error) {
	tx := r.db.Debug().Begin()
	err := tx.Debug().Model(&models.Transaction{}).Create(transaction).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return transaction, tx.Commit().Error
}

func (r *transactionRepositoryImpl) Find(transaction_id uint64) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	err := r.db.Debug().Model(&models.Transaction{}).Where("id = ?", transaction_id).Find(transaction).Error
	return transaction, err
}

func (r *transactionRepositoryImpl) FindAll() ([]*models.Transaction, error) {
	transaction := []*models.Transaction{}
	err := r.db.Debug().Model(&models.Transaction{}).Find(&transaction).Error
	return transaction, err
}

func (r *transactionRepositoryImpl) FindByPayeeId(payee_id uint64) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	err := r.db.Debug().Model(&models.Transaction{}).Where("payee = ?", payee_id).Find(transaction).Error
	return transaction, err
}

func (r *transactionRepositoryImpl) FindByPayerId(payer_id uint64) (*models.Transaction, error) {
	transaction := &models.Transaction{}
	err := r.db.Debug().Model(&models.Transaction{}).Where("payer = ?", payer_id).Find(transaction).Error
	return transaction, err
}

func (r *transactionRepositoryImpl) Delete(transaction_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Transaction{}).Where("id = ?", transaction_id).Delete(&models.Transaction{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
