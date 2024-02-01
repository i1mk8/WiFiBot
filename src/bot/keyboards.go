// Клавиатуры с командами для взаимодействия с ботом
package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	// Клавиатура главного меню
	MainMenuKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ВКЛ"),
			tgbotapi.NewKeyboardButton("ВЫКЛ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Расписание"),
		),
	)

	// Клавиатура меню управления расписанием
	ScheduleKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("ВКЛ"),
			tgbotapi.NewKeyboardButton("ВЫКЛ"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Просмотр"),
			tgbotapi.NewKeyboardButton("Изменить"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Назад"),
		),
	)

	// Клавиатура с кнопкой отмена (используется, когда юзеру нужно отправить какую-нибудь информацию боту. например, время включения/выключения Wi-FI)
	CancelKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Отмена"),
		),
	)
)
