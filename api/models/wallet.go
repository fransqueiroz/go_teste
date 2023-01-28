package models

func (p *Wallet) Validate() error {
	return nil
}

type Wallet struct {
	Model
	User_id uint64  `gorm:"not null" json:"user_id"`
	Value   float32 `gorm:"default:0.00" json:"value"`
}
