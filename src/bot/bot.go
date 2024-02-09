// Код для получения/отправки сообщений в telegram
package bot

import (
	"crypto/tls"
	"crypto/x509"
	"embed"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/i1mk8/WifiBot/ConfigManager"
	WifiManager "github.com/i1mk8/WifiBot/WiFiManager"
	"github.com/i1mk8/WifiBot/bot/states"
	"github.com/i1mk8/WifiBot/utils"
)

const (
	CertPath = "cert.pem"
)

var (
	//go:embed cert.pem
	fs  embed.FS
	bot *tgbotapi.BotAPI
)

// Запуск бота
func StartBot() {

	// Сертификат необходим для подключения к telegram по https (на роутере сертификат отсутсвует)
	cert, err := fs.ReadFile(CertPath)
	if err != nil {
		log.Panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		},
	}

	bot, err = tgbotapi.NewBotAPIWithClient(ConfigManager.GetConfig().BotToken, tgbotapi.APIEndpoint, httpClient)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Авторизован как @%s", bot.Self.UserName)

	handleNewMessages()
}

// Отправка сообщения юзеру
func sendMessage(message tgbotapi.MessageConfig) {
	_, err := bot.Send(message)
	if err != nil {
		log.Panic(err)
	}
}

// Получение новых сообщений от юзера
func handleNewMessages() {
	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	messages := bot.GetUpdatesChan(update)
	for message := range messages {
		if message.Message != nil {

			state, _ := states.GetUser(message.Message.From.ID)
			if state != nil {

				switch state.State {

				// Главное меню
				case states.MainMenu:
					handleMainMenu(message)

				// Меню управления расписанием
				case states.ScheduleMenu:
					handleScheduleMenu(message)

				// Меню редактированиия времени выключения Wi-Fi
				case states.EditScheduleDown:
					handleEditScheduleDown(message)

				// Меню редактированиия времени включения Wi-Fi
				case states.EditScheduleUp:
					handleEditScheduleUp(message)
				}

			} else {
				// Если у юзера не установлено никакое состояние, то проверяем, что он в списке пользователей бота и отправляем главное меню
				config := ConfigManager.GetConfig()
				if utils.Int64InSlice(config.Users, message.Message.From.ID) {
					sendMainMenu(message.Message.Chat.ID)
				}
			}
		}
	}
}

// Отправка юзеру главного меню
func sendMainMenu(userId int64) {
	states.SetUser(userId, states.MainMenu)

	message := tgbotapi.NewMessage(userId, "Главное меню")
	message.ReplyMarkup = MainMenuKeyboard

	sendMessage(message)
}

// Реакция на команды юзера в главном меню
func handleMainMenu(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)
	newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "")
	switch messageText {

	// Включение Wi-FI
	case "вкл":
		WifiManager.InterfacesUp()
		newMessage.Text = "Успешно!"

	// Выключение Wi-FI
	case "выкл":
		WifiManager.InterfacesDown()
		newMessage.Text = "Успешно!"

	// Управление расписанием
	case "расписание":
		sendScheduleMenu(message.Message.From.ID)
		return

	default:
		return
	}

	sendMessage(newMessage)
}

// Получаем статус работы расписания в текстовом виде
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

// Отправка юзеру меню управления расписанием
func sendScheduleMenu(userId int64) {
	states.SetUser(userId, states.ScheduleMenu)

	message := tgbotapi.NewMessage(userId, "Управление расписанием")
	message.ReplyMarkup = ScheduleKeyboard

	sendMessage(message)
}

// Реакция на команды юзера в меню управления расписанием
func handleScheduleMenu(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)
	newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "")
	switch messageText {

	// Включение расписания
	case "вкл":
		config := ConfigManager.GetConfig()
		config.ScheduleEnabled = true
		ConfigManager.SetConfig(config)

		newMessage.Text = "Успешно!"

	// Выключение расписания
	case "выкл":
		config := ConfigManager.GetConfig()
		config.ScheduleEnabled = false
		ConfigManager.SetConfig(config)

		newMessage.Text = "Успешно!"

	// Просмотр статуса расписания
	case "просмотр":
		newMessage.Text = getScheduleStatus()

	// Редактирование расписания
	case "изменить":
		statusMessage := tgbotapi.NewMessage(message.Message.Chat.ID, getScheduleStatus())
		sendMessage(statusMessage)

		states.SetUser(message.Message.From.ID, states.EditScheduleDown)
		newMessage.Text = "Введите время выключения\nФормат: HH:MM"
		newMessage.ReplyMarkup = CancelKeyboard

	// Переход назад, в главное меню
	case "назад":
		sendMainMenu(message.Message.Chat.ID)
		return

	default:
		return
	}

	sendMessage(newMessage)
}

/*
Получение двух чисел из текста, который отправил юзер. Если текст не является числами, то возвращается nil.
Например:
20:30 -> 20, 30
06:30 -> 6, 30
рандомный текст -> nil
*/
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

// Реакция на команды юзера в меню редактирования расписания (время выключения Wi-FI)
func handleEditScheduleDown(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)

	// Переход назад, в меню управлением расписанием
	if messageText == "отмена" {
		sendScheduleMenu(message.Message.Chat.ID)

	} else {
		// Проверка и сохранение нового расписания
		newMessage := tgbotapi.NewMessage(message.Message.Chat.ID, "")
		data := parseEditScheduleMessage(messageText)

		if data != nil {
			config := ConfigManager.GetConfig()
			config.ScheduleEnabled = true
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

// Реакция на команды юзера в меню редактирования расписания (время включения Wi-FI)
func handleEditScheduleUp(message tgbotapi.Update) {
	messageText := strings.ToLower(message.Message.Text)

	// Переход назад, в меню управлением расписанием
	if messageText == "отмена" {
		sendScheduleMenu(message.Message.Chat.ID)

	} else {
		// Проверка и сохранение нового расписания
		data := parseEditScheduleMessage(messageText)

		if data != nil {
			config := ConfigManager.GetConfig()
			config.ScheduleEnabled = true
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
