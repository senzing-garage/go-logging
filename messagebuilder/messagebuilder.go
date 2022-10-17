/*
Package helper ...
*/
package messagebuilder

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/senzing/go-logging/messageformat"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func init() {
	logger = New()
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (logger *Logger) printf(debugLevelName string, format string, v ...interface{}) *Logger {
	var message string
	calldepth := 3
	if format == noFormat {
		v := append(v, 0)
		copy(v[1:], v[0:])
		v[0] = debugLevelName + " "
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(debugLevelName+" "+format, v...)
	}
	log.Output(calldepth, message)
	return logger
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func SetLevel(level Level) *Logger { return logger.SetLevel(level) }
func (logger *Logger) SetLevel(level Level) *Logger {
	logger.isPanic = level <= LevelPanic
	logger.isFatal = level <= LevelFatal
	logger.isError = level <= LevelError
	logger.isWarn = level <= LevelWarn
	logger.isInfo = level <= LevelInfo
	logger.isDebug = level <= LevelDebug
	logger.isTrace = level <= LevelTrace
	return logger
}

func New() *Logger {
	return new(Logger)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build an error function.
func BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	return logger.BuildError(idTemplate, errorNumber, message, details...)
}

// Build an error method.
func (logger *Logger) BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	errorMessage := messageformat.BuildMessage(
		logger.BuildMessageId(idTemplate, errorNumber),
		logger.BuildMessageLevel(errorNumber, message),
		message,
		details...,
	)
	return errors.New(errorMessage)
}

// Build log message function.
func BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string {
	return logger.BuildMessage(idTemplate, errorNumber, message, details...)
}

// Build log message method.
func (logger *Logger) BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string {
	return messageformat.BuildMessage(
		logger.BuildMessageId(idTemplate, errorNumber),
		logger.BuildMessageLevel(errorNumber, message),
		message,
		details...,
	)
}

// Build log message function.
func BuildMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) string {
	return logger.BuildMessageFromError(idTemplate, errorNumber, message, err, details...)
}

// Build log message method.
func (logger *Logger) BuildMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) string {
	return messageformat.BuildMessageFromError(
		logger.BuildMessageId(idTemplate, errorNumber),
		logger.BuildMessageLevel(errorNumber, message),
		message,
		anError,
		details...,
	)
}

// Build log message function.
func BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) string {
	return logger.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
}

// Build log message method.
func (logger *Logger) BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) string {
	return messageformat.BuildMessageFromErrorUsingMap(
		logger.BuildMessageId(idTemplate, errorNumber),
		logger.BuildMessageLevel(errorNumber, message),
		message,
		anError,
		details,
	)
}

// Build log message function.
func BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string {
	return logger.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
}

// Build log message method.
func (logger *Logger) BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string {
	return messageformat.BuildMessageUsingMap(
		logger.BuildMessageId(idTemplate, errorNumber),
		logger.BuildMessageLevel(errorNumber, message),
		message,
		details,
	)
}

// Construct a unique message id function.
func BuildMessageId(idTemplate string, errorNumber int) string {
	return logger.BuildMessageId(idTemplate, errorNumber)
}

// Construct a unique message id method.
func (logger *Logger) BuildMessageId(idTemplate string, errorNumber int) string {
	return fmt.Sprintf(idTemplate, errorNumber)
}

// Based on the errorNumber and Senzing error code, get the message level function.
func BuildMessageLevel(errorNumber int, message string) string {
	return logger.BuildMessageLevel(errorNumber, message)
}

// Based on the errorNumber and Senzing error code, get the message level method.
func (logger *Logger) BuildMessageLevel(errorNumber int, message string) string {

	var result = "unknown"

	// Create a list of sorted keys.

	messageLevelKeys := make([]int, 0, len(MessageLevelMap))
	for key := range MessageLevelMap {
		messageLevelKeys = append(messageLevelKeys, key)
	}
	sort.Ints(messageLevelKeys)

	// Using the sorted message number, find the level.

	for _, messageLevelKey := range messageLevelKeys {
		if errorNumber < messageLevelKey {
			result = MessageLevelMap[messageLevelKey]
			break
		}
	}

	// Determine the level of a specific Senzing error.

	messageSplits := strings.Split(message, "|")
	for key, value := range SenzingErrorsMap {
		if messageSplits[0] == key {
			result = value
			break
		}
	}

	// Determine if message has error code.

	return result
}

// Write log record based on message level function.
func LogBasedOnLevel(messageLevel string, messageJson string) {
	logger.LogBasedOnLevel(messageLevel, messageJson)
}

// Write log record based on message level method.
func (logger *Logger) LogBasedOnLevel(messageLevel string, messageJson string) {
	switch messageLevel {
	case "info":
		logger.Info(messageJson)
	case "warning":
		logger.Warn(messageJson)
	case "error":
		logger.Error(messageJson)
	case "debug":
		logger.Debug(messageJson)
	case "trace":
		logger.Trace(messageJson)
	case "retryable":
		logger.Info(messageJson)
	case "reserved":
		logger.Info(messageJson)
	case "fatal":
		logger.Fatal(messageJson)
	case "panic":
		logger.Panic(messageJson)
	default:
		logger.Info(messageJson)
	}
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	return logger.LogMessage(idTemplate, errorNumber, message, details...)
}

