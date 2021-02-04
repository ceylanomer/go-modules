package customercontext

import (
	"database/sql"
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

//CreateDBIfNotExist is creates db if not exist
func CreateDBIfNotExist() (string, error) {
	serverConnection, serverConnectionError := sql.Open("mssql", os.Getenv("SERVER_URL"))
	if serverConnectionError == nil {
		isServerConnectedError := serverConnection.Ping()
		if isServerConnectedError == nil {
			databaseConnection, databaseConnectionError := sql.Open("mssql", os.Getenv("DATABASE_URL"))
			if databaseConnectionError != nil {
				fmt.Println("Database Error: " + databaseConnectionError.Error())
				return "Database Error: " + databaseConnectionError.Error(), databaseConnectionError
			} else {
				isDatabaseConnectedError := databaseConnection.Ping()
				if isDatabaseConnectedError != nil {
					_, creatingDatabaseError := serverConnection.Exec("CREATE DATABASE goDB")
					if creatingDatabaseError != nil {
						fmt.Println("Database Creating Error: " + creatingDatabaseError.Error())
						return "Database Creating Error: " + creatingDatabaseError.Error(), creatingDatabaseError
					}
				} else {
					fmt.Println("Database already exist.")
					return "Database already exist.", nil
				}
			}
		} else {
			fmt.Println("Server Error: " + isServerConnectedError.Error())
			return "Server Error: " + isServerConnectedError.Error(), isServerConnectedError
		}
	} else {
		fmt.Println(serverConnectionError.Error())
		return serverConnectionError.Error(), serverConnectionError
	}
	return "nil", nil
}
