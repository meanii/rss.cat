package commands

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// Help provides usage instructions for the bot.
func Help(b *gotgbot.Bot, ctx *ext.Context) error {
	message := `*RSS.cat Help*

Here are the available commands:

/start - Welcome message and introduction
/add <rss-url> - Subscribe to an RSS feed
/remove <rss-url> - Remove an RSS feed from your subscriptions
/myfeeds - List all RSS feeds you have added
/stats - Show bot statistics (users, RSS feeds, notifications sent)
/help - Show this help message

*How to use:*
- Use /add followed by a valid RSS feed URL to subscribe.
- Use /remove followed by a valid RSS feed URL to unsubscribe (removal is soft, notification count is preserved).
- Use /myfeeds to view all your added RSS feeds.
- Use /stats to see bot statistics.
- You will receive updates when new items are published.

If you have any questions or need support, feel free to reach out.`
	_, err := ctx.EffectiveMessage.Reply(b, message, &gotgbot.SendMessageOpts{
		ParseMode: "Markdown",
	})
	return err
}
