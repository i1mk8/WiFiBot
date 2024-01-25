package ConfigManager

type ConfigStruct struct {
	BotToken           string  `json:"bot_token"`
	Users              []int64 `json:"bot_users"`
	Timezone           string  `json:"timezone"`
	ScheduleEnabled    bool    `json:"schedule_enabled"`
	ScheduleDownHour   int     `json:"schedule_down_hour"`
	ScheduleDownMinute int     `json:"schedule_down_minute"`
	ScheduleUpHour     int     `json:"schedule_up_hour"`
	ScheduleUpMinute   int     `json:"schedule_up_minute"`
}
