/*
The MessageLocationDefault implementation returns an empty string for a location value.
*/
package messagelocation

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The MessageLocationDefault type is for returning an empty string for location value.
type MessageLocationDefault struct {
	CallerSkip int
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLocation method returns an empty string for a location value.
func (messageLocation *MessageLocationDefault) MessageLocation(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	result := ""
	// https://pkg.go.dev/runtime#Caller
	pc, file, line, ok := runtime.Caller(messageLocation.CallerSkip)
	if ok {
		callingFunction := runtime.FuncForPC(pc)
		runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
		functionName := runtimeFunc.ReplaceAllString(callingFunction.Name(), "$1")
		filename := filepath.Base(file)
		result = fmt.Sprintf("In %s() at %s:%d", functionName, filename, line)
	}

	return result, err
}
