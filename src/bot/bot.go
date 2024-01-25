package bot

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/i1mk8/WifiBot/ConfigManager"
	WifiManager "github.com/i1mk8/WifiBot/WiFiManager"
	"github.com/i1mk8/WifiBot/bot/states"
	"github.com/i1mk8/WifiBot/utils"
)

var bot *tgbotapi.BotAPI

func sendMessage(message tgbotapi.MessageConfig) {
	_, err := bot.Send(message)
	if err != nil {
		log.Panic(err)
	}
}

func StartBot() {
	var err error
	bot, err = tgbotapi.NewBotAPI(ConfigManager.GetConfig().BotToken)

	if err != nil {
		log.Panic(err)
	}
	log.Printf("Авторизован как @%s", bot.Self.UserName)

	handleNewMessages()
}

func handleNewMessages() {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	messages := bot.GetUpdatesChan(update)
	for message := range messages {
		if message.Message != nil {

			state, _ := states.GetUser(message.Message.From.ID)
			if state != nil {

				switch state.State {

				case states.MainMenu:
					handleMainMenu(message)

				case states.ScheduleMenu:
					handleScheduleMenu(message)

				case states.EditScheduleDown:
					handleEditScheduleDown(message)

				case states.EditScheduleUp:
					handleEditScheduleUp(message)
				}

			} else {
				config := ConfigManager.GetConfig()
				if utils.Int64InSlice(config.Users, message.Message.From.ID) {
					sendMainMenu(message.Message.Chat.ID)
				}
			}
		}
	}
}

func sendMainMenu(userId int64) {
	states.SetUser(userId, states.MainMenu)

	message := tgbotapi.NewMessage(userId, "Главное меню")
	message.ReplyMarkup = MainMenuKeyboard

	sendMessage(message)
}

func handleMainMenu(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)
	newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "")
	switch messageText {

	case "вкл":
		WifiManager.InterfacesUp()
		newMessage.Text = "Успешно!"

	case "выкл":
		WifiManager.InterfacesDown()
		newMessage.Text = "Успешно!"

	case "расписание":
		sendScheduleMenu(message.Message.From.ID)
		return

	default:
		return
	}

	sendMessage(newMessage)
}

func getScheduleStatus() string {
	config := ConfigManager.GetConfig()

	var status string
	if config.ScheduleEnabled {
		status = "Включено"
	} else {
		status = "Выключено"
	}

	scheduleDownHour := utils.IntToString(config.ScheduleDownHour)
	scheduleDownMinute := utils.IntToString(config.ScheduleDownMinute)

	scheduleUpHour := utils.IntToString(config.ScheduleUpHour)
	scheduleUpMinute := utils.IntToString(config.ScheduleUpMinute)

	return fmt.Sprintf("Статус: %s\nВЫКЛ: %s:%s\nВКЛ: %s:%s", status, scheduleDownHour, scheduleDownMinute, scheduleUpHour, scheduleUpMinute)
}

func sendScheduleMenu(userId int64) {
	states.SetUser(userId, states.ScheduleMenu)

	message := tgbotapi.NewMessage(userId, "Управление расписанием")
	message.ReplyMarkup = ScheduleKeyboard

	sendMessage(message)
}

func handleScheduleMenu(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)
	newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "")
	switch messageText {

	case "вкл":
		config := ConfigManager.GetConfig()
		config.ScheduleEnabled = true
		ConfigManager.SetConfig(config)

		newMessage.Text = "Успешно!"

	case "выкл":
		config := ConfigManager.GetConfig()
		config.ScheduleEnabled = false
		ConfigManager.SetConfig(config)

		newMessage.Text = "Успешно!"

	case "просмотр":
		newMessage.Text = getScheduleStatus()

	case "изменить":
		statusMessage := tgbotapi.NewMessage(message.Message.Chat.ID, getScheduleStatus())
		sendMessage(statusMessage)

		states.SetUser(message.Message.From.ID, states.EditScheduleDown)
		newMessage.Text = "Введите время выключения\nФормат: HH:MM"
		newMessage.ReplyMarkup = CancelKeyboard

	case "назад":
		sendMainMenu(message.Message.Chat.ID)
		return

	default:
		return
	}

	sendMessage(newMessage)
}

func parseEditScheduleMessage(text string) *[2]int {
	data := strings.Split(text, ":")

	if len(data) == 2 {
		result := make([]int, 2)

		for index, value := range data {
			intValue, err := strconv.Atoi(value)

			if err != nil {
				return nil
			}

			result[index] = intValue
		}

		return (*[2]int)(result)
	}

	return nil
}

func handleEditScheduleDown(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)

	if messageText == "отмена" {
		sendScheduleMenu(message.Message.Chat.ID)

	} else {
		newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "")
		data := parseEditScheduleMessage(messageText)

		if data != nil {
			config := ConfigManager.GetConfig()
			config.ScheduleDownHour = data[0]
			config.ScheduleDownMinute = data[1]
			ConfigManager.SetConfig(config)

			infoMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "Успешно!")
			sendMessage(infoMessage)

			states.SetUser(message.Message.From.ID, states.EditScheduleUp)
			newMessage.Text = "Введите время включения\nФормат: HH:MM"
			newMessage.ReplyMarkup = CancelKeyboard

		} else {
			newMessage.Text = "Неверный формат!"
		}

		sendMessage(newMessage)
	}
}

func handleEditScheduleUp(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)

	if messageText == "отмена" {
		sendScheduleMenu(message.Message.Chat.ID)

	} else {
		data := parseEditScheduleMessage(messageText)

		if data != nil {
			config := ConfigManager.GetConfig()
			config.ScheduleUpHour = data[0]
			config.ScheduleUpMinute = data[1]
			ConfigManager.SetConfig(config)

			infoMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "Успешно!")
			sendMessage(infoMessage)

			sendScheduleMenu(message.Message.Chat.ID)

		} else {
			newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "Неверный формат!")
			sendMessage(newMessage)
		}
	}
}
