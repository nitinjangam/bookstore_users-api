package users

import (
	"fmt"

	usersdb "github.com/nitinjangam/bookstore_users-api/datasources/mysql/users_db"
	"github.com/nitinjangam/bookstore_users-api/logger"
	"github.com/nitinjangam/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail      = "email_UNIQUE"
	noRowsInResult        = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

var userDB = make(map[int64]*User)
var emptyUser *User

//Get function
func (user *User) Get() *errors.RestErr {
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error while trying to prepare GetUser statement", err)
		return errors.CreateNewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("Error while trying to get user", err)
		return errors.CreateNewInternalServerError("database error")
	}
	return nil
}

//Save function
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error while trying to prepare save user statement", err)
		return errors.CreateNewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("Error while trying to save user", err)
		return errors.CreateNewInternalServerError("database error")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error while trying to get last insert id", err)
		return errors.CreateNewInternalServerError("database error")
	}

	user.ID = userID
	return nil
}

//Update function
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error while trying to prepare update user statement", err)
		return errors.CreateNewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("Error while trying to update user", err)
		return errors.CreateNewInternalServerError("database error")
	}
	return nil
}

//Delete function
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error while trying to prepare delete user statement", err)
		return errors.CreateNewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		logger.Error("Error while trying to delete user", err)
		return errors.CreateNewInternalServerError("database error")
	}
	return nil
}

//FindByStatus function
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("Error while trying to prepare find user by status statement", err)
		return nil, errors.CreateNewInternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("Error while trying to find user by status", err)
		return nil, errors.CreateNewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error while trying to get info from rows in query find user by status", err)
			return nil, errors.CreateNewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.CreateNewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
