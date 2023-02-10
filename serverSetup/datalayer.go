package main

import (
	"context"
	"customerSite/serverSetup/azuredb"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Controller struct {
	Database *sql.DB
}

type userDetails struct {
	ID           int
	name         string
	email        string
	passwordHash string
	accessLevel  int
}

type issueDetails struct {
	ID          int    `json:"id"`
	IssueType   int    `json:"type"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	User        int    `json:"user"`
	LastUpdate  int    `json:"date"`
}

func validateUserCredentials(email string, passwordHash string) (bool, userDetails) {

	emailUser := "aaa@gmail.com"
	hashCheck := "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8"

	if email == emailUser && passwordHash == hashCheck {
		return true, userDetails{
			ID:    123,
			name:  "ram",
			email: emailUser,
		}
	}
	return false, userDetails{}
}

func enterNewIssue(ctx context.Context, ct *Controller, issueInput issueDetails) bool {
	const (
		userTable = "DataSchema.IssueDetails"
		columns   = " IssueType, Description, Status, UserID, LastUpdate "
	)
	var values string
	values = fmt.Sprintf(" %d, '%s' , %d, %d, %d ",
		issueInput.IssueType,
		issueInput.Description,
		issueInput.Status,
		issueInput.User,
		time.Now().Unix())
	fmt.Println(issueInput)
	tsql := fmt.Sprintf(
		`INSERT INTO %s ( %s ) VALUES ( %s); 
		select isNull(SCOPE_IDENTITY(), -1);`,
		userTable, columns, values,
	)
	fmt.Println(tsql)
	_, err := submitNewEntryInDb(ctx, ct.Database, tsql)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func submitNewEntryInDb(ctx context.Context, db *sql.DB, queryString string) (int64, error) {

	var err error
	//check db is available
	if db == nil {
		err = errors.New(" db is null, transaction didn't go through")
		return -1, err
	}
	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return -1, err
	}

	stmt, err := db.Prepare(queryString)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
	)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}
	return newID, nil
}

// backend access funcrtions
func (ct *Controller) getSingleQueryResult(ctx context.Context, query string, dataStruct any) (any, error) {
	azuredb.Check_ping()
	return dataStruct, nil
}
