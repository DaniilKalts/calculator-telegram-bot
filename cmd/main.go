package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Нужен токен от телеграм-бота для запуска. Укажите в переменных окружения (.env)")
	}
}
