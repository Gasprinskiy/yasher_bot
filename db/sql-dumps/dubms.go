package sqldumps

const CreateChatsSql = `CREATE TABLE IF NOT EXISTS chats (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	chat_id TEXT NOT NULL,
	last_run DATETIME
);`

const CreateUsersSql = `CREATE TABLE IF NOT EXISTS participants (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	chat_id TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	user_name TEXT NOT NULL,
	FOREIGN KEY (chat_id) REFERENCES chats(chat_id) ON DELETE CASCADE ON UPDATE CASCADE
);`

const CreateUsersScoreSql = `CREATE TABLE IF NOT EXISTS participants_score (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	chat_id TEXT NOT NULL,
	participant_id INTEGER NOT NULL,
	score_count INTEGER NOT NULL,
	is_last_winner BOOLEAN NOT NULL DEFAULT 0,
	FOREIGN KEY (participant_id) REFERENCES participants(id) ON DELETE CASCADE ON UPDATE CASCADE
);`
