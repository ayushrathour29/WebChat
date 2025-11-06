package store

import(
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn := "host=localhost user=postgres password=ayush123 dbname=webchat port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalf("Failed to connect database: %v", err)
	}

	DB = db
	log.Println("Database connected successfully")
}