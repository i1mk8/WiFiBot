package utils

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Выполнение команды
func Execute(name string, args []string) string {
	result := exec.Command(name, args...)
	stdout, stderr := result.Output()

	if stderr != nil {
		log.Panic(stderr)
	}
	return string(stdout)
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

/*
Получаем текущее время на роутере.
Эта функция необходима, так как time.Now() работает не корректно и всегда возвращает время по UTC, вместо установленной временной зоны
*/
func GetCurrentTime() (int, int) {
	output := Execute("date", []string{"+%H,%M,"})
	parsedOutput := strings.Split(output, ",")

	hour, _ := strconv.Atoi(parsedOutput[0])
	minute, _ := strconv.Atoi(parsedOutput[1])

	return hour, minute
}
