package customercontext

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//ConnectDB connection utils to db
func ConnectDB() (*gorm.DB, error) {
	connString := "sqlserver://sa:Sa123456@localhost:1433?database=goDB"
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
