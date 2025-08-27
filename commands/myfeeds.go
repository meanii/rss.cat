package commands

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/meanii/rss.cat/database"
)

// MyFeeds lists all RSS URLs added by the owner.
func MyFeeds(b *gotgbot.Bot, ctx *ext.Context) error {
	ownerId := ctx.EffectiveMessage.From.Id
	var feeds []database.Rss
	database.SqlDB.Where("owner_id = ?", ownerId).Find(&feeds)
	if len(feeds) == 0 {
		_, _ = ctx.EffectiveMessage.Reply(b, "You have not added any RSS feeds yet.",
			&gotgbot.SendMessageOpts{
				ParseMode: "Markdown",
				LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
					IsDisabled: true,
				}})
		return nil
	}
	msg := "*Your RSS Feeds:*\n"
	for i, feed := range feeds {
		msg += fmt.Sprintf("%d. [%s](%s)\n", i+1, feed.Link, feed.Link)
	}
	_, _ = ctx.EffectiveMessage.Reply(b, msg, &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	return nil
}
