package commands

import (
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/meanii/rss.cat/database"
	"github.com/meanii/rss.cat/util"
	"github.com/mmcdole/gofeed"
)

// StartRSSBackgroundJob launches a goroutine to check RSS feeds and send updates.
func StartRSSBackgroundJob(bot *gotgbot.Bot) {
	go func() {
		for {
			var feeds []database.Rss
			database.SqlDB.Find(&feeds)
			fp := gofeed.NewParser()
			for _, feed := range feeds {
				parsed, err := fp.ParseURL(feed.Link)
				if err != nil || parsed == nil || len(parsed.Items) == 0 {
					log.Printf("Failed to parse RSS: %s", feed.Link)
					continue
				}
				latest := parsed.Items[0]
				// Use util.GetItemUniqueID for best practice unique identifier
				itemID := util.GetItemUniqueID(latest)
				if feed.LastItemGUID != itemID {
					// Update last item GUID in DB
					database.SqlDB.Model(&feed).Update("last_item_guid", itemID)
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
					// Increment NotificationCount for this feed
					database.SqlDB.Model(&feed).Update("notification_count", feed.NotificationCount+1)
				}
			}
			time.Sleep(5 * time.Minute) // Check every 5 minutes
		}
	}()
}
