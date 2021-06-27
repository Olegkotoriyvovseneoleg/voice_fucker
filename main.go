package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)
var badWords = []string{
	"в очко себе это надиктуй",
	"что за ебанный писк?",
	"у меня кошка, когда блюет, ито звучит лучше",
	"судя по интонации у тебя явные признаки ДЦП",
}

// GetBadWord ...
func GetBadWord() string {
	return badWords[rand.Intn(len(badWords))]
}

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
		// if message from channel
		if update.ChannelPost != nil {
      log.Printf("Chat message %s", update.ChannelPost.Chat.ID)
  		// if message from user
  		if update.ChannelPost != nil {
        log.Printf("Chat message text %s", update.Message.Text)
  		}

      if update.ChannelPost.Voice != nil {
        log.Printf("Chat message is voice")
        log.Printf("Chat message duration %d", update.Message.Voice.Duration)
      }

			replyMessage(update.ChannelPost.Chat.ID, update.ChannelPost.MessageID, bot)
		}

		// if message from user
		if update.Message != nil {
      log.Printf("Message text %s", update.Message.Text)
		}

    if update.Message.Voice != nil {
      log.Printf("Message is voice")
      log.Printf("message duration %d", update.Message.Voice.Duration)
			replyMessage(update.Message.Chat.ID, update.Message.MessageID, bot)
    }
	}
}

func replyMessage(chatId int64, messageId int, bot *tgbotapi.BotAPI) {
  replyMessage := tgbotapi.NewMessage(chatId, GetBadWord())
	replyMessage.ParseMode = "HTML"
	replyMessage.ReplyToMessageID = messageId

	bot.Send(replyMessage)
}
