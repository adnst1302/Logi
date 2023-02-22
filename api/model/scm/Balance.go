package scm

import (
	"gorm.io/gorm"
	"time"
)

type BalanceTransaction struct {
	gorm.Model
	TrxId     string
	UserId    string
	Amount    string
	Flow      string
	TypeTrans string
	CreateBy  string
	ApproveBy string
	TrxTime   time.Time
}
