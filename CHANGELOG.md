# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.5.0] - 2024-06-07

### Changed in 1.5.0

- From `LoggerDefault` to `BasicLogger`
- From `LoggerInterface` to `Logger`
- From `LoggingImpl` to `BasicLogging`
- From `LoggingInterface` to `Logging`

## [1.4.1] - 2024-04-19

### Changed in 1.4.1

- Update dependencies
  - github.com/stretchr/testify v1.9.0
  - golang.org/x/exp v0.0.0-20240416160154-fe59bbe5cc7f

## [1.4.0] - 2023-12-29

### Changed in 1.4.0

- Renamed module to `github.com/senzing-garage/go-logging`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-messaging v1.4.0
  - golang.org/x/exp v0.0.0-20231226003508-02704c960a9b

## [1.3.3] - 2023-10-16

### Changed in 1.3.3

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-messaging v0.3.3
  - golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63

## [1.3.2] - 2023-08-04

### Changed in 1.3.2

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-messaging v0.3.2
  - golang.org/x/exp v0.0.0-20230801115018-d63ba01acd4b

## [1.3.1] - 2023-07-13

### Changed in 1.3.1

- Update dependencies
  - github.com/senzing-garage/go-messaging v0.3.1

## [1.3.0] - 2023-07-06

### Changed in 1.3.0

- Using "github.com/senzing-garage/go-messaging/messenger" to manage message format.

### Deleted in 1.3.0

- `messagedate`
- `messagedetails`
- `messageduration`
- `messageerrors`
- `messageformat`
- `messageid`
- `messagelevel`
- `messagelocation`
- `messagelogger`
- `messagestatus`
- `messagetext`
- `messagetime`
- ``

## [1.2.6] - 2023-06-16

### Changed in 1.2.6

- Update dependencies
  - github.com/senzing-garage/go-messaging v0.2.2
  - github.com/stretchr/testify v1.8.4
  - golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1

## [1.2.5] - 2023-05-23

### Changed in 1.2.5

- Changed `Error()` to `NewError()`

## [1.2.4] - 2023-05-23

### Changed in 1.2.4

- Added documentation on `Log()`, `Error()`, and logging output.
- Removed a layer of "caller depth" in `logging.Error()`

## [1.2.3] - 2023-05-11

### Changed in 1.2.3

- Update dependencies
  - github.com/senzing-garage/go-messaging v0.2.1

## [1.2.2] - 2023-05-11

### Changed in 1.2.2

- Changed use of `slog.New(...)`
- Update dependencies
  - golang.org/x/exp v0.0.0-20230510235704-dd950f8aeaea

## [1.2.1] - 2023-04-19

### Added in 1.2.1

- Change to UTC time.

## [1.2.0] - 2023-04-17

### Added in 1.2.0

- `go-logging/logging`

### Deprecated in 1.2.0

- The following packages will probably be removed in version 2.0.0:
  - `logger`
  - `messagedate`
  - `messagedetails`
  - `messageduration`
  - `messageerrors`
  - `messageformat`
  - `messageid`
  - `messagelevel`
  - `messagelocation`
  - `messagelogger`
  - `messagestatus`
  - `messagetext`
  - `messagetime`

## [1.1.3] - 2023-01-04

### Added in 1.1.3

- Fix ineffective assignments and spelling errors

## [1.1.2] - 2023-01-04

### Added in 1.1.2

- Support `messagelogger.Level` type as a  `messagelogger.New()` parameter

## [1.1.1] - 2022-11-18

### Added in 1.1.1

- `NewSenzingApiLogger()`
- Ability to choose `location` level using `CallerSkip`

### Changed in 1.1.1

- Improved `status` and `level` determination
- Improved `details` representation
- Improved godoc documentation
- Improved testing

## [1.1.0] - 2022-11-15

### Added in 1.1.0

- `date`, `time`, `duration`, `level`, and `location` fields

### Changed in 1.1.0

- Examples reflect TRACE, DEBUG, INFO, WARN, ERROR, FATAL, PANIC ordering
- Removed non-JSON elements to JSON messages

## [1.0.1] - 2022-11-08

### Changed in 1.0.1

- Added `messagelogger.NewSenzingLogger()` for use in Senzing applications
- Made distinct package implementations for NewSenzingLogger
- All `messageloglevel` implementations honor a specific log level in the details parameter
- Exposed `LevelToTextMap` and `TextToLevelMap`
- Fixed HTML escaping of JSON
- Log `nil` better
- Removed quotes from `map[]` output.
- Improved test cases

## [1.0.0] - 2022-11-01

### Added to 1.0.0

- Initial implementation
