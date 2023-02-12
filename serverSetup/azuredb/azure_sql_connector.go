package azuredb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var Db *sql.DB

type DbConfig struct {
	Server   string `toml:"server"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

func (db *DbConfig) GetConnString() string {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		db.Server, db.User, db.Password, db.Port, db.Database)
	return connString
}

func InstantiateDBpool(connString string) (*sql.DB, error) {
	var err error
	// Create connection pool
	Db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
		//log.Fatal("Error creating connection pool: ", err.Error())
	}
	return Db, nil
}

func Check_ping(db *sql.DB) bool {
	var err error
	//db, err := InstantiateDBpool(GetConnString())
	db.Conn(context.Background())
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf(fmt.Sprintf("Db: %s , was connected", ""))
	return true
}
