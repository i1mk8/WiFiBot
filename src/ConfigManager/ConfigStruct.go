// Список существующих полей в конфиге
package ConfigManager

type ConfigStruct struct {
	BotToken           string  `json:"bot_token"`            // Токен telegram бота
	Users              []int64 `json:"bot_users"`            // Список id юзеров бота
	ScheduleEnabled    bool    `json:"schedule_enabled"`     // Включено ли расписание
	ScheduleDownHour   int     `json:"schedule_down_hour"`   // Час, когда Wi-Fi выключается
	ScheduleDownMinute int     `json:"schedule_down_minute"` // Минута, когда Wi-Fi выключается
	ScheduleUpHour     int     `json:"schedule_up_hour"`     // Час, когда Wi-Fi включается
	ScheduleUpMinute   int     `json:"schedule_up_minute"`   // Минута, когда Wi-Fi включается
}
