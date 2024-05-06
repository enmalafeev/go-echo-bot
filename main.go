package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	botToken := os.Getenv("botToken")

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatId := tu.ID(update.Message.Chat.ID)
			_, _ = bot.CopyMessage(
				tu.CopyMessage(
					chatId,
					chatId,
					update.Message.MessageID,
				),
			)
		}
	}
}