// Inspect the error to see what the level is and log based on the level method.
func (logger *Logger) LogMessage(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	var err error = nil
	messageLevel := logger.BuildMessageLevel(errorNumber, message)
	messageJson := logger.BuildMessage(idTemplate, errorNumber, message, details...)
	logger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) error {
	return logger.LogMessageFromError(idTemplate, errorNumber, message, err, details...)
}

// Inspect the error to see what the level is and log based on the level method.
func (logger *Logger) LogMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) error {
	var err error = nil

	messageLevel := logger.BuildMessageLevel(errorNumber, message)
	messageJson := logger.BuildMessageFromError(idTemplate, errorNumber, message, anError, details...)
	logger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) error {
	return logger.LogMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
}

// Inspect the error to see what the level is and log based on the level method.
func (logger *Logger) LogMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) error {
	var err error = nil

	messageLevel := logger.BuildMessageLevel(errorNumber, message)
	messageJson := logger.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, anError, details)
	logger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// Inspect the error to see what the level is and log based on the level function.
func LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
	return logger.LogMessageUsingMap(idTemplate, errorNumber, message, details)
}

// Inspect the error to see what the level is and log based on the level method.
func (logger *Logger) LogMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) error {
	var err error = nil
	messageLevel := logger.BuildMessageLevel(errorNumber, message)
	messageJson := logger.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
	logger.LogBasedOnLevel(messageLevel, messageJson)
	return err
}

// ----------------------------------------------------------------------------
// Public IsXXX() routines
// ----------------------------------------------------------------------------

func IsPanic() bool { return logger.IsPanic() }
func (logger *Logger) IsPanic() bool {
	return logger.isPanic
}

func IsFatal() bool { return logger.IsFatal() }
func (logger *Logger) IsFatal() bool {
	return logger.isFatal
}

func IsError() bool { return logger.IsError() }
func (logger *Logger) IsError() bool {
	return logger.isError
}

func IsWarn() bool { return logger.IsWarn() }
func (logger *Logger) IsWarn() bool {
	return logger.isWarn
}

func IsInfo() bool { return logger.IsInfo() }
func (logger *Logger) IsInfo() bool {
	return logger.isInfo
}

func IsDebug() bool { return logger.IsDebug() }
func (logger *Logger) IsDebug() bool {
	return logger.isDebug
}

func IsTrace() bool { return logger.IsTrace() }
func (logger *Logger) IsTrace() bool {
	return logger.isTrace
}

// ----------------------------------------------------------------------------
// Public XXX() routines for leveled logging.
// ----------------------------------------------------------------------------

// --- Trace ------------------------------------------------------------------

func Trace(v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Trace(v ...interface{}) *Logger {
	if logger.isTrace {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func Tracef(format string, v ...interface{}) *Logger {
	if logger.IsTrace() {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

func (logger *Logger) Tracef(format string, v ...interface{}) *Logger {
	if logger.isTrace {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

// --- Debug ------------------------------------------------------------------

func Debug(v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Debug(v ...interface{}) *Logger {
	if logger.isDebug {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func Debugf(format string, v ...interface{}) *Logger {
	if logger.IsDebug() {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

func (logger *Logger) Debugf(format string, v ...interface{}) *Logger {
	if logger.isDebug {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

// --- Info -------------------------------------------------------------------

func Info(v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Info(v ...interface{}) *Logger {
	if logger.isInfo {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func Infof(format string, v ...interface{}) *Logger {
	if logger.IsInfo() {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

func (logger *Logger) Infof(format string, v ...interface{}) *Logger {
	if logger.isInfo {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

// --- Warn -------------------------------------------------------------------

func Warn(v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Warn(v ...interface{}) *Logger {
	if logger.isWarn {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func Warnf(format string, v ...interface{}) *Logger {
	if logger.IsWarn() {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

func (logger *Logger) Warnf(format string, v ...interface{}) *Logger {
	if logger.isWarn {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}

// --- Error ------------------------------------------------------------------

func Error(v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func (logger *Logger) Error(v ...interface{}) *Logger {
	if logger.isError {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func Errorf(format string, v ...interface{}) *Logger {
	if logger.IsError() {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

func (logger *Logger) Errorf(format string, v ...interface{}) *Logger {
	if logger.isError {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

// --- Fatal ------------------------------------------------------------------

func Fatal(v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func (logger *Logger) Fatal(v ...interface{}) *Logger {
	if logger.isFatal {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func Fatalf(format string, v ...interface{}) *Logger {
	if logger.IsFatal() {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

func (logger *Logger) Fatalf(format string, v ...interface{}) *Logger {
	if logger.isFatal {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

// --- Panic ------------------------------------------------------------------

func Panic(v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func (logger *Logger) Panic(v ...interface{}) *Logger {
	if logger.isPanic {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func Panicf(format string, v ...interface{}) *Logger {
	if logger.IsPanic() {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

func (logger *Logger) Panicf(format string, v ...interface{}) *Logger {
	if logger.isPanic {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}
