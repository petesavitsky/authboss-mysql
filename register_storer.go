package authbossmysql

import (
	"database/sql"
	"time"

	"github.com/volatiletech/authboss"
)

const createUserStatement = "INSERT INTO user(email, password, created, updated) values (?,?,?,?);"

type RegisterStorer struct {
	userStorer      UserStorer
	createStatement *sql.Stmt
}

func (registerStorer RegisterStorer) Create(key string, attr authboss.Attributes) error {
	email, _ := attr.String(authboss.StoreEmail)
	password, _ := attr.String(authboss.StorePassword)
	created := time.Now()
	updated := time.Now()
	_, err := registerStorer.createStatement.Exec(email, password, created, updated)
	if err != nil {
		return authboss.ErrUserFound
	}
	return nil
}

func (registerStorer RegisterStorer) Put(key string, attr authboss.Attributes) error {
	return registerStorer.userStorer.Put(key, attr)
}

func (registerStorer RegisterStorer) Get(key string) (interface{}, error) {
	return registerStorer.userStorer.Get(key)
}

func (registerStorer RegisterStorer) Close() {
	defer registerStorer.createStatement.Close()
	defer registerStorer.userStorer.getStatement.Close()
	defer registerStorer.userStorer.updateStatement.Close()
}
