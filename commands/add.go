package commands

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/meanii/rss.cat/database"
	"github.com/mmcdole/gofeed"
)

// start introduces the bot.
func Add(b *gotgbot.Bot, ctx *ext.Context) error {

	text := ctx.EffectiveMessage.Text
	args := strings.Fields(text)
	if len(args) < 2 {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"❌ *Invalid command format.*\nPlease use: `/add <rss-url>`",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	rssURL := args[1]

	// Check for duplicate RSS for this user
	var existing database.Rss
	result := database.SqlDB.Where("link = ? AND owner_id = ?", rssURL, ctx.EffectiveMessage.From.Id).First(&existing)
	if result.Error == nil {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"ℹ️ This RSS feed is already added to your subscriptions.",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	// Validate RSS feed
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(rssURL)
	if err != nil || feed == nil {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"❌ *Failed to add RSS feed.*\nReason: The provided URL is not a valid RSS feed.",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	// Get the latest item GUID if available
	var lastItemGUID string
	if len(feed.Items) > 0 {
		lastItemGUID = feed.Items[0].GUID
	}

	// Store in database
	rss := database.Rss{
		Link:         rssURL,
		OwnerId:      ctx.EffectiveMessage.From.Id,
		LastItemGUID: lastItemGUID,
	}
	err = database.SqlDB.Create(&rss).Error
	if err != nil {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"❌ *Database error.*\nCould not save your RSS feed. Please try again later.",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	// Success reply
	_, _ = ctx.EffectiveMessage.Reply(
		b,
		"✅ *RSS feed added successfully!*\nYou will now receive updates from this feed.",
		&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
	)
	return nil
}
