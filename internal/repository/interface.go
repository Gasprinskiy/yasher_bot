package repository

import (
	"time"
	"yasher_bot/entity/chat"
)

type Chat interface {
	AddNewChat(chatID string) error
	GetChatById(chatID string) (chat.Chat, error)
	SetChatLastRun(chatID string, lastRun time.Time) error
	AddGameParticipants(participant chat.AddParticipantParam) error
	GetChatParticipant(chatID string, userId int) (chat.Participant, error)
	FindChatParticipants(chatID string) ([]chat.Participant, error)
	FindUserScoreById(chatID string, userID int) (int, error)
	SetNewUserScore(score chat.Score) error
	UnMarkLastWinner(chatID string) error
	UpdateUserScore(score chat.Score) error
	FindLastWinner(chatID string) (chat.Participant, error)
	FindChatParticipantsWithScore(chatID string) ([]chat.ParticipantWithScore, error)
}
