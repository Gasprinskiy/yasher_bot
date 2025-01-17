package db

import (
	"database/sql"
	sqldumps "yasher_bot/db/sql-dumps"

	_ "github.com/mattn/go-sqlite3"
)

type TestUser struct {
	chat_id   string
	user_id   int
	user_name string
}

func CreateDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	// 3 -4522854310 363693580 Ressurection21
	// 4 -4522854310 590546224 samgasper
	// 5 -4522854310 333209275 beastmodeonyeaa
	// 6 -4522854310 86131386 Kurliker
	// 7 -4522854310 153373681 FarMK1703
	// 8 -4522854310 761843289 k_abdullaev

	// -1002394430051

	// testUsers := []TestUser{
	// 	{
	// 		chat_id:   "-1002394430051",
	// 		user_id:   363693580,
	// 		user_name: "Ressurection21",
	// 	},
	// 	{
	// 		chat_id:   "-1002394430051",
	// 		user_id:   590546224,
	// 		user_name: "samgasper",
	// 	},
	// 	{
	// 		chat_id:   "-1002394430051",
	// 		user_id:   333209275,
	// 		user_name: "beastmodeonyeaa",
	// 	},
	// 	{
	// 		chat_id:   "-1002394430051",
	// 		user_id:   86131386,
	// 		user_name: "Kurliker",
	// 	},
	// 	{
	// 		chat_id:   "-1002394430051",
	// 		user_id:   153373681,
	// 		user_name: "FarMK1703",
	// 	},
	// 	{
	// 		chat_id:   "-1002394430051",
	// 		user_id:   761843289,
	// 		user_name: "k_abdullaev",
	// 	},
	// }

	db.Exec(sqldumps.CreateChatsSql)
	db.Exec(sqldumps.CreateUsersSql)
	db.Exec(sqldumps.CreateUsersScoreSql)

	// for _, testUser := range testUsers {
	// 	query := `INSERT INTO participants (chat_id, user_id, user_name) VALUES (?, ?, ?)`
	// 	db.Exec(
	// 		query,
	// 		testUser.chat_id,
	// 		testUser.user_id,
	// 		testUser.user_name,
	// 	)
	// }

	return db, nil
}
