package models

type User struct {
	Model
	Name      string `gorm:"size:200;not null" json:"name"`
	CPF       string `gorm:"size:200;not null;unique" json:"cpf"`
	Email     string `gorm:"size:150;not null;unique" json:"email"`
	Password  string `gorm:"size:300" json:"password"`
	User_type string `json:"user_type" gorm:"user_type:enum('F', 'J');default:'F'"`
}
