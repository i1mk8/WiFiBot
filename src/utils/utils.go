package utils

import (
	"log"
	"os/exec"
	"strconv"
)

func Execute(name string, args []string) {
	result := exec.Command(name, args...)
	_, stderr := result.Output()

	if stderr != nil {
		log.Panic(stderr)
	}
}

func Int64InSlice(slice []int64, value int64) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}

	return false
}

func IntToString(value int) string {
	result := strconv.Itoa(value)
	if value < 10 {
		result = "0" + result
	}
	return result
}

func SaveFileSystem() {
	Execute("fs", []string{"save"})
}
