// Package telegram handles all Telegram-side logic.
package telegram

import (
	"log"
	"io/ioutil"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ritlug/teleirc/internal"
)

/*
Client contains information for the Telegram bridge, including
the TelegramSettings needed to run the bot
*/
type Client struct {
	api       *tgbotapi.BotAPI
	Settings  internal.TelegramSettings
	sendToIrc func(string)
}

/*
Creating variables for logging
*/
var (
	logFlags    = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	Info        = log.New(os.Stdout, "INFO: ", logFlags)
	Error       = log.New(os.Stderr, "ERROR: ", logFlags)
)

/*
NewClient creates a new Telegram bot client
*/
func NewClient(settings internal.TelegramSettings, tgapi *tgbotapi.BotAPI) *Client {
	Info.Println("Creating new Telegram bot client...")
	return &Client{api: tgapi, Settings: settings}
}

/*
SendMessage sends a message to the Telegram channel specified in the settings
*/
func (tg *Client) SendMessage(msg string) {
	newMsg := tgbotapi.NewMessage(tg.Settings.ChatID, "")
	newMsg.Text = msg
	tg.api.Send(newMsg)
}

/*
StartBot adds necessary handlers to the client and then connects,
returning any errors that occur
*/
func (tg *Client) StartBot(errChan chan<- error, sendMessage func(string)) {
	Info.Println("Starting up Telegram bot...")
	var err error
	tg.api, err = tgbotapi.NewBotAPI(tg.Settings.Token)
	if err != nil {
		Error.Println("Failed to connect to Telegram")
		errChan <- err
	}
	tg.sendToIrc = sendMessage

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tg.api.GetUpdatesChan(u)
	if err != nil {
		errChan <- err
		Error.Println(err)
	}

	// TODO: Move these lines into the updateHandler when available
	for update := range updates {
		if update.Message == nil {
			continue
		}

		messageHandler(tg, update)
	}
	errChan <- nil
}
