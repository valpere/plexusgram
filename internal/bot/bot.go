/*
Copyright © 2025 Valentyn Solomko <valentyn.solomko@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/valpere/plexusgram/internal"
)

// Bot represents the Telegram bot instance
type Bot struct {
	client *gotgbot.Bot
	config *internal.Config
}

// New creates a new Bot instance
func New(cfg *internal.Config) (*Bot, error) {
	client, err := gotgbot.NewBot(cfg.BotToken, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot: %w", err)
	}

	return &Bot{
		client: client,
		config: cfg,
	}, nil
}

// Start begins polling for updates
func (b *Bot) Start() error {
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(bot *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Printf("Error handling update: %v", err)
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	// Register handlers
	b.registerHandlers(dispatcher)

	updater := ext.NewUpdater(dispatcher, nil)

	// Get bot info
	botInfo, err := b.client.GetMe(nil)
	if err != nil {
		return fmt.Errorf("failed to get bot info: %w", err)
	}

	log.Printf("Bot started: @%s", botInfo.Username)

	// Start polling
	err = updater.StartPolling(b.client, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to start polling: %w", err)
	}

	log.Println("Bot is running. Press Ctrl+C to stop.")
	updater.Idle()

	return nil
}

// registerHandlers registers all command and message handlers
func (b *Bot) registerHandlers(dispatcher *ext.Dispatcher) {
	// Command handlers
	dispatcher.AddHandler(handlers.NewCommand("start", b.handleStart))
	dispatcher.AddHandler(handlers.NewCommand("help", b.handleHelp))
}

// handleStart handles the /start command
func (b *Bot) handleStart(bot *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(bot, "Welcome to Plexusgram! 👋\n\nUse /help to see available commands.", nil)
	return err
}

// handleHelp handles the /help command
func (b *Bot) handleHelp(bot *gotgbot.Bot, ctx *ext.Context) error {
	helpText := `Available commands:

/start - Start the bot
/help - Show this help message

This is a demonstration Telegram bot built with Go and gotgbot.`

	_, err := ctx.EffectiveMessage.Reply(bot, helpText, nil)
	return err
}
