package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os/exec"
	"strings"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("XXXXXXXXXxtgbottokenXXXXXXXXXXX")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			text := strings.Split(update.Message.Text, " ")
			command := text[0]
			switch command {
			case "start":
				log.Printf("com start")
				reply := "hi im bot"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)

			case "ls":
				log.Printf("com ls")
				out, err := exec.Command("/bin/ls", "-1").Output()
				if err != nil {
					log.Fatal(err)
				}
				reply := string(out)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)

		    default:
				log.Printf("com def")
				out, err := exec.Command("/bin/sh", "-c", command).Output()
				if err != nil {
					log.Fatal(err)
				}
				reply := string(out)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)
			}
		}
	}
}
