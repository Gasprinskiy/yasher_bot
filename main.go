package main

import (
	"log"
	"yasher_bot/db"
	"yasher_bot/external"
	"yasher_bot/internal/repository/sqllite"
	"yasher_bot/internal/usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7660939129:AAH3lPmRL5b6yeCPQZH35MTGCK0I3I4Sx0k")
	if err != nil {
		log.Panic(err)
	}

	db, err := db.CreateDb()
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	chatRepo := sqllite.NewChatRepository(db)
	gameUsecase := usecase.NewGameUsecase(chatRepo)

	external.StartUpdatesListening(bot, gameUsecase)
}
