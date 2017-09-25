package authbossmysql

import (
	"database/sql"
	"time"

	"github.com/volatiletech/authboss"
)

const emailField = "email"
const passwordField = "password"
const createdField = "created"
const updatedField = "updated"
const updateUserStatement = "UPDATE user SET email = ?, password = ?, updated = ? where email = ?;"
const getUserStatement = "SELECT email, password, created, updated from user where email = ?;"

type User struct {
	email    string
	password string
	created  time.Time
	updated  time.Time
}

type UserStorer struct {
	updateStatement *sql.Stmt
	getStatement    *sql.Stmt
}

func (storer UserStorer) Put(key string, attr authboss.Attributes) error {
	// TODO: handle cases where email or password are not passed in
	email, _ := attr.String(authboss.StoreEmail)
	password, _ := attr.String(authboss.StorePassword)
	updated := time.Now()
	result, err := storer.updateStatement.Exec(email, password, updated)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil && rowsAffected < 1 {
		return authboss.ErrUserNotFound
	}
	return nil
}

func (storer UserStorer) Get(key string) (interface{}, error) {
	row := storer.getStatement.QueryRow(key)
	user := User{}
	err := row.Scan(&user)
	if err != nil && err == sql.ErrNoRows {
		return nil, authboss.ErrUserNotFound
	} else if err == nil {
		return user, nil
	}
	return nil, err
}

func (storer UserStorer) Close() {
	defer storer.getStatement.Close()
	defer storer.updateStatement.Close()
}
