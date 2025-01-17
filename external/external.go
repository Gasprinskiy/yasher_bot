package external

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"yasher_bot/constants/commands"
	"yasher_bot/entity/chat"
	"yasher_bot/internal/usecase"
	"yasher_bot/tools/random"

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

	botUserName := bot.Self.UserName

	for update := range updates {
		switch update.Message.Text {
		case fmt.Sprintf("%s@%s", commands.StartCommand, botUserName):
			layer.HandleStart(update)

		case fmt.Sprintf("%s@%s", commands.RegisterCommand, botUserName):
			layer.CheckBotStarted(
				update,
				func(update tgbotapi.Update) {
					layer.HandleRegister(update)
				},
			)

		case fmt.Sprintf("%s@%s", commands.RunTheGameCommand, botUserName):
			layer.CheckBotStarted(
				update,
				func(update tgbotapi.Update) {
					layer.HandleRunTheGame(update)
				},
			)

		case fmt.Sprintf("%s@%s", commands.GetTopWinners, botUserName):
			layer.CheckBotStarted(
				update,
				func(update tgbotapi.Update) {
					layer.HandleTopWinners(update)
				},
			)

		case fmt.Sprintf("%s@%s", commands.GetParticipantsList, botUserName):
			layer.HandleGetParticipantsList(update)

		case fmt.Sprintf("%s@%s", commands.HealthCheck, botUserName):
			layer.HandleHealthCheck(update)
		}

		// Debug code
		if strings.Contains(update.Message.Text, fmt.Sprintf("%s@%s", "/rand", botUserName)) {
			splitedMessage := strings.Split(update.Message.Text, " ")[1]
			i, err := strconv.Atoi(splitedMessage)
			if err != nil {
				fmt.Println("Could not parse string number: ", err)
			}
			layer.handleMessageSend(update, fmt.Sprintf("Shit: %d", random.MakeRandomNumber(i)))
		}
		//

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

func (e *External) HandleGetParticipantsList(update tgbotapi.Update) {
	message := e.usecase.GetGameParticipantsListMessage(e.getChatIdAsString(update))
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
