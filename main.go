package main

import (
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/meanii/rss.cat/commands"
	"github.com/meanii/rss.cat/database"
	"github.com/meanii/rss.cat/middleware"
)

func main() {
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN enviorment varible is empty")
	}

	// connecting sqlite3
	_, err := database.NewSqlConn("rss.cat")
	if err != nil {
		panic("failed to connect sqlite3")
	}

	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		BotClient: middleware.LoggerMiddleware(),
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Start background RSS job
	commands.StartRSSBackgroundJob(b)

	// Create updater and dispatcher.
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	updater := ext.NewUpdater(dispatcher, nil)

	// defining commands
	dispatcher.AddHandler(handlers.NewCommand("start", commands.Start))
	dispatcher.AddHandler(handlers.NewCommand("add", commands.Add))
	dispatcher.AddHandler(handlers.NewCommand("help", commands.Help))
	dispatcher.AddHandler(handlers.NewCommand("myfeeds", commands.MyFeeds))
	dispatcher.AddHandler(handlers.NewCommand("stats", commands.Stats))
	dispatcher.AddHandler(handlers.NewCommand("remove", commands.Remove))

	// Start receiving updates.
	err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
