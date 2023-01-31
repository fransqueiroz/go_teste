package services

import (
	"fmt"
	"strconv"

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

	user_id := wallet.User_id
	wallet_user, err := s.walletRepository.FindByUserId(user_id)
	if err != nil {
		return err
	}

	balance := float32(wallet_user.Value) + float32(wallet.Value)

	value_stg := fmt.Sprintf("%.2f", balance)

	value, err := strconv.ParseFloat(value_stg, 32)

	wallet.Value = float32(value)

	err = s.walletRepository.UpdateByUserId(wallet)
	return err
}

func (s *walletServiceImpl) Delete(wallet_id uint64) error {
	err := s.walletRepository.Delete(wallet_id)
	return err
}
