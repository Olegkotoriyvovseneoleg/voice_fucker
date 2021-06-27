package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func _check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	botAPIKey, isSet := os.LookupEnv("TELEGRAM_BOT_API_KEY_VOICE")
	if !isSet {
		log.Panic("No bot api key found. Please set TELEGRAM_BOT_API_KEY env")
	}
	// create bot using token, client
	bot, err := tgbotapi.NewBotAPI(botAPIKey)
	_check(err)

	// debug mode on
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// set update interval
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 1000

	updates, err := bot.GetUpdatesChan(u)
	_check(err)

	// get new updates
	for update := range updates {
		// if message from user
		if update.Message != nil {
	    if update.Message.Voice != nil {
	      log.Printf("Message is voice, duration %d", update.Message.Voice.Duration)

				replyMessage(update.Message.Chat.ID, update.Message.MessageID, bot)
	    }
		}
	}
}

func replyMessage(chatId int64, messageId int, bot *tgbotapi.BotAPI) {
  replyMessage := tgbotapi.NewMessage(chatId, GetBadWord())
	replyMessage.ParseMode = "HTML"
	replyMessage.ReplyToMessageID = messageId

	bot.Send(replyMessage)
}
