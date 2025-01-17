package usecase

import (
	"fmt"
	"time"
	"yasher_bot/constants/global"
	"yasher_bot/constants/messages"
	"yasher_bot/entity/chat"
	"yasher_bot/internal/repository"
	"yasher_bot/tools/chronos"
	"yasher_bot/tools/random"
	"yasher_bot/tools/slice"
)

type GameUsecase struct {
	repo repository.Chat
}

func NewGameUsecase(repo repository.Chat) *GameUsecase {
	return &GameUsecase{repo}
}

func (u *GameUsecase) GetStarted(chatId string) string {
	var message string

	_, err := u.repo.GetChatById(chatId)
	if err != nil && err != global.ErrNoData {
		fmt.Println("Ошибка при получении чата: ", err.Error())
		return message
	}

	if err == global.ErrNoData {
		err := u.repo.AddNewChat(chatId)
		if err != nil {
			fmt.Println("Ошибка при создании чата: ", err.Error())
			return message
		}

		message = messages.HelloMessage
	} else {
		message = messages.AlreadyStartedMessage
	}

	return message
}

func (u *GameUsecase) IsBotStartedMessage(chatId string) string {
	_, err := u.repo.GetChatById(chatId)
	if err != nil && err != global.ErrNoData {
		fmt.Println("Ошибка при получении чата: ", err.Error())
		return ""
	}

	if err == global.ErrNoData {
		return messages.BotIsNotStarted
	}

	return ""
}

func (u *GameUsecase) RegisterParticipants(participant chat.AddParticipantParam) string {
	var message string

	_, err := u.repo.GetChatParticipant(participant.ChatID, participant.UserID)
	if err != nil && err != global.ErrNoData {
		fmt.Println("Ошибка при получении учатсника: ", err.Error())
		return message
	}

	if err == global.ErrNoData {
		err := u.repo.AddGameParticipants(participant)
		if err != nil {
			fmt.Println("Ошибка при добавлении участника: ", err.Error())
			return message
		}

		message = fmt.Sprintf(messages.RegisteredMessage, participant.UserName)
	} else {
		message = fmt.Sprintf(messages.AlreadyRegisteredMessage, participant.UserName)
	}

	return message
}

func (u *GameUsecase) TodayWinnerFoundMessage(chatId string) string {
	var message string

	chat, err := u.repo.GetChatById(chatId)
	if err != nil {
		fmt.Println("Ошибка при получении данных по чату: ", err.Error())
		return message
	}

	if chat.LastRun != nil && chronos.IsToday(*chat.LastRun) {
		lastWinner, err := u.repo.FindLastWinner(chatId)
		if err != nil {
			fmt.Println("Ошибка при получении последнего победтиеля по чату: ", err.Error())
			return message
		}

		messageIndex := random.MakeRandomNumber(len(messages.WinnerAlreadyFoundMessages))
		message = fmt.Sprintf(messages.WinnerAlreadyFoundMessages[messageIndex], lastWinner.UserName)
	}

	return message
}

func (u *GameUsecase) GetBeforeRunMessage() string {
	messageIndex := random.MakeRandomNumber(len(messages.SearchInProgressMessages))

	return messages.SearchInProgressMessages[messageIndex]
}

func (u *GameUsecase) CheckParticipantsMessage(chatId string) ([]chat.ParticipantWithScore, string) {
	data, err := u.repo.FindChatParticipants(chatId)
	if err != nil && err != global.ErrNoData {
		fmt.Println("Ошибка при получении учатсников: ", err.Error())
		return data, ""
	}

	if err == global.ErrNoData {
		return data, messages.NoParticipantsMessage
	}

	if len(data) == 1 {
		return data, fmt.Sprintf(messages.TooFewParticipantsMessage, data[0].UserName)
	}

	return data, ""
}

func (u *GameUsecase) RunTheGame(participants []chat.ParticipantWithScore, chatId string) string {
	data := []chat.ParticipantWithScore{}
	for _, participant := range participants {
		if participant.IsLastWinner && (participant.ScoreCount.Valid && participant.ScoreCount.GetInt() >= 3) {
			continue
		}
		data = append(data, participant)
	}

	data = slice.ShuffleArray(data)

	winnerIndex := random.MakeRandomNumber(len(data))
	winner := data[winnerIndex]

	err := u.repo.UnMarkLastWinner(chatId)
	if err != nil {
		fmt.Println("Ошибка при обновлении последнего победителя: ", err.Error())
		return ""
	}

	newScore := chat.Score{
		ChatID:        chatId,
		ParticipantID: winner.ID,
		ScoreCount:    1 + winner.ScoreCount.GetInt(),
	}

	if winner.ScoreCount.Valid {
		err = u.repo.UpdateUserScore(newScore)
		if err != nil {
			fmt.Println("Ошибка при обновлении cчета победителя: ", err.Error())
			return ""
		}
	} else {
		err = u.repo.SetNewUserScore(newScore)
		if err != nil {
			fmt.Println("Ошибка при создании cчета победителя: ", err.Error())
			return ""
		}
	}

	lastRun := time.Now()
	err = u.repo.SetChatLastRun(chatId, lastRun)
	if err != nil {
		fmt.Println("Ошибка при обновлении последней даты запуска игры: ", err.Error())
		return ""
	}

	messageIndex := random.MakeRandomNumber(len(messages.WinnerMessages))

	return fmt.Sprintf(messages.WinnerMessages[messageIndex], winner.UserName)
}

func (u *GameUsecase) GetGameParticipantsListMessage(chatId string) string {
	data, err := u.repo.GetChatParticipantList(chatId)
	if err != nil && err != global.ErrNoData {
		fmt.Println("Ошибка при получении участников игры: ", err.Error())
		return ""
	}

	if err == global.ErrNoData {
		return messages.NoParticipantsMessage
	}

	message := messages.ParticipantsListMessage

	for i, participant := range data {
		message += fmt.Sprintf("%d. %s\n", i+1, participant.UserName)
	}

	message += messages.ParticipantsListMessageEnd

	return message
}

func (u *GameUsecase) GetTopWinners(chatId string) string {
	data, err := u.repo.FindChatParticipantsWithScore(chatId)
	if err != nil && err != global.ErrNoData {
		fmt.Println("Ошибка при получении топа победителей: ", err.Error())
		return ""
	}

	if err == global.ErrNoData {
		return messages.TopWinnersEmpty
	}

	message := messages.TopWinnersMessage

	for i, participant := range data {
		message += fmt.Sprintf("%d. %s: %d раз(а)\n", i+1, participant.UserName, participant.ScoreCount.GetInt())
	}

	return message
}

func (u *GameUsecase) GetHealthCheckMessage() string {
	messageIndex := random.MakeRandomNumber(len(messages.HealthCheckMessages))
	return messages.HealthCheckMessages[messageIndex]
}
