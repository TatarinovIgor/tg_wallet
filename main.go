package main

import (
	"github.com/yanzay/tbot"
	"handlers"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	bot, err := tbot.NewServer(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/answer", "42")
	bot.HandleFunc("/timer {seconds}", timerHandler)

	bot.HandleFunc("/sign_in", signInHandler())
	bot.HandleFunc("/log_in", logInHandler)
	bot.HandleFunc("/get_balance", getBalanceHandler)
	bot.HandleFunc("/pay_in", payInHandler)
	bot.HandleFunc("/pay_out", payOutHandler)
	bot.HandleFunc("/transfer", transferHandler)
	bot.HandleFunc("/about", aboutHandler)

	bot.ListenAndServe()
}

func timerHandler(m *tbot.Message) {
	// m.Vars contains all variables, parsed during routing
	secondsStr := m.Vars["seconds"]
	// Convert string variable to integer seconds value
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		m.Reply("Invalid number of seconds")
		return
	}
	m.Replyf("Timer for %d seconds started", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}
