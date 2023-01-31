package services

import (
	"errors"

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
	var InvalidUser = errors.New("Invalid user type to perform transfers")
	var InvalidValue = errors.New("Insufficient balance to carry out transaction")
	payer_id := transaction.Payer
	payee_id := transaction.Payee

	user_payer, err := s.userRepository.Find(payer_id)
	user_payee, err := s.userRepository.Find(payee_id)
	if err != nil {
		return transaction, err
	}

	if user_payer.User_type == "J" {
		err = InvalidUser
		return transaction, err
	}

	wallet_payer, err := s.walletRepository.FindByUserId(user_payer.ID)
	wallet_payee, err := s.walletRepository.FindByUserId(user_payee.ID)
	if err != nil {
		return transaction, err
	}

	if float32(wallet_payer.Value) < float32(transaction.Value) {
		err = InvalidValue
		return transaction, err
	}

	transaction, err = s.transactionRepository.Save(transaction)

	balancePayer := float32(wallet_payer.Value) - float32(transaction.Value)
	balancePayee := float32(wallet_payee.Value) + float32(transaction.Value)

	walletPayee := &models.Wallet{}
	walletPayee.User_id = user_payee.ID
	walletPayee.Value = float32(balancePayee)

	walletPayer := &models.Wallet{}
	walletPayer.User_id = user_payer.ID
	walletPayer.Value = float32(balancePayer)

	err = s.walletRepository.UpdateByUserId(walletPayer)
	err = s.walletRepository.UpdateByUserId(walletPayee)

	return transaction, err
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
