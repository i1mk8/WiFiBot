// Код для управления конфигом бота
package ConfigManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/i1mk8/WifiBot/utils"
)

const (
	ConfigPath = "wifi_bot.json" // Путь до конфига
)

// Получение конфига
func GetConfig() ConfigStruct {
	file, err := os.Open(ConfigPath)
	if err != nil {
		log.Panic(err)
	}

	bytes, _ := ioutil.ReadAll(file)
	var config ConfigStruct
	json.Unmarshal(bytes, &config)

	file.Close()
	return config
}

// Сохранение конфига
func SetConfig(config ConfigStruct) {
	configJson, err := json.Marshal(config)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(ConfigPath, configJson, 0644)
	if err != nil {
		log.Panic(err)
	}

	utils.SaveFileSystem()
}
