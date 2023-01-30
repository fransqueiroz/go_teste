package models

type Transaction struct {
	Model
	Value float64 `gorm:"default:0.00" json:"value"`
	Payer uint64  `gorm:"not null" json:"payer"`
	Payee uint64  `gorm:"not null" json:"payee"`
}
