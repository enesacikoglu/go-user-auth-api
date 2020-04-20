package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"go-user-auth-api/infrastructure/configuration"
	"time"
)

type MSSqlConnection struct {
	db *sql.DB
}

func NewMSSqlConnection(dbConfig configuration.DatabaseConfig) *MSSqlConnection {
	connectionString := GetConnectionString(dbConfig)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		errMessage := "Error connecting to trendyol db"
		panic(errMessage)
	}
	db.SetMaxOpenConns(dbConfig.MaxOpenConnection)
	db.SetMaxIdleConns(dbConfig.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Hour)

	return &MSSqlConnection{
		db: db,
	}
}

func GetConnectionString(dbConfig configuration.DatabaseConfig) string {
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", dbConfig.UserName, dbConfig.Password, dbConfig.Host, dbConfig.Database)
	return connectionString
}
