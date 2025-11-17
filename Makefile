.PHONY: run build

BIN_FILE = tg-demo-bot.exe

run:
	go run cmd/bot/main.go



clean:
	del .\*.exe

build:
	go build -v -o ${BIN_FILE} ./cmd/bot/main.go
	

.DEFAULT_GOAL := build