package WifiManager

import (
	"fmt"
	"log"
	"time"

	"github.com/i1mk8/WifiBot/ConfigManager"
	"github.com/i1mk8/WifiBot/utils"
)

const (
	CommandUp   = "iwpriv %s set RadioOn=1"
	CommandDown = "iwpriv %s set RadioOn=0"
)

var interfaces = [2]string{"ra0", "rai0"}

func InterfacesUp() {
	for _, element := range interfaces {
		utils.Execute(fmt.Sprintf(CommandUp, element))
	}
}

func InterfacesDown() {
	for _, element := range interfaces {
		utils.Execute(fmt.Sprintf(CommandDown, element))
	}
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
