package model

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model

	WalletID uuid.UUID `gorm:"primaryKey"`
	Amount   float64
}
