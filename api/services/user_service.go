package services

import (
	"errors"

	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/fransqueiroz/go_teste/api/repository"
)

type UserService interface {
	Post(*models.User) (*models.User, error)
	GetUser(uint64) (*models.User, error)
	FindAll() ([]*models.User, error)
	Update(*models.User) error
	Delete(uint64) error
}

type userServiceImpl struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
}

func NewUserService(userRepository repository.UserRepository, walletRepository repository.WalletRepository) *userServiceImpl {
	return &userServiceImpl{userRepository, walletRepository}
}

func Validate(user *models.User) error {
	var ErrUserEmptyName = errors.New("user.name cannot be empty")
	var ErrUserNameMaxLen = errors.New("user.name max length is 200")
	var ErrUserEmptyCPF = errors.New("user.cpf cannot be empty")
	// var ErrUserWrongFormatCPF = errors.New("user.cpf must be XXX.XXX.XXX-XX")
	var ErrUserEmptyEmail = errors.New("user.email cannot be empty")
	var ErrUserEmptyPassword = errors.New("user.password cannot be empty")

	if user.Name == "" {
		return ErrUserEmptyName
	}

	if len(user.Name) > 300 {
		return ErrUserNameMaxLen
	}

	if user.CPF == "" {
		return ErrUserEmptyCPF
	}

	if user.Email == "" {
		return ErrUserEmptyEmail
	}

	if user.Password == "" {
		return ErrUserEmptyPassword
	}

	return nil
}

func (s *userServiceImpl) Post(user *models.User) (*models.User, error) {

	var ErrUserExists = errors.New("user already registered")

	query := "email = '" + user.Email + "'"
	user_response, _ := s.userRepository.GetUserForQuery(query)

	if user.Email == user_response.Email {
		return user, ErrUserExists
	}

	query = "cpf = '" + user.CPF + "'"
	user_response, _ = s.userRepository.GetUserForQuery(query)

	if user.CPF == user_response.CPF {
		return user, ErrUserExists
	}

	err := Validate(user)
	if err != nil {
		return user, err
	}

	user, err = s.userRepository.Save(user)

	if err != nil {
		return user, err
	}

	walletStruct := &models.Wallet{}

	walletStruct.User_id = user.ID

	_, err = s.walletRepository.Save(walletStruct)

	if err != nil {
		return user, err
	}

	return user, err
}

func (s *userServiceImpl) GetUser(user_id uint64) (*models.User, error) {
	user, err := s.userRepository.Find(user_id)
	return user, err
}

func (s *userServiceImpl) FindAll() ([]*models.User, error) {
	var err error
	user, err := s.userRepository.FindAll()
	return user, err
}

func (s *userServiceImpl) Update(user *models.User) error {
	err := Validate(user)
	if err != nil {
		return err
	}

	err = s.userRepository.Update(user)
	return err
}

func (s *userServiceImpl) Delete(user_id uint64) error {

	wallet, err := s.walletRepository.FindByUserId(user_id)
	if err != nil {
		return err
	}

	if wallet.Value > 0 {
		return errors.New("user with value greater than 0")
	}

	err = s.userRepository.Delete(user_id)

	if err != nil {
		return err
	}

	return err
}
