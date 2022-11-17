/*
The MessageLocationSenzing implementation returns a string in the format "In Function() at filename.go:nnn".
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

// The MessageLocationSenzing type is for returning a string in the format "In Function() at filename.go:nnn".
type MessageLocationSenzing struct {
	CallerSkip int // Number of stacks to ascend. See https://pkg.go.dev/runtime#Caller
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// The MessageLocation method returns a string in the format "In Function() at filename.go:nnn".
func (messageLocation *MessageLocationSenzing) MessageLocation(messageNumber int, details ...interface{}) (string, error) {
	var err error = nil
	result := ""

	// Determine number of stacks to ascend.

	callerSkip := messageLocation.CallerSkip
	for _, value := range details {
		switch typedValue := value.(type) {
		case CallerSkip:
			callerSkip = int(typedValue)
		}
	}

	// See https://pkg.go.dev/runtime#Caller

	pc, file, line, ok := runtime.Caller(callerSkip)
	if ok {
		callingFunction := runtime.FuncForPC(pc)
		runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
		functionName := runtimeFunc.ReplaceAllString(callingFunction.Name(), "$1")
		filename := filepath.Base(file)
		result = fmt.Sprintf("In %s() at %s:%d", functionName, filename, line)
	}

	return result, err
}
