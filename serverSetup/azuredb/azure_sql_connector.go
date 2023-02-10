package azuredb

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

var Db *sql.DB
var server = "1"
var port = 11
var user = "q"
var password = "1"
var database = "ma"

func GetConnString() string {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
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

func Check_ping() bool {
	var err error
	db, err := InstantiateDBpool(GetConnString())
	db.Conn(context.Background())
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf(fmt.Sprintf("Db: %s , was connected", ""))
	return true
}
