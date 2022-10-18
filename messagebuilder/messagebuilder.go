/*
Package helper ...
*/
package messagebuilder

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func init() {
	messagebuilder = New()
}

// ----------------------------------------------------------------------------
// Public Setters
// ----------------------------------------------------------------------------

func New() *MessageBuilder {
	return new(MessageBuilder)
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build an error function.
func BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	return messagebuilder.BuildError(idTemplate, errorNumber, message, details...)
}

// Build an error method.
func (messagebuilder *MessageBuilder) BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error {
	errorMessage := messageformat.BuildMessage(
		messagebuilder.BuildMessageId(idTemplate, errorNumber),
		messagebuilder.BuildMessageLevel(errorNumber, message),
		message,
		details...,
	)
	return errors.New(errorMessage)
}

// Build log message function.
func BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string {
	return messagebuilder.BuildMessage(idTemplate, errorNumber, message, details...)
}

// Build log message method.
func (messagebuilder *MessageBuilder) BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string {
	return messageformat.BuildMessage(
		messagebuilder.BuildMessageId(idTemplate, errorNumber),
		messagebuilder.BuildMessageLevel(errorNumber, message),
		message,
		details...,
	)
}

// Build log message function.
func BuildMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) string {
	return messagebuilder.BuildMessageFromError(idTemplate, errorNumber, message, err, details...)
}

// Build log message method.
func (messagebuilder *MessageBuilder) BuildMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) string {
	return messageformat.BuildMessageFromError(
		messagebuilder.BuildMessageId(idTemplate, errorNumber),
		messagebuilder.BuildMessageLevel(errorNumber, message),
		message,
		anError,
		details...,
	)
}

// Build log message function.
func BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) string {
	return messagebuilder.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
}

// Build log message method.
func (messagebuilder *MessageBuilder) BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) string {
	return messageformat.BuildMessageFromErrorUsingMap(
		messagebuilder.BuildMessageId(idTemplate, errorNumber),
		messagebuilder.BuildMessageLevel(errorNumber, message),
		message,
		anError,
		details,
	)
}

// Build log message function.
func BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string {
	return messagebuilder.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
}

// Build log message method.
func (messagebuilder *MessageBuilder) BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string {
	return messageformat.BuildMessageUsingMap(
		messagebuilder.BuildMessageId(idTemplate, errorNumber),
		messagebuilder.BuildMessageLevel(errorNumber, message),
		message,
		details,
	)
}

// Construct a unique message id function.
func BuildMessageId(idTemplate string, errorNumber int) string {
	return messagebuilder.BuildMessageId(idTemplate, errorNumber)
}

// Construct a unique message id method.
func (messagebuilder *MessageBuilder) BuildMessageId(idTemplate string, errorNumber int) string {
	return fmt.Sprintf(idTemplate, errorNumber)
}

// Based on the errorNumber and Senzing error code, get the message level function.
func BuildMessageLevel(errorNumber int, message string) string {
	return messagebuilder.BuildMessageLevel(errorNumber, message)
}

// Based on the errorNumber and Senzing error code, get the message level method.
func (messagebuilder *MessageBuilder) BuildMessageLevel(errorNumber int, message string) string {

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
