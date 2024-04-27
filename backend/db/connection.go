package db

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dqx0/camp042728_19/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func NewDB() *gorm.DB {
	once.Do(func() {
		// 環境変数から接続情報を取得
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
		}

		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbName := os.Getenv("POSTGRES_DB")
		dbHost := os.Getenv("POSTGRES_HOST")
		dbPort := os.Getenv("POSTGRES_PORT")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)
		var errOpen error
		db, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if errOpen != nil {
			log.Fatalf("データベース接続に失敗しました: %v", errOpen)
		}

		// モデルの自動マイグレーション
		errAutoMigrate := db.AutoMigrate(&model.User{}, &model.Expense{}, &model.Allowance{})
		if errAutoMigrate != nil {
			log.Fatalf("自動マイグレーションに失敗しました: %v", errAutoMigrate)
		}
	})

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
