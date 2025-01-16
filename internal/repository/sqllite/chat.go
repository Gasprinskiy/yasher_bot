package sqllite

import (
	"database/sql"
	"fmt"
	"time"
	"yasher_bot/constants/global"
	"yasher_bot/entity/chat"
	"yasher_bot/internal/repository"
)

type chatRepository struct {
	db *sql.DB
}

func NewChatRepository(db *sql.DB) repository.Chat {
	return &chatRepository{db}
}

func (r *chatRepository) AddNewChat(chatID string) error {
	query := `INSERT INTO chats (chat_id) VALUES (?)`
	_, err := r.db.Exec(query, chatID)
	return err
}

func (r *chatRepository) GetChatById(chatID string) (chat.Chat, error) {
	data := chat.Chat{}
	found := false

	query := `
	SELECT 
		c.id, 
		c.chat_id,
		c.last_run
	FROM chats c
		WHERE c.chat_id = ?`

	rows, err := r.db.Query(query, chatID)

	if err != nil {
		return data, err
	}

	for rows.Next() {
		found = true

		var id int
		var chat_id string
		var last_run *time.Time
		err = rows.Scan(&id, &chat_id, &last_run)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}
		data.ID = id
		data.ChatID = chat_id
		data.LastRun = last_run
	}

	if !found {
		return data, global.ErrNoData
	}

	return data, nil
}

func (r *chatRepository) SetChatLastRun(chatID string, lastRun time.Time) error {
	query := `UPDATE chats SET last_run = ? WHERE chat_id = ?`
	_, err := r.db.Exec(query, lastRun, chatID)
	return err
}

func (r *chatRepository) AddGameParticipants(participant chat.AddParticipantParam) error {
	query := `INSERT INTO participants (chat_id, user_id, user_name) VALUES (?, ?, ?)`
	_, err := r.db.Exec(
		query,
		participant.ChatID,
		participant.UserID,
		participant.UserName,
	)
	return err
}

func (r *chatRepository) GetChatParticipant(chatID string, userId int) (chat.Participant, error) {
	data := chat.Participant{}
	found := false

	query := `
	SELECT 
		p.id, 
		p.chat_id,
		p.user_id,
		p.user_name
	FROM participants p
		WHERE p.chat_id = ?
		AND p.user_id = ?`

	rows, err := r.db.Query(query, chatID, userId)

	if err != nil {
		return data, err
	}

	for rows.Next() {
		found = true

		var id int
		var chat_id string
		var user_id int
		var user_name string
		err = rows.Scan(&id, &chat_id, &user_id, &user_name)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}
		data.ID = id
		data.ChatID = chat_id
		data.UserID = user_id
		data.UserName = user_name
	}

	if !found {
		return data, global.ErrNoData
	}

	return data, nil
}

func (r *chatRepository) GetChatParticipantList(chatID string) ([]chat.Participant, error) {
	data := []chat.Participant{}

	query := `
	SELECT 
		p.id, 
		p.chat_id,
		p.user_id,
		p.user_name
	FROM participants p
		WHERE p.chat_id = ?`

	rows, err := r.db.Query(query, chatID)

	if err != nil {
		return data, err
	}

	for rows.Next() {
		var id int
		var chat_id string
		var user_id int
		var user_name string
		err = rows.Scan(&id, &chat_id, &user_id, &user_name)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}

		participant := chat.Participant{
			ID:       id,
			ChatID:   chat_id,
			UserID:   user_id,
			UserName: user_name,
		}
		data = append(data, participant)
	}

	if len(data) == 0 {
		return data, global.ErrNoData
	}

	return data, nil
}

