package main

import (
	"log"
	"os"
	"yasher_bot/db"
	"yasher_bot/external"
	"yasher_bot/internal/repository/sqllite"
	"yasher_bot/internal/usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Ошибка загрузки .env файла: ", err)
	}

	botToken := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic("Ошибка при подключении к API бота: ", err)
	}

	db, err := db.CreateDb()
	if err != nil {
		log.Panic("Ошибка при создании БД: ", err)
	}

	defer db.Close()

	chatRepo := sqllite.NewChatRepository(db)
	gameUsecase := usecase.NewGameUsecase(chatRepo)

	external.StartUpdatesListening(bot, gameUsecase)
}
