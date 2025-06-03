FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o calculator-telegram-bot ./cmd/main.go

EXPOSE 8080

CMD ["./calculator-telegram-bot"]
