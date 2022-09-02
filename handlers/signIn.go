package handlers

import (
	"github.com/yanzay/tbot"
)

func signInHandler(m *tbot.Message, bot tbot.Server) {
	m.Replyf("Please register at following link:")
	whitelist := []string{"yanzay", "user2"}
	bot.AddMiddleware(tbot.NewAuth(whitelist))
}
