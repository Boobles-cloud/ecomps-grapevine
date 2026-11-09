package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DbService struct {
	handler *DbHandler
}

type DbHandler struct {
	dbConn *sql.DB
}

func CreateNewDbHandler() (*DbHandler, error) {

	db, err := sql.Open("mysql", os.Getenv("db_connection_string"))

	if err != nil {
		return &DbHandler{}, err
	}

	return &DbHandler{
		dbConn: db,
	}, nil
}

func CreateNewDbService(dh *DbHandler) *DbService {
	return &DbService{
		handler: dh,
	}
}
