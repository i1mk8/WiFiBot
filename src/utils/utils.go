package utils

import (
	"log"
	"os/exec"
	"strconv"
)

// Выполнение команды
func Execute(name string, args []string) {
	result := exec.Command(name, args...)
	_, stderr := result.Output()

	if stderr != nil {
		log.Panic(stderr)
	}
}

// Проверка наличия в списке числа
func Int64InSlice(slice []int64, value int64) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}

	return false
}

/*
Конвертация числа в строку.
Примеры:
18 -> 18
6 -> 06
*/
func IntToString(value int) string {
	result := strconv.Itoa(value)
	if value < 10 {
		result = "0" + result
	}
	return result
}

// Сохранение файловой системы (чтобы конфиг не сбрасывался при перезагрузке роутера)
func SaveFileSystem() {
	Execute("fs", []string{"save"})
}
