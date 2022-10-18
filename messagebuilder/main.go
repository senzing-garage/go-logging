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

type MessageBuilder struct{}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

var messagebuilder *MessageBuilder

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
