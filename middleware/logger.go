package middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/meanii/rss.cat/database"
	"gorm.io/gorm"
)

type LoggerBotClient struct {
	gotgbot.BotClient
}

func (b LoggerBotClient) RequestWithContext(ctx context.Context, token string, method string, params map[string]string, data map[string]gotgbot.FileReader, opts *gotgbot.RequestOpts) (json.RawMessage, error) {

	if chatID, ok := params["chat_id"]; ok && len(chatID) > 0 {
		chatIDInt64, err := strconv.ParseInt(chatID, 10, 64)
		if err != nil {
			return nil, err
		}
		var user database.User
		result := database.SqlDB.Where("user_id = ?", chatIDInt64).First(&user)
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			// User not found, create new
			user = database.User{UserId: chatIDInt64}
			database.SqlDB.Create(&user)
		} else {
			// User exists, update fields if needed
			database.SqlDB.Save(&user)
		}
	}

	// Call the next bot client instance in the middleware chain.
	val, err := b.BotClient.RequestWithContext(ctx, token, method, params, data, opts)
	if err != nil {
		// Middlewares can also be used to increase error visibility, in case they aren't logged elsewhere.
		log.Println("warning, got an error:", err)
	}
	return val, err
}

func LoggerMiddleware() gotgbot.BotClient {
	return &LoggerBotClient{
		BotClient: &gotgbot.BaseBotClient{
			Client:             http.Client{},
			UseTestEnvironment: false,
			DefaultRequestOpts: &gotgbot.RequestOpts{
				Timeout: gotgbot.DefaultTimeout,
				APIURL:  gotgbot.DefaultAPIURL,
			},
		},
	}
}
