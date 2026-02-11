package database
import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func ConnectDB() *gorm.DB{
	dsn:=os.Getenv("DB_DSN")
	if dsn==""{
		dsn="host=localhost user=postgres password=strongpassword dbname=inventory_db port=5432 sslmode=disable"
	}
	db, err:=gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")
	return db
}