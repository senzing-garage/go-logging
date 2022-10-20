/*
Package helper ...
*/
package messagebuilder

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageBuilderImpl struct{}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// var messageBuilderInstance *MessageBuilderImpl

// Important:  The number listed is one more than the highest number for the MessageLevel.
// Message ranges:
// 0000-0999 info
// 1000-1999 warning
// 2000-2999 error
// 3000-3999 debug
// 4000-4999 trace
// 5000-5999 reserved-
// 6000-6999 retryable
// 7000-7999 reserved-2
// 8000-8999 fatal
// 9000-9999 panic
var MessageLevelMap = map[int]string{
	1000:  "info",
	2000:  "warning",
	3000:  "error",
	4000:  "debug",
	5000:  "trace",
	7000:  "retryable",
	9000:  "reserved",
	10000: "fatal",
}

var SenzingErrorsMap = map[string]string{
	"0002E":  "info",
	"0007E":  "error",
	"0023E":  "error",
	"0024E":  "error",
	"0025E":  "error",
	"0026E":  "error",
	"0027E":  "error",
	"0032E":  "error",
	"0034E":  "error",
	"0035E":  "error",
	"0036E":  "error",
	"0048E":  "fatal",
	"0051E":  "error",
	"0053E":  "fatal",
	"0054E":  "error",
	"0061E":  "error",
	"0062E":  "error",
	"0064E":  "error",
	"1007E":  "error",
	"2134E":  "error",
	"30020":  "error",
	"30103E": "error",
	"30110E": "error",
	"30111E": "error",
	"30112E": "error",
	"30121E": "error",
	"30122E": "error",
	"30123E": "error",
	"9000E":  "error",
}

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

func New() *MessageBuilderImpl {
	return new(MessageBuilderImpl)
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

// func init() {
// 	messageBuilderInstance = New()
// }

// ----------------------------------------------------------------------------
// Instance functions
// ----------------------------------------------------------------------------

// Build an error function.
// func BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error {
// 	return messageBuilderInstance.BuildError(idTemplate, errorNumber, message, details...)
// }

// Build log message function.
// func BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string {
// 	return messageBuilderInstance.BuildMessage(idTemplate, errorNumber, message, details...)
// }

// Build log message function.
// func BuildMessageFromError(idTemplate string, errorNumber int, message string, err error, details ...interface{}) string {
// 	return messageBuilderInstance.BuildMessageFromError(idTemplate, errorNumber, message, err, details...)
// }

// Build log message function.
// func BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, err error, details map[string]interface{}) string {
// 	return messageBuilderInstance.BuildMessageFromErrorUsingMap(idTemplate, errorNumber, message, err, details)
// }

// Build log message function.
// func BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string {
// 	return messageBuilderInstance.BuildMessageUsingMap(idTemplate, errorNumber, message, details)
// }

// Construct a unique message id function.
// func BuildMessageId(idTemplate string, errorNumber int) string {
// 	return messageBuilderInstance.BuildMessageId(idTemplate, errorNumber)
// }

// Based on the errorNumber and Senzing error code, get the message level function.
// func BuildMessageLevel(errorNumber int, message string) string {
// 	return messageBuilderInstance.BuildMessageLevel(errorNumber, message)
// }

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build an error method.
// func (messagebuilder *MessageBuilderImpl) BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error {
// 	errorMessage := messageformat.BuildMessage(
// 		messagebuilder.BuildMessageId(idTemplate, errorNumber),
// 		messagebuilder.BuildMessageLevel(errorNumber, message),
// 		message,
// 		details...,
// 	)
// 	return errors.New(errorMessage)
// }

// Build log message method.
// func (messagebuilder *MessageBuilderImpl) BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string {
// 	return messageformat.BuildMessage(
// 		messagebuilder.BuildMessageId(idTemplate, errorNumber),
// 		messagebuilder.BuildMessageLevel(errorNumber, message),
// 		message,
// 		details...,
// 	)
// }

// Build log message method.
// func (messagebuilder *MessageBuilderImpl) BuildMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) string {
// 	return messageformat.BuildMessageFromError(
// 		messagebuilder.BuildMessageId(idTemplate, errorNumber),
// 		messagebuilder.BuildMessageLevel(errorNumber, message),
// 		message,
// 		anError,
// 		details...,
// 	)
// }

// Build log message method.
// func (messagebuilder *MessageBuilderImpl) BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) string {
// 	return messageformat.BuildMessageFromErrorUsingMap(
// 		messagebuilder.BuildMessageId(idTemplate, errorNumber),
// 		messagebuilder.BuildMessageLevel(errorNumber, message),
// 		message,
// 		anError,
// 		details,
// 	)
// }

// Build log message method.
// func (messagebuilder *MessageBuilderImpl) BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string {
// 	return messageformat.BuildMessageUsingMap(
// 		messagebuilder.BuildMessageId(idTemplate, errorNumber),
// 		messagebuilder.BuildMessageLevel(errorNumber, message),
// 		message,
// 		details,
// 	)
// }

// Construct a unique message id method.
// func (messagebuilder *MessageBuilderImpl) BuildMessageId(idTemplate string, errorNumber int) string {
// 	return fmt.Sprintf(idTemplate, errorNumber)
// }

// Based on the errorNumber and Senzing error code, get the message level method.
// func (messagebuilder *MessageBuilderImpl) BuildMessageLevel(errorNumber int, message string) string {

// 	var result = "unknown"

// 	// Create a list of sorted keys.

// 	messageLevelKeys := make([]int, 0, len(MessageLevelMap))
// 	for key := range MessageLevelMap {
// 		messageLevelKeys = append(messageLevelKeys, key)
// 	}
// 	sort.Ints(messageLevelKeys)

// 	// Using the sorted message number, find the level.

// 	for _, messageLevelKey := range messageLevelKeys {
// 		if errorNumber < messageLevelKey {

// 			result = MessageLevelMap[messageLevelKey]
// 			break
// 		}
// 	}

// 	// Determine the level of a specific Senzing error.

// 	messageSplits := strings.Split(message, "|")
// 	for key, value := range SenzingErrorsMap {
// 		if messageSplits[0] == key {
// 			result = value
// 			break
// 		}
// 	}

// 	// Determine if message has error code.

// 	return result
// }
