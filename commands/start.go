package commands

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// start introduces the bot.
func Start(b *gotgbot.Bot, ctx *ext.Context) error {
	message := fmt.Sprintf(
		"*Welcome to @%s!*\n\nI am your personal RSS feed assistant.\n\n*How to get started?*\n- Use `/add <rss-url>` to subscribe to any RSS feed.\n- I will notify you whenever there are new updates.\n\nIf you need help, just type `/help`.",
		b.User.FirstName,
	)
	_, err := ctx.EffectiveMessage.Reply(b, message, &gotgbot.SendMessageOpts{
		ParseMode: "Markdown",
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
