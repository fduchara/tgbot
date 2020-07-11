export GOPATH := $(PWD)
export GOBIN := $(PWD)

install:
	go get github.com/go-telegram-bot-api/telegram-bot-api

build:
	go build -v -o bot bot.go
