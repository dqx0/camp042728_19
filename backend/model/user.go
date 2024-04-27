package model

type User struct {
	UserID   uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
