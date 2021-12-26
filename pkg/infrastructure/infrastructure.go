package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/y-mabuchi/torasemi-todo-api/pkg/config"
)

var db *sql.DB

func NewDB(conf config.DBConfig) *sql.DB {
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}
	instanceConnectionName := os.Getenv("CLOUD_SQL_CONNECTION_NAME")
	var dsn string
	if instanceConnectionName != "" {
		dsn = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true&loc=Local&charset=utf8mb4&&multiStatements=true", conf.DBUser, conf.DBPass, socketDir, instanceConnectionName, conf.DBName)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4&multiStatements=true", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName)
	}

	dbTemp, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("action=sql.Open, status=error: %v", err)
	}

	db = dbTemp
	if err := db.Ping(); err != nil {
		log.Fatalf("action=db.Ping, status=error: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
