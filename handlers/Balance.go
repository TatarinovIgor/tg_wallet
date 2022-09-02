package handlers

import (
	"github.com/yanzay/tbot"
	"strconv"
	"time"
)

func getBalanceHandler(m *tbot.Message) {
	// m.Vars contains all variables, parsed during routing
	m.Replyf("Getting your current balance)
	time.Sleep(time.Duration(seconds) * time.Second)
	m.Reply("Time out!")
}
