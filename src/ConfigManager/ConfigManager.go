package ConfigManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	ConfigPath = "wifi_bot.json"
)

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

func SetConfig(config ConfigStruct) {
	configJson, err := json.Marshal(config)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(ConfigPath, configJson, 0644)
	if err != nil {
		log.Panic(err)
	}
}
