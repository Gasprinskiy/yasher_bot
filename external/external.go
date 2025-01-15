package external

import (
	"fmt"
	"log"
	"time"
	"yasher_bot/constants/commands"
	"yasher_bot/entity/chat"
	"yasher_bot/internal/usecase"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type External struct {
	bot     *tgbotapi.BotAPI
	usecase *usecase.GameUsecase
}

func StartUpdatesListening(
	bot *tgbotapi.BotAPI,
	usecase *usecase.GameUsecase,
) {
	layer := External{
		bot,
		usecase,
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		switch update.Message.Text {
		case fmt.Sprintf("%s@%s", commands.StartCommand, bot.Self.UserName):
			layer.HandleStart(update)

		case fmt.Sprintf("%s@%s", commands.RegisterCommand, bot.Self.UserName):
			layer.HandleRegister(update)

		case fmt.Sprintf("%s@%s", commands.RunTheGameCommand, bot.Self.UserName):
			layer.HandleRunTheGame(update)
		}
		fmt.Println("update: ", update.Message.Text)
	}
}

func (e *External) HandleStart(update tgbotapi.Update) {
	chatId := fmt.Sprintf("%d", update.Message.Chat.ID)

	message := e.usecase.GetStarted(chatId)
	if message == "" {
		return
	}

	botMsg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	botMsg.ParseMode = "HTML"

	e.bot.Send(botMsg)
}

func (e *External) HandleRegister(update tgbotapi.Update) {
	param := chat.AddParticipantParam{
		ChatID:   fmt.Sprintf("%d", update.Message.Chat.ID),
		UserID:   update.Message.From.ID,
		UserName: update.Message.From.UserName,
	}

	message := e.usecase.RegisterParticipants(param)
	if message == "" {
		return
	}

	botMsg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	botMsg.ParseMode = "HTML"

	e.bot.Send(botMsg)
}

func (e *External) HandleRunTheGame(update tgbotapi.Update) {
	chatId := fmt.Sprintf("%d", update.Message.Chat.ID)

	participants, checkParticipantsMessage := e.usecase.CheckParticipantsMessage(chatId)
	if checkParticipantsMessage != "" {
		botMsg := tgbotapi.NewMessage(update.Message.Chat.ID, checkParticipantsMessage)
		botMsg.ParseMode = "HTML"

		e.bot.Send(botMsg)
		return
	}

	winnerFoundMessage := e.usecase.TodayWinnerFoundMessage(chatId)
	if winnerFoundMessage != "" {
		botMsg := tgbotapi.NewMessage(update.Message.Chat.ID, winnerFoundMessage)
		botMsg.ParseMode = "HTML"

		e.bot.Send(botMsg)
		return
	}

	beforeRunMessage := e.usecase.GetBeforeRunMessage()
	botMsg := tgbotapi.NewMessage(update.Message.Chat.ID, beforeRunMessage)
	botMsg.ParseMode = "HTML"

	e.bot.Send(botMsg)

	go func() {
		time.Sleep(800 * time.Millisecond)

		message := e.usecase.RunTheGame(participants, chatId)
		if message == "" {
			return
		}

		botMsg = tgbotapi.NewMessage(update.Message.Chat.ID, message)

		e.bot.Send(botMsg)
	}()
}
