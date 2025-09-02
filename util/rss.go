package util

import "github.com/mmcdole/gofeed"

// GetItemUniqueID returns a unique identifier for an RSS item using best practices.
// Prefers GUID, then Link, then Title+Published.
func GetItemUniqueID(item *gofeed.Item) string {
	if item == nil {
		return ""
	}
	if item.GUID != "" {
		return item.GUID
	}
	if item.Link != "" {
		return item.Link
	}
	return item.Title + item.Published
}
