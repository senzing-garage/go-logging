/*
Package helper ...
*/
package messagestatus

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageStatusSenzingApi struct{}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var senzingApiErrorsMap = map[string]string{
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

// Important:  The number listed is one more than the highest number for the MessageLevel.
// Message ranges:
// 0000-0999 info
// 1000-1999 warning
// 2000-2999 error
// 3000-3999 debug
// 4000-4999 trace
// 5000-5999 fatal
// 6000-6999 panic
var MessageLevelMap = map[int]logger.Level{
	1000: logger.LevelInfo,
	2000: logger.LevelWarn,
	3000: logger.LevelError,
	4000: logger.LevelDebug,
	5000: logger.LevelTrace,
	6000: logger.LevelFatal,
	7000: logger.LevelPanic,
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageStatusSenzingApi) MessageStatus(errorNumber int, details ...interface{}) (string, error) {
	var err error = nil
	var result = ""

	mapIndex := "0000E"

	result, ok := senzingApiErrorsMap[mapIndex]
	if ok {
		return result, err
	}

	return result, err
}
