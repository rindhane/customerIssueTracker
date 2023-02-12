package main

import (
	"context"
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
	Name         string
	Email        string
	PasswordHash string
	Phone        string
	AccessLevel  int
}

type issueDetails struct {
	ID          int    `json:"id"`
	IssueType   int    `json:"type"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	User        int    `json:"user"`
	LastUpdate  int    `json:"date"`
}

func validateUserCredentials(ctx context.Context, ct *Controller, email string, passwordHash string) (bool, userDetails) {
	var userFound userDetails
	userFound, _ = isUserAuth(ctx, ct, email, passwordHash)
	if email == userFound.Email && passwordHash == userFound.PasswordHash {
		return true, userDetails{
			ID:          userFound.ID,
			Name:        userFound.Name,
			Email:       userFound.Email,
			AccessLevel: userFound.AccessLevel,
		}
	}
	fmt.Println(userFound)
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
	//fmt.Println(issueInput)
	tsql := fmt.Sprintf(
		`INSERT INTO %s ( %s ) VALUES ( %s); 
		select isNull(SCOPE_IDENTITY(), -1);`,
		userTable, columns, values,
	)
	//fmt.Println(tsql)
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

// backend access functions
func (ct *Controller) getSingleResultCheckQuery(ctx context.Context, query string) (int64, error) {
	var db *sql.DB
	db = ct.Database
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

	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
	)
	var rowID int64
	err = row.Scan(&rowID)
	if err != nil {
		return -1, err
	}
	return rowID, nil
}

func setPasswordToUser(ctx context.Context, ct *Controller, email string, password string) (int64, error) {
	var db *sql.DB
	db = ct.Database
	tsql :=
		`BEGIN 
			IF NOT EXISTS (SELECT * FROM DataSchema.UserDetails
						WHERE Email='%[1]s'
						)
				BEGIN 
					INSERT INTO DataSchema.UserDetails (Email, PasswordHash)
					VALUES ('%[1]s', '%[2]s')
				END 
			ELSE
				BEGIN
					UPDATE DataSchema.UserDetails 
					SET PasswordHash = '%[2]s'
					WHERE Email= '%[1]s'
				END
		END
		` // ref https://stackoverflow.com/questions/20971680/sql-server-insert-if-not-exists
	query := fmt.Sprintf(tsql, email, password)
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
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	return result.RowsAffected()
}

func isUserAuth(ctx context.Context, ct *Controller, email string, password string) (userDetails, error) {
	tsql :=
		`BEGIN 
				SELECT ID, Name, Email, Phone, AccessLevel, PasswordHash FROM DataSchema.UserDetails
				WHERE Email='%[1]s' and PasswordHash='%[2]s'
		END
		`
	var resultUser userDetails
	var db *sql.DB
	db = ct.Database
	query := fmt.Sprintf(tsql, email, password)
	var err error
	//check db is available
	if db == nil {
		err = errors.New(" db is null, transaction didn't go through")
		return resultUser, err
	}
	// Check if database is alive.
	err = db.PingContext(ctx)
	if err != nil {
		return resultUser, err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return resultUser, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(
		ctx,
	)
	var (
		tempName  sql.NullString
		tempPhone sql.NullString
		tempLevel sql.NullInt64
		tempHash  sql.NullString
	)
	err = row.Scan(&resultUser.ID, &tempName, &resultUser.Email,
		&tempPhone, &tempLevel, &tempHash)

	if tempName.Valid {
		resultUser.Name = tempName.String
	}
	if tempPhone.Valid {
		resultUser.Phone = tempPhone.String
	}
	if tempLevel.Valid {
		resultUser.AccessLevel = int(tempLevel.Int64)
	}
	if tempHash.Valid {
		resultUser.PasswordHash = tempHash.String
	}

	if err != nil {
		fmt.Println(err.Error())
		return resultUser, err
	}
	return resultUser, nil
}
