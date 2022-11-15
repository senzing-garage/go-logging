# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
