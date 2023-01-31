package models

func ValidateTransaction(transaction Transaction, user User, wallet Wallet) error {
	return nil
}

type Transaction struct {
	Model
	Value float32 `gorm:"default:0.00" json:"value"`
	Payer uint64  `gorm:"not null" json:"payer"`
	Payee uint64  `gorm:"not null" json:"payee"`
}
