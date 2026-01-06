# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

go-logging is a Go logging library built on Go's experimental `slog` package. It provides message-number-based logging where the message number determines the log level, text, and identification.

## Build and Development Commands

```bash
# Build
make build                      # Build binary for current platform
make build-all                  # Build for all platforms (darwin/linux/windows, amd64/arm64)

# Test
make test                       # Run all tests with gotestfmt output
go test -v -run TestName ./...  # Run a single test

# Lint
make lint                       # Run all linters (golangci-lint, govulncheck, cspell)
make golangci-lint              # Run only golangci-lint
make fix                        # Run auto-fixers for lint issues

# Coverage
make coverage                   # Run tests with coverage, opens HTML report
make check-coverage             # Run coverage check against threshold

# Dependencies
make dependencies-for-development  # Install dev tools (golangci-lint, gotestfmt, etc.)
make dependencies                  # Update Go module dependencies

# Other
make run                        # Run main.go directly
make documentation              # Start godoc server
make clean                      # Clean build artifacts and caches
```

## Architecture

### Two Logging Packages

**`logger/`** - Simple traditional logger interface with method-based log levels:

- `Debug()`, `Info()`, `Warn()`, `Error()`, `Fatal()`, `Panic()` methods
- Includes formatted variants: `Debugf()`, `Infof()`, etc.
- Global default instance + ability to create custom instances

**`logging/`** - Message-number-based logging for Senzing tools:

- Single `Log(messageNumber, details...)` method
- Message number determines log level based on ranges:
  - 0-999: TRACE, 1000-1999: DEBUG, 2000-2999: INFO, 3000-3999: WARN
  - 4000-4999: ERROR, 5000-5999: FATAL, 6000-6999: PANIC
- Outputs structured JSON with fields: time, level, id, text, location, details, errors
- `NewSenzingLogger()` creates loggers for senzing-tools with component IDs
- Built on `go-messaging/messenger` for message formatting

### Key Types

- `logging.Logging` interface - Main interface for message-number logging
- `logging.BasicLogging` - Implementation of Logging interface
- `logger.Logger` interface - Traditional logger interface
- `logger.BasicLogger` - Implementation of Logger interface

### Guards

Both packages support guard methods (`IsDebug()`, `IsTrace()`, etc.) to avoid expensive logging calls when the level won't be printed:

```go
if logger.IsDebug() {
    logger.Log(1001, expensiveOperation())
}
```

## Dependencies

- `github.com/senzing-garage/go-messaging` - Message formatting and templates
- `github.com/senzing-garage/go-helpers` - Error wrapping utilities
- `golang.org/x/exp/slog` - Structured logging foundation

## Linting Configuration

Linting config is in `.github/linters/.golangci.yaml`. Key settings:

- Max line length: 120 characters
- Max cyclomatic complexity: 20
- Exhaustruct exclusions for BasicLogger, BasicLogging, ExtractedValues
- Uses gofumpt for formatting
