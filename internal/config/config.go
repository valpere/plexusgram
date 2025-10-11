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
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/valpere/plexusgram/internal"
)

// Load reads configuration from environment variables
// Automatically loads .env file if it exists
func Load() (*internal.Config, error) {
	// Load .env file if it exists (silently ignore if not found)
	_ = godotenv.Load()

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	debug := os.Getenv("DEBUG") == "true"

	return &internal.Config{
		BotToken: botToken,
		Debug:    debug,
	}, nil
}
