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
			layer.CheckBotStarted(
				update,
				func(update tgbotapi.Update) {
					layer.HandleRegister(update)
				},
			)

		case fmt.Sprintf("%s@%s", commands.RunTheGameCommand, bot.Self.UserName):
			layer.CheckBotStarted(
				update,
				func(update tgbotapi.Update) {
					layer.HandleRunTheGame(update)
				},
			)

		case fmt.Sprintf("%s@%s", commands.GetTopWinners, bot.Self.UserName):
			layer.CheckBotStarted(
				update,
				func(update tgbotapi.Update) {
					layer.HandleTopWinners(update)
				},
			)

		case fmt.Sprintf("%s@%s", commands.HealthCheck, bot.Self.UserName):
			layer.HandleHealthCheck(update)
		}
		fmt.Println("update: ", update.Message.Text)
	}
}

func (e *External) HandleStart(update tgbotapi.Update) {
	message := e.usecase.GetStarted(e.getChatIdAsString(update))
	if message == "" {
		return
	}

	e.handleMessageSend(update, message)
}

func (e *External) HandleRegister(update tgbotapi.Update) {
	param := chat.AddParticipantParam{
		ChatID:   e.getChatIdAsString(update),
		UserID:   update.Message.From.ID,
		UserName: update.Message.From.UserName,
	}

	message := e.usecase.RegisterParticipants(param)
	if message == "" {
		return
	}

	e.handleMessageSend(update, message)
}

func (e *External) HandleRunTheGame(update tgbotapi.Update) {
	chatId := e.getChatIdAsString(update)

	participants, checkParticipantsMessage := e.usecase.CheckParticipantsMessage(chatId)
	if checkParticipantsMessage != "" {
		e.handleMessageSend(update, checkParticipantsMessage)
		return
	}

	winnerFoundMessage := e.usecase.TodayWinnerFoundMessage(chatId)
	if winnerFoundMessage != "" {
		e.handleMessageSend(update, winnerFoundMessage)
		return
	}

	beforeRunMessage := e.usecase.GetBeforeRunMessage()
	e.handleMessageSend(update, beforeRunMessage)

	go func() {
		time.Sleep(800 * time.Millisecond)

		message := e.usecase.RunTheGame(participants, chatId)
		if message == "" {
			return
		}

		e.handleMessageSend(update, message)
	}()
}

func (e *External) HandleTopWinners(update tgbotapi.Update) {
	message := e.usecase.GetTopWinners(e.getChatIdAsString(update))
	if message == "" {
		return
	}

	e.handleMessageSend(update, message)
}

func (e *External) CheckBotStarted(
	update tgbotapi.Update,
	onStartedFunc func(update tgbotapi.Update),
) {
	startedMessage := e.usecase.IsBotStartedMessage(e.getChatIdAsString(update))
	if startedMessage != "" {
		e.handleMessageSend(update, startedMessage)
		return
	}

	onStartedFunc(update)
}

func (e *External) HandleHealthCheck(update tgbotapi.Update) {
	message := e.usecase.GetHealthCheckMessage()
	e.handleMessageSend(update, message)
}

func (e *External) getChatIdAsString(update tgbotapi.Update) string {
	return fmt.Sprintf("%d", update.Message.Chat.ID)
}

func (e *External) handleMessageSend(update tgbotapi.Update, message string) {
	botMsg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	botMsg.ParseMode = "HTML"

	_, err := e.bot.Send(botMsg)
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения: ", err.Error())
	}
}
