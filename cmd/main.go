package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Knetic/govaluate"
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
	bot.Debug = false

	commands := []tgbotapi.BotCommand{
		{Command: "start", Description: "Запустить бота"},
		{Command: "sourcecode", Description: "Узнать исходный код бота"},
	}
	if _, err := bot.Request(tgbotapi.NewSetMyCommands(commands...)); err != nil {
		log.Fatalf("Не удалось установить команды: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		chatID := update.Message.Chat.ID
		userInput := update.Message.Text

		switch userInput {
		case "/start":
			message := tgbotapi.NewMessage(
				chatID,
				"Привет! Нужно вычислить африметическое выражение?\n\nТы по адрессу!\n\nНабери любой пример и я его решу!",
			)
			bot.Send(message)
		case "/sourcecode":
			message := tgbotapi.NewMessage(
				chatID,
				"Ссылка на репозиторий проекта: github.com/DaniilKalts/calculator-telegram-bot",
			)
			bot.Send(message)
		default:
			result, err := evaluateExpression(userInput)

			var reply string
			if err != nil {
				reply = "Ты неправильно записал пример. Попробуй ещё раз!"
			} else {
				reply = userInput + " = " + result
			}

			bot.Send(tgbotapi.NewMessage(chatID, reply))
		}

	}
}

func evaluateExpression(expression string) (string, error) {
	evaluatedExpression, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	result, err := evaluatedExpression.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result), nil
}
