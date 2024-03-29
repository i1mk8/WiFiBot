// Код для управления Wi-Fi
package WifiManager

import (
	"time"

	"github.com/i1mk8/WifiBot/ConfigManager"
	"github.com/i1mk8/WifiBot/utils"
)

var (
	commandName = "iwpriv"
	commandArgs = [3]string{"", "set", ""}
	interfaces  = [2]string{"ra0", "rai0"} // Интерфейсы Wi-Fi
)

/*
Установка состояния работы интерфейсов.
1 - Включено
0 - Выключено
*/
func setInterfacesState(state string) {
	commandArgs[2] = "RadioOn=" + state

	for _, element := range interfaces {
		commandArgs[0] = element
		utils.Execute(commandName, commandArgs[:])
	}
}

// Включение Wi-Fi
func InterfacesUp() {
	setInterfacesState("1")
}

// Выключение Wi-Fi
func InterfacesDown() {
	setInterfacesState("0")
}

// Авто включение/выключение Wi-Fi по расписанию
func Auto() {
	for true {
		config := ConfigManager.GetConfig()
		if config.ScheduleEnabled {

			hour, minute := utils.GetCurrentTime()
			if hour == config.ScheduleDownHour && minute == config.ScheduleDownMinute {
				InterfacesDown()
			} else if hour == config.ScheduleUpHour && minute == config.ScheduleUpMinute {
				InterfacesUp()
			}

		}
		time.Sleep(60 * time.Second)
	}
}
