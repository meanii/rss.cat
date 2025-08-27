package commands

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/meanii/rss.cat/database"
)

// Stats provides bot statistics: number of users and RSS feeds.
func Stats(b *gotgbot.Bot, ctx *ext.Context) error {
	var userCount int64
	var rssCount int64
	var notifCount int64
	database.SqlDB.Model(&database.User{}).Count(&userCount)
	database.SqlDB.Model(&database.Rss{}).Count(&rssCount)
	database.SqlDB.Model(&database.Rss{}).Select("SUM(notification_count)").Scan(&notifCount)
	msg := fmt.Sprintf(
		"*RSS.cat Bot Stats*\n\nI am currently servicing *%d* RSS feeds and have reached *%d* users!\n\nA total of *%d* RSS notifications have been sent across all users.",
		rssCount, userCount, notifCount,
	)
	_, _ = ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	return nil
}
