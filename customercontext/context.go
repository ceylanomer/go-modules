package customercontext

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//ConnectDB connection utils to db
func ConnectDB() (*gorm.DB, error) {
	godotenv.Load()
	connString := os.Getenv("DATABASE_URL")
	fmt.Println(connString)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
