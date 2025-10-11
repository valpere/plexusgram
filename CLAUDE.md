# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**plexusgram** - A demonstration Telegram bot built with Go and the gotgbot framework.

This is a portfolio project designed to showcase software engineering skills through progressive implementation phases:
- Phase 1: Simple chatbot functionality
- Phase 2: Game features
- Phase 3: Utility functions and service integrations

**Tech Stack**: Go 1.24.6, gotgbot framework

## Project Context

Additional context files are located in `../context-plexusgram/`:
- `00-prompts.md` - Project requirements and planning discussions
- `Valentyn_Solomko_Software_Engineer-new.md` - Developer CV/background

When files referenced in code or documentation aren't found in the working directory, check the context directory.

## Development Commands

### Building and Running
```bash
# Run the application
go run main.go

# Build binary
go build -o plexusgram main.go

# Run tests
go test ./...

# Run specific test
go test -run TestName ./path/to/package
```

### Dependency Management
```bash
# Add new dependency
go get github.com/package/name

# Update dependencies
go get -u ./...

# Tidy dependencies
go mod tidy

# Vendor dependencies (if needed)
go mod vendor
```

### Code Quality
```bash
# Format code
go fmt ./...

# Lint (if golangci-lint is installed)
golangci-lint run

# Vet code
go vet ./...
```

### Testing
```bash
# Run tests with coverage
go test -cover ./...

# Generate coverage profile
go test -coverprofile=coverage.out ./...

# View coverage in browser
go tool cover -html=coverage.out
```

## Architecture

### Current State
The project is in early initialization stage with:
- Basic `main.go` entry point (currently empty main function)
- Go module configuration (`go.mod`)
- MIT License

### Planned Architecture
Following SOLID principles and Go best practices from global CLAUDE.md:

**Expected Package Structure:**
- `internal/` - Internal application packages
  - `types.go` - Common types used across packages
  - `bot/` - Telegram bot core functionality
  - `handlers/` - Command and message handlers
  - `storage/` - Data persistence layer
  - `services/` - Business logic and external integrations

**Type Organization:**
- Common types in `internal/types.go`
- Package-specific types in `internal/{package}/types.go`
- Higher-level types override lower-level ones in case of conflicts

### Design Principles (from global CLAUDE.md)
- **DRY, YAGNI, KISS** - Avoid duplication, implement only what's needed, keep it simple
- **SOLID** - Single responsibility, open/closed, Liskov substitution, interface segregation, dependency inversion
- **GRASP** - Information expert, creator, low coupling, high cohesion

## Development Workflow

1. This is a portfolio project with progressive complexity phases
2. Each phase should be completable within approximately one week
3. Start with local development, then progress to cloud deployment
4. Focus on showcasing software engineering skills rather than building production-ready features

## Environment Setup

Required environment variables (create `.env` file):
```
TELEGRAM_BOT_TOKEN=your_bot_token_here
```

Note: `.env` is gitignored for security.
