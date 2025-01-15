package db

import (
	"database/sql"
	sqldumps "yasher_bot/db/sql-dumps"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	db.Exec(sqldumps.CreateChatsSql)
	db.Exec(sqldumps.CreateUsersSql)
	db.Exec(sqldumps.CreateUsersScoreSql)

	return db, nil
}
