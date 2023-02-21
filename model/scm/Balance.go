package scm

import (
	"gorm.io/gorm"
	"time"
)

type BalanceUser struct {
	gorm.Model
	PlayId  string
	Amount  float32
	Type    string
	TrxID   string
	TrxTime time.Time
}
