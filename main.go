package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	phrases := []string{
		"Do you always look this good?",
		"You have such beautiful eyes.",
		"I can't stop thinking about your smile.",
		"I wonder what you would do if I were there right now?",
		"Are you always this sweet or is it just for me?",
		"Guess who I'm thinking about right now?",
		"Do you think I can read minds? Because I know what you're thinking!",
		"If we had a date, where would you take me?",
		"Do you like surprises? I have one for you!",
		"Tell me something interesting about yourself.",
		"It's always a pleasure to talk to you, you know?",
		"What do you do in your free time?",
		"You're so confident, don't you have any weaknesses?",
		"Maybe we should see who's better at board games?",
		"Are you always this confident or is it just with me?",
		"Is there something you've always wanted to ask me?",
		"Guess what I want to tell you...",
		"Do you want to know my secret?",
		"You're always so attentive, it's very attractive.",
		"I like how you listen, it's rare.",
		"I can talk to you about anything.",
	}

	rand.Seed(time.Now().UnixNano())

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msgText := phrases[rand.Intn(len(phrases))]

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			bot.Send(msg)
		}
	}
}
