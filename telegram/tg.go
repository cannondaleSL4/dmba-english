package telegram

import (
	"flag"
	"fmt"
	"github.com/dmba-english/db"
	"github.com/dmba-english/service"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var tokenTg = flag.String("tokenTg", os.Getenv("TOKEN_TELEGRAM"), "your token")
var userWords map[int][]*db.Dict = make(map[int][]*db.Dict)

//var bot tgbotapi.BotAPI
//
//func init() {
//	bot, err := tgbotapi.NewBotAPI(*tokenTg)
//	if err != nil {
//		log.Panic(err)
//	}
//	bot.Debug = true
//}

func Tg() {

	bot, err := tgbotapi.NewBotAPI(*tokenTg)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 15

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {

		command := update.Message.Command()
		if command == "" {
			// Здесь логика для "обычных" сообщений
		} else if command == "learn" {
			learn(bot, update)
		} else if command == "exam" {
			exam(bot, update)
		}

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
	}
}

func learn(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	userId := update.Message.From.ID
	words := service.GetWords(userId)
	userWords[userId] = words
	for _, element := range words {
		fmt.Println(element)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, element.Word+" : "+element.Word_translate)
		//msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}

func exam(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	//userId := update.Message.From.ID
	//
	//for _, element := range userWords[userId] {
	//	fmt.Println(element)
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, element.Word + " : " + element.Word_translate)
	//	//msg.ReplyToMessageID = update.Message.MessageID
	//	bot.Send(msg)
	//}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 15

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {

		command := update.Message.Command()
		fmt.Println(command)
		//if command == "" {
		//	// Здесь логика для "обычных" сообщений
		//} else if command == "learn" {
		//	learn(bot, update)
		//} else if command == "exam" {
		//	exam(bot, update)
		//}
		//
		//if update.Message == nil { // ignore any non-Message Updates
		//	continue
		//}
	}
}
