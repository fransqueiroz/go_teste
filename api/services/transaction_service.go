package services

import (
	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/fransqueiroz/go_teste/api/repository"
)

type TransactionService interface {
	Post(*models.Transaction) (*models.Transaction, error)
	Find(uint64) (*models.Transaction, error)
	FindByPayerId(uint64) (*models.Transaction, error)
	FindByPayeeId(uint64) (*models.Transaction, error)
	FindAll() ([]*models.Transaction, error)
	Delete(uint64) error
}

type transactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
	userRepository        repository.UserRepository
	walletRepository      repository.WalletRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository, userRepository repository.UserRepository, walletRepository repository.WalletRepository) *transactionServiceImpl {
	return &transactionServiceImpl{transactionRepository, userRepository, walletRepository}
}

func (s *transactionServiceImpl) Post(transaction *models.Transaction) (*models.Transaction, error) {

	wallet, err := s.transactionRepository.Save(transaction)

	return wallet, err
}

func (s *transactionServiceImpl) Find(transaction_id uint64) (*models.Transaction, error) {
	transaction, err := s.transactionRepository.Find(transaction_id)
	return transaction, err
}

func (s *transactionServiceImpl) FindByPayerId(payee_id uint64) (*models.Transaction, error) {
	transaction, err := s.transactionRepository.FindByPayeeId(payee_id)
	return transaction, err
}

func (s *transactionServiceImpl) FindByPayeeId(payer_id uint64) (*models.Transaction, error) {
	transaction, err := s.transactionRepository.FindByPayeeId(payer_id)
	return transaction, err
}

func (s *transactionServiceImpl) FindAll() ([]*models.Transaction, error) {
	transaction, err := s.transactionRepository.FindAll()
	return transaction, err
}

func (s *transactionServiceImpl) Delete(transaction_id uint64) error {
	err := s.transactionRepository.Delete(transaction_id)
	return err
}
