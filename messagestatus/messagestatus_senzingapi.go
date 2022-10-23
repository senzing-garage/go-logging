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

var senzingApiStatusMap = map[int]string{
	1000: "user-input-error",
	2000: "retryable",
}

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

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// TODO:
func (messagelevel *MessageStatusSenzingApi) MessageStatus(errorNumber int, text string) (string, error) {
	var err error = nil
	var result = ""

	result, ok := senzingApiStatusMap[errorNumber]
	if ok {
		return result, err
	}

	return result, err
}
