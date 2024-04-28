package model

import (
	"time"
)

type Expense struct {
	ExpenseID       uint      `gorm:"primaryKey"`
	UserID          uint      `gorm:"index"`
	Date            time.Time `gorm:"index"`
	Title           string    `gorm:"type:varchar(255)" json:"title"`
	AmountSpent     int       `gorm:"int" json:"money"`
	RemainingAmount int       `gorm:"int"`
}
