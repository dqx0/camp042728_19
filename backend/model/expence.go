package model

import (
	"time"
)

type Expense struct {
	ExpenseID       uint      `gorm:"primaryKey"`
	UserID          uint      `gorm:"index"`
	Date            time.Time `gorm:"index"`
	Title           string    `gorm:"type:varchar(255)" json:"title"`
	AmountSpent     float64   `gorm:"type:numeric(10,2)" json:"money"`
	RemainingAmount float64   `gorm:"type:numeric(10,2)"`
}
