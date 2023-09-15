package utils

import (
	"fmt"
	"os/exec"
	"time"
)

func ParserDur(durStr string) time.Duration {

	dur, err := time.ParseDuration(durStr)
	if err != nil {
		fmt.Printf("cant parse duration %s", err)
	}
	return dur
}

func Execut() string {
	cmd := exec.Command("uptime")

	// Выполняем команду и получаем вывод
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Ошибка выполнения команды:", err)
		// return
	}

	// Преобразуем вывод в строку и выводим его
	result := string(output)
	return result
}
