# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**plexusgram** - A Multi-Language Knowledge Base Bot built with Go and the gotgbot framework.

This is a portfolio project designed to showcase software engineering skills through progressive implementation phases:
- Phase 1: Basic bot infrastructure and command handling
- Phase 2: Knowledge base core (article storage, search, categories)
- Phase 3: Advanced features (i18n, full-text search, RBAC, versioning, analytics)

**Business Value**: Corporate knowledge management, customer support automation

**Tech Stack**: Go 1.24.6, gotgbot framework, PostgreSQL (planned), Elasticsearch (planned)

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
  - `handlers/` - Command and message handlers for bot interactions
  - `storage/` - Data persistence layer (repository pattern)
    - `models/` - Data models (Article, Category, User, etc.)
    - `repository/` - Repository interfaces and implementations
  - `services/` - Business logic layer
    - `knowledge/` - Knowledge base operations (CRUD, search)
    - `i18n/` - Internationalization support
    - `auth/` - Authentication and authorization (RBAC)
  - `search/` - Full-text search integration (Elasticsearch)
  - `analytics/` - Usage tracking and metrics

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
4. Focus on showcasing software engineering skills for corporate knowledge management use cases

## Implementation Priorities

**Phase 2 Priorities (Knowledge Base Core):**
1. Define data models (Article, Category, Tag)
2. Implement repository pattern with PostgreSQL
3. Add article CRUD commands (/add, /search, /view, /delete)
4. Implement basic search functionality
5. Add category/tag organization

**Phase 3 Priorities (Advanced Features):**
1. Multi-language content support (i18n)
2. Elasticsearch integration for full-text search
3. RBAC system (admin, editor, viewer roles)
4. Content versioning and change tracking
5. Usage analytics and reporting

## Environment Setup

Required environment variables (create `.env` file):
```
# Telegram Bot
TELEGRAM_BOT_TOKEN=your_bot_token_here

# Database (Phase 2+)
DATABASE_URL=postgres://user:password@localhost:5432/plexusgram
DATABASE_MAX_CONNECTIONS=25

# Elasticsearch (Phase 3+)
ELASTICSEARCH_URL=http://localhost:9200
ELASTICSEARCH_INDEX=knowledge_base

# Optional
DEBUG=false
LOG_LEVEL=info
```

Note: `.env` is gitignored for security.
- to memorize