package authbossmysql

import (
	"database/sql"
	"log"
	// imports the mysql package for database creation
	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

func GetDB(dataSourceName string) (*sql.DB, error) {
	return sql.Open(driverName, dataSourceName)
}

func GetUserStorer(db *sql.DB) UserStorer {
	updateUser, err := db.Prepare(updateUserStatement)
	if err != nil {
		log.Fatalf("Failed to create update user statement [%v]", err)
	}
	getUser, err := db.Prepare(getUserStatement)
	if err != nil {
		log.Fatalf("Failed to create get user statement [%v]", err)
	}
	return UserStorer{updateStatement: updateUser, getStatement: getUser}
}

func GetRegisterStorer(db *sql.DB) RegisterStorer {
	userStorer := GetUserStorer(db)
	createUser, err := db.Prepare(createUserStatement)
	if err != nil {
		log.Fatalf("Failed to create create user statement [%v]", err)
	}
	return RegisterStorer{userStorer: userStorer, createStatement: createUser}
}
