package commands

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/meanii/rss.cat/database"
)

// Remove removes an RSS feed for the user by soft-deleting it, preserving notification count.
func Remove(b *gotgbot.Bot, ctx *ext.Context) error {
	text := ctx.EffectiveMessage.Text
	args := strings.Fields(text)
	if len(args) < 2 {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"❌ *Invalid command format.*\nPlease use: `/remove <rss-url>`",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	rssURL := args[1]
	var rss database.Rss
	result := database.SqlDB.Where("link = ? AND owner_id = ?", rssURL, ctx.EffectiveMessage.From.Id).First(&rss)
	if result.Error != nil {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"❌ *RSS feed not found in your subscriptions.*",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	// Soft delete using GORM's built-in soft delete
	err := database.SqlDB.Delete(&rss).Error
	if err != nil {
		_, _ = ctx.EffectiveMessage.Reply(
			b,
			"❌ *Failed to remove RSS feed.*",
			&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
		)
		return nil
	}

	_, _ = ctx.EffectiveMessage.Reply(
		b,
		"✅ RSS feed removed.",
		&gotgbot.SendMessageOpts{ParseMode: "Markdown"},
	)
	return nil
}
