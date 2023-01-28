package repository

import (
	"time"

	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/jinzhu/gorm"
)

type WalletRepository interface {
	Save(*models.Wallet) (*models.Wallet, error)
	Find(uint64) (*models.Wallet, error)
	FindByUserId(uint64) (*models.Wallet, error)
	FindAll() ([]*models.Wallet, error)
	UpdateByUserId(*models.Wallet) error
	Delete(uint64) error
}

type walletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) *walletRepositoryImpl {
	return &walletRepositoryImpl{db}
}

func (r *walletRepositoryImpl) Save(wallet *models.Wallet) (*models.Wallet, error) {
	tx := r.db.Debug().Begin()
	err := tx.Debug().Model(&models.Wallet{}).Create(wallet).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return wallet, tx.Commit().Error
}

func (r *walletRepositoryImpl) Find(wallet_id uint64) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	err := r.db.Debug().Model(&models.Wallet{}).Where("id = ?", wallet_id).Find(wallet).Error
	return wallet, err
}

func (r *walletRepositoryImpl) FindByUserId(user_id uint64) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	err := r.db.Debug().Model(&models.Wallet{}).Where("user_id = ?", user_id).Find(wallet).Error
	return wallet, err
}

func (r *walletRepositoryImpl) FindAll() ([]*models.Wallet, error) {
	wallet := []*models.Wallet{}
	err := r.db.Debug().Model(&models.Wallet{}).Find(&wallet).Error
	return wallet, err
}

func (r *walletRepositoryImpl) UpdateByUserId(wallet *models.Wallet) error {
	tx := r.db.Begin()

	columns := map[string]interface{}{
		"user_id":    wallet.User_id,
		"value":      wallet.Value,
		"updated_at": time.Now(),
	}

	err := tx.Debug().Model(&models.Wallet{}).Where("user_id = ?", wallet.User_id).UpdateColumns(columns).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *walletRepositoryImpl) Delete(wallet_id uint64) error {
	tx := r.db.Begin()
	err := tx.Debug().Model(&models.Wallet{}).Where("id = ?", wallet_id).Delete(&models.Wallet{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
