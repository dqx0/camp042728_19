package model

type Allowance struct {
	AllowanceID uint `gorm:"primaryKey"`
	UserID      uint `gorm:"index"`
	Month       int  `gorm:"index" json:"month"`
	Year        int  `gorm:"index" json:"year"`
	Allowance   int  `gorm:"int" json:"money"`
}
