package model

type Allowance struct {
	AllowanceID uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"index"`
	Month       int     `gorm:"index"`
	Year        int     `gorm:"index"`
	Allowance   float64 `gorm:"type:numeric(10,2)"`
}
