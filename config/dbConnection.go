package config

import (
	"fmt"
	"os"

	"my-fiber-app/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")

	// Формируем строку подключения
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbhost, dbuser, dbpassword, dbname, dbport)

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("DB connection failed: " + err.Error()) // выводим ошибку при неудачном подключении
	}

	DB = db
	fmt.Println("DB connected successfully")

	AutoMigrate(db)
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Discount{},
		&models.Order{},
		&models.Payment{},
		&models.Product{},
	)
}
