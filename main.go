package main

import (
	"log"
	"regexp"
  "strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
)

// SystemConfig contains system config
var SystemConfig ConfigData

func _check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {

	// load environment variables
	SystemConfig = loadConfigFromEnv()

	// create bot using token, client
	bot, err := tgbotapi.NewBotAPI(SystemConfig.botAPIKey)
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
				var user *tgbotapi.User
				user = update.Message.From
				go newMessageFromBastard(user.UserName, update.Message.Chat.ID)
	    }

			niceParse, _ := regexp.Compile("/top")
			if niceParse.MatchString(update.Message.Text) {
				replyTop3(update.Message.Chat.ID, update.Message.MessageID, bot)
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

func replyTop3(chatId int64, messageId int, bot *tgbotapi.BotAPI) {
	log.Println("Get top 3")
  answer := "Топ дибилов:\n"

	for _, bastardRate := range getTop3(chatId) {
		answer = answer + bastardRate.name + " : " + strconv.Itoa(bastardRate.count) + "\n"
	}

  replyMessage := tgbotapi.NewMessage(chatId, answer)
	replyMessage.ParseMode = "HTML"
	replyMessage.ReplyToMessageID = messageId

	bot.Send(replyMessage)
}
