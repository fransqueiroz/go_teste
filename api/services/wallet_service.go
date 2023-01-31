package services

import (
	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/fransqueiroz/go_teste/api/repository"
)

type WalletService interface {
	Post(*models.Wallet) (*models.Wallet, error)
	Find(uint64) (*models.Wallet, error)
	FindByUserId(uint64) (*models.Wallet, error)
	FindAll() ([]*models.Wallet, error)
	Update(*models.Wallet) error
	Delete(uint64) error
}

type walletServiceImpl struct {
	walletRepository repository.WalletRepository
}

func NewWalletService(walletRepository repository.WalletRepository) *walletServiceImpl {
	return &walletServiceImpl{walletRepository}
}

func (s *walletServiceImpl) Post(wallet *models.Wallet) (*models.Wallet, error) {

	wallet, err := s.walletRepository.Save(wallet)
	if err != nil {
		return wallet, err
	}

	return wallet, err
}

func (s *walletServiceImpl) Find(wallet_id uint64) (*models.Wallet, error) {
	wallet, err := s.walletRepository.Find(wallet_id)
	return wallet, err
}

func (s *walletServiceImpl) FindByUserId(user_id uint64) (*models.Wallet, error) {
	wallet, err := s.walletRepository.FindByUserId(user_id)
	return wallet, err
}

func (s *walletServiceImpl) FindAll() ([]*models.Wallet, error) {
	wallet, err := s.walletRepository.FindAll()
	return wallet, err
}

func (s *walletServiceImpl) Update(wallet *models.Wallet) error {
	err := s.walletRepository.UpdateByUserId(wallet)
	return err
}

func (s *walletServiceImpl) Delete(wallet_id uint64) error {
	err := s.walletRepository.Delete(wallet_id)
	return err
}
