package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		panic("Нужен токен от телеграм-бота для запуска. Укажите в переменных окружения (.env)")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true // чтобы подробно видеть логи

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		chatID := update.Message.Chat.ID
		userInput := update.Message.Text

		if userInput == "/start" {
			message := tgbotapi.NewMessage(
				chatID,
				"Привет! Нужно вычислить африметическое выражение?\n\nТы по адрессу!\n\nНабери любой пример и я его решу!",
			)
			bot.Send(message)
			continue
		}
	}
}
