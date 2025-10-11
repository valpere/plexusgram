# plexusgram

A Multi-Language Knowledge Base Bot built with Go and the gotgbot framework.

## Overview

Plexusgram is a Telegram bot designed for corporate knowledge management and customer support automation. This portfolio project demonstrates advanced software engineering skills through progressive implementation phases, showcasing best practices in Go development, clean architecture, and enterprise-grade bot development.

**Business Value:**
- Corporate knowledge management
- Customer support automation
- Multi-language content delivery
- Searchable knowledge repository

## Features

**Phase 1 (Current):**
- Basic bot initialization and connection
- Command handlers (`/start`, `/help`)
- Environment-based configuration
- Long polling for updates

**Phase 2 - Knowledge Base Core:**
- Article storage and retrieval
- Search functionality
- Category/tag organization
- Basic content management commands

**Phase 3 - Advanced Features:**
- Multi-language support (i18n)
- Full-text search with PostgreSQL/Elasticsearch
- Role-based access control (RBAC)
- Content versioning system
- Usage analytics and tracking

## Tech Stack

- **Language:** Go 1.24.6
- **Framework:** gotgbot v2
- **Storage:** PostgreSQL (planned), Elasticsearch (planned)
- **Architecture:** Clean architecture with SOLID principles
- **Patterns:** Repository pattern, dependency injection

## Project Structure

```
plexusgram/
├── internal/
│   ├── types.go              # Common types
│   ├── config/
│   │   └── config.go         # Configuration loader
│   └── bot/
│       └── bot.go            # Bot core and handlers
├── main.go                   # Application entry point
├── Makefile                  # Build automation
├── .env.example              # Environment template
└── go.mod                    # Dependencies
```

## Getting Started

### Prerequisites

- Go 1.24.6 or higher
- A Telegram bot token from [@BotFather](https://t.me/botfather)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/valpere/plexusgram.git
cd plexusgram
```

2. Install dependencies:
```bash
make deps
```

3. Create environment file:
```bash
cp .env.example .env
```

4. Add your bot token to `.env`:
```bash
TELEGRAM_BOT_TOKEN=your_bot_token_here
```

### Running the Bot

**Using Make:**
```bash
make run
```

**Using Go directly:**
```bash
source .env && go run main.go
```

**Using the binary:**
```bash
make build
source .env && ./plexusgram
```

## Development

### Available Make Commands

Run `make help` to see all available commands:

- `make build` - Build the application binary
- `make run` - Run the application
- `make test` - Run all tests
- `make test-coverage` - Run tests with coverage
- `make fmt` - Format code
- `make vet` - Run go vet
- `make check` - Run formatting and vetting
- `make clean` - Remove build artifacts

### Project Guidelines

This project follows design principles documented in the global and project-level CLAUDE.md files:
- **DRY, YAGNI, KISS** - Simple, maintainable code
- **SOLID principles** - Clean architecture
- **GRASP patterns** - Proper responsibility assignment

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Generate HTML coverage report
make test-coverage-html
```

## Configuration

Environment variables (`.env` file):

- `TELEGRAM_BOT_TOKEN` - **Required.** Your Telegram bot token from BotFather
- `DEBUG` - Optional. Set to `true` for debug logging (default: `false`)

## License

MIT License - see [LICENSE](LICENSE) file for details.

## Author

Valentyn Solomko <valentyn.solomko@gmail.com>
