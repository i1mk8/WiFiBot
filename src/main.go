package main

import (
	WifiManager "github.com/i1mk8/WifiBot/WiFiManager"
	"github.com/i1mk8/WifiBot/bot"
)

func main() {
	go WifiManager.Auto()
	bot.StartBot()
}
