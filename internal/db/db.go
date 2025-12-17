package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"book.com/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// channel to take data from user request/handler function and store that in DB using reciever go-routine.
var BookCh = make(chan models.Book, 100)

func Init() {
	//DB Credentials loading
	godotenv.Load(".env")
	Db_name := os.Getenv("DB_NAME")
	Db_user := os.Getenv("DB_USER")
	Db_pass := os.Getenv("DB_PASSWORD")
	Port := os.Getenv("DB_PORT")
	Host := os.Getenv("DB_HOST")
	Ssl := os.Getenv("SSLMODE")

	if Db_name == "" || Db_user == "" || Db_pass == "" || Port == "" || Host == "" || Ssl == "" {
		log.Fatal("Something is missing in DB credentials.")
	}
	//DSN creation
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		Host, Db_user, Db_pass, Db_name, Port, Ssl,
	)
	//Logging service
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	//checking if connection is established or not.
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get sql.DB", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}
	fmt.Println("Database connection successful!")
}
