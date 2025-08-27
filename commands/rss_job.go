package commands

import (
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/meanii/rss.cat/database"
	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

// StartRSSBackgroundJob launches a goroutine to check RSS feeds and send updates.
func StartRSSBackgroundJob(bot *gotgbot.Bot) {
	go func() {
		for {
			var feeds []database.Rss
			database.SqlDB.Where("active = ?", true).Find(&feeds)
			fp := gofeed.NewParser()
			for _, feed := range feeds {
				parsed, err := fp.ParseURL(feed.Link)
				if err != nil || parsed == nil || len(parsed.Items) == 0 {
					log.Printf("Failed to parse RSS: %s", feed.Link)
					continue
				}
				latest := parsed.Items[0]
				if feed.LastItemGUID != latest.GUID {
					// Update last item GUID in DB
					database.SqlDB.Model(&feed).Update("LastItemGUID", latest.GUID)
					// Extract website from feed.Link
					website := feed.Link
					if parsed.FeedLink != "" {
						website = parsed.FeedLink
					}
					msg := website + "\n[" + latest.Title + "](" + latest.Link + ")"
					// Send to owner
					_, _ = bot.SendMessage(feed.OwnerId, msg, &gotgbot.SendMessageOpts{ParseMode: "Markdown", LinkPreviewOptions: &gotgbot.LinkPreviewOptions{
						IsDisabled: true,
					}})
					// Increment notification count for owner
					database.SqlDB.Model(&database.User{}).Where("user_id = ?", feed.OwnerId).Update("NotificationCount", gorm.Expr("NotificationCount + 1"))
					// TODO: Send to subscribers if implemented
				}
			}
			time.Sleep(5 * time.Minute)
		}
	}()
}
