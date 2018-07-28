package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/samurang87/availabot/calendar_checker"
	"github.com/yanzay/tbot"
)

// DefaultHandler receives all messages sent by Telegram to the bot
func DefaultHandler(message *tbot.Message) {

	now := time.Now()
	busyCal, err := calendar_checker.GetBusyCalendar(now)

	if err != nil {
		log.Fatal(err)
	}

	result, err := calendar_checker.GetNextThreeEvenings(now, busyCal)

	if err != nil {
		log.Fatal(err)
	}

	message.Reply(fmt.Sprint(result))
}

func main() {

	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.HandleDefault(DefaultHandler)

	bot.ListenAndServe()

}
