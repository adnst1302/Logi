package scm

import (
	"gorm.io/gorm"
	"time"
)

type BalanceMember struct {
	gorm.Model
	Username string
	Amount   float32
	Type     string
	TrxID    string
	TrxTime  time.Time
}
