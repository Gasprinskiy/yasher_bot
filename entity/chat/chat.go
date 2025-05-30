package chat

import (
	"time"
	"yasher_bot/tools/sqlnull"
)

type Chat struct {
	ID      int
	ChatID  string
	LastRun *time.Time
}

type Participant struct {
	ID       int
	ChatID   string
	UserID   int
	UserName string
}

type AddParticipantParam struct {
	ChatID   string
	UserID   int
	UserName string
}

type Score struct {
	ChatID        string
	ParticipantID int
	ScoreCount    int
}

type ParticipantWithScore struct {
	ID           int
	ChatID       string
	UserID       int
	UserName     string
	ScoreCount   sqlnull.NullInt64
	IsLastWinner bool
}