func (r *chatRepository) FindChatParticipants(chatID string) ([]chat.Participant, error) {
	data := []chat.Participant{}

	query := `
	SELECT 
		p.id, 
		p.chat_id,
		p.user_id,
		p.user_name
	FROM participants p
		WHERE p.chat_id = ?`

	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		var id int
		var chat_id string
		var user_id int
		var user_name string
		err = rows.Scan(&id, &chat_id, &user_id, &user_name)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}

		participant := chat.Participant{
			ID:       id,
			ChatID:   chat_id,
			UserID:   user_id,
			UserName: user_name,
		}
		data = append(data, participant)
	}

	if len(data) == 0 {
		return data, global.ErrNoData
	}

	return data, nil
}

func (r *chatRepository) FindUserScoreById(chatID string, userID int) (int, error) {
	var score_count int
	var found bool

	query := `
	SELECT 
		ps.score_count
	FROM participants_score ps
		WHERE ps.chat_id = ?
		AND ps.participant_id = ?`

	rows, err := r.db.Query(query, chatID, userID)
	if err != nil {
		return score_count, err
	}

	for rows.Next() {
		found = true
		err = rows.Scan(&score_count)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}
	}

	if !found {
		return score_count, global.ErrNoData
	}

	return score_count, nil
}

func (r *chatRepository) SetNewUserScore(score chat.Score) error {
	query := `INSERT INTO participants_score (chat_id, participant_id, score_count, is_last_winner) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(
		query,
		score.ChatID,
		score.ParticipantID,
		score.ScoreCount,
		1,
	)
	return err
}

func (r *chatRepository) UnMarkLastWinner(chatID string) error {
	query := `UPDATE participants_score SET is_last_winner = 0 WHERE chat_id = ?`
	_, err := r.db.Exec(query, chatID)
	return err
}

func (r *chatRepository) UpdateUserScore(score chat.Score) error {
	query := `
	UPDATE 
		participants_score 
	SET 
		score_count = ?,
		is_last_winner = 1
	WHERE chat_id = ?
	AND participant_id = ?`
	_, err := r.db.Exec(query, score.ScoreCount, score.ChatID, score.ParticipantID)
	return err
}

func (r *chatRepository) FindLastWinner(chatID string) (chat.Participant, error) {
	data := chat.Participant{}
	found := false

	query := `
	SELECT 
		p.id, 
		p.chat_id,
		p.user_id,
		p.user_name
	FROM participants_score ps
		JOIN participants p ON p.id = ps.participant_id
	WHERE p.chat_id = ?
	AND ps.is_last_winner = 1`

	rows, err := r.db.Query(query, chatID)

	if err != nil {
		return data, err
	}

	for rows.Next() {
		found = true

		var id int
		var chat_id string
		var user_id int
		var user_name string
		err = rows.Scan(&id, &chat_id, &user_id, &user_name)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}
		data.ID = id
		data.ChatID = chat_id
		data.UserID = user_id
		data.UserName = user_name
	}

	if !found {
		return data, global.ErrNoData
	}

	return data, nil
}

func (r *chatRepository) FindChatParticipantsWithScore(chatID string) ([]chat.ParticipantWithScore, error) {
	data := []chat.ParticipantWithScore{}

	query := `
	SELECT 
		p.id, 
		p.chat_id,
		p.user_id,
		p.user_name,
		ps.score_count
	FROM participants p
		JOIN participants_score ps ON ps.participant_id = p.id 
	WHERE p.chat_id = ?
	ORDER BY ps.score_count DESC
	LIMIT 10
	OFFSET 0`

	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return data, err
	}

	for rows.Next() {
		var id int
		var chat_id string
		var user_id int
		var user_name string
		var score_count int

		err = rows.Scan(&id, &chat_id, &user_id, &user_name, &score_count)
		if err != nil {
			fmt.Printf("Ошибка при чтении строки: %v", err)
		}

		participant := chat.ParticipantWithScore{
			ID:         id,
			ChatID:     chat_id,
			UserID:     user_id,
			UserName:   user_name,
			ScoreCount: score_count,
		}
		data = append(data, participant)
	}

	if len(data) == 0 {
		return data, global.ErrNoData
	}

	return data, nil
}
