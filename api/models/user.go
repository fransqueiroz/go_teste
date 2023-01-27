package models

import "errors"

var ErrUserEmptyName = errors.New("user.name cannot be empty")
var ErrUserNameMaxLen = errors.New("user.name max length is 200")
var ErrUserEmptyCPF = errors.New("user.cpf cannot be empty")
var ErrUserWrongFormatCPF = errors.New("user.cpf must be XXX.XXX.XXX-XX")
var ErrUserEmptyEmail = errors.New("user.email cannot be empty")
var ErrUserEmptyPassword = errors.New("user.password cannot be empty")

func (p *User) Validate() error {
	if p.Name == "" {
		return ErrUserEmptyName
	}

	return nil
}

type User struct {
	Model
	Name      string `gorm:"size:200;not null" json:"name"`
	CPF       string `gorm:"size:12 not null" json:"cpf"`
	Email     string `gorm:"size:150" json:"email"`
	Password  string `gorm:"size:300" json:"password"`
	User_type string `json:"user_type" gorm:"user_type:enum('F', 'J');default:'F'"`
}
