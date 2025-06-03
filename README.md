# calculator-telegram-bot

## ТЗ:
Создать простого телеграм бота, который работает как калькулятор
вычисляет простые арифметические выражения
Нужно отправить ссылку на репозиторий, язык использовать Go
Так же отправить ссылку на работающего бота

### Информация о боте:
- **Ник бота:** @KD_Calculator_bot
- **Ссылка:** https://t.me/@KD_Calculator_bot

### Как запустить:

1. С помощью Докера

```bash
docker build -t calculator-bot .
docker run -p 8080:8080 calculator-bot
```

2. Без Докера

```
go run ./cmd/main.go
```

### Скриншоты:

![2025-06-03_19-40](https://github.com/user-attachments/assets/8fdb4b75-ff32-438d-a38b-b3b50ac77ddf)

![2025-06-03_19-40_1](https://github.com/user-attachments/assets/17725080-22b0-428d-91b0-0e52334ced8a)

P.S Он уже выставлен на удаленном сервере
