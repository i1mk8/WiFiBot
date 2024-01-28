package WifiManager

import (
	"log"
	"time"

	"github.com/i1mk8/WifiBot/ConfigManager"
	"github.com/i1mk8/WifiBot/utils"
)

var (
	commandName = "iwpriv"
	commandArgs = [3]string{"", "set", ""}
	interfaces  = [2]string{"ra0", "rai0"}
)

func setInterfacesState(state string) {
	commandArgs[2] = "RadioOn=" + state

	for _, element := range interfaces {
		commandArgs[0] = element
		utils.Execute(commandName, commandArgs[:])
	}
}

func InterfacesUp() {
	setInterfacesState("1")
}

func InterfacesDown() {
	setInterfacesState("0")
}

func Auto() {
	for true {
		config := ConfigManager.GetConfig()
		if config.ScheduleEnabled {

			location, err := time.LoadLocation(config.Timezone)
			if err != nil {
				log.Panic(err)
			}

			nowTime := time.Now().In(location)
			if nowTime.Hour() == config.ScheduleDownHour && nowTime.Minute() == config.ScheduleDownMinute {
				InterfacesDown()
			} else if nowTime.Hour() == config.ScheduleUpHour && nowTime.Minute() == config.ScheduleUpMinute {
				InterfacesUp()
			}

		}
		time.Sleep(60 * time.Second)
	}
}
