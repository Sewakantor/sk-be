package config

import (
	"fmt"
	"github.com/joho/godotenv"
	_propertyRepo "github.com/sewakantor/sw-be/repository/databases/property"
	_usersRepo "github.com/sewakantor/sw-be/repository/databases/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

	})
	if err != nil {
		panic(err.Error())
	}
	dbMigrate(db)

	return db
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_usersRepo.Users{}, &_propertyRepo.Complex{}, &_propertyRepo.Building{}, &_propertyRepo.Review{})
}