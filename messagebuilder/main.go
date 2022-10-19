// The logger package is a set of method to help with common tasks.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagebuilder

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageBuilderImpl struct{}

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageBuilderInterface interface {
	BuildError(idTemplate string, errorNumber int, message string, details ...interface{}) error
	BuildMessage(idTemplate string, errorNumber int, message string, details ...interface{}) string
	BuildMessageFromError(idTemplate string, errorNumber int, message string, anError error, details ...interface{}) string
	BuildMessageFromErrorUsingMap(idTemplate string, errorNumber int, message string, anError error, details map[string]interface{}) string
	BuildMessageUsingMap(idTemplate string, errorNumber int, message string, details map[string]interface{}) string
	BuildMessageId(idTemplate string, errorNumber int) string
	BuildMessageLevel(errorNumber int, message string) string
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var messageBuilderInstance *MessageBuilderImpl

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
