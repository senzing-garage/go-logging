/*
go-logging/main.go implements examples.
*/
package main

import (
	"errors"
	"fmt"

	"github.com/senzing/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var (
	idMessages = map[int]string{
		0:    "TRACE has %s",
		1000: "DEBUG has %s",
		2000: "INFO has %s",
		3000: "WARN has %s",
		4000: "ERROR has %s",
		5000: "FATAL has %s",
		6000: "PANIC has %s",
	}

	idStatuses = map[int]string{
		2000: "SUCCESS",
		4000: "FAILURE",
		6000: "DISASTER",
	}
	messageIdTemplate = "my-id-%04d"
	callerSkip        = 3
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func testLogger(logger logging.LoggingInterface) {
	err1 := errors.New("example error #1")
	err2 := errors.New("example error #2")

	logger.Log(0000, "TRACE level", err1, err2)
	logger.Log(1000, "DEBUG level", err1, err2)
	logger.Log(2000, "INFO level", err1, err2)
	logger.Log(3000, "WARN level", err1, err2)
	logger.Log(4000, "ERROR level", err1, err2)
	logger.Log(5000, "FATAL level", err1, err2)
	logger.Log(6000, "PANIC level", err1, err2)
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	//

	// ------------------------------------------------------------------------
	// Simple logger
	// Message ids translate into log levels:
	//    0 -  999 TRACE
	// 1000 - 1999 DEBUG
	// 2000 - 2999 INFO
	// 3000 - 3999 WARN
	// 4000 - 4999 ERROR
	// 5000 - 5999 FATAL
	// 6000 - 6999 PANIC
	//
	// Notice that no "text" field shows up.  That's because id messages
	// haven't been defined.   That will be seen in "logger2".
	// ------------------------------------------------------------------------

	logger1, err := logging.New()
	if err != nil {
		fmt.Println(err)
	}

	logger1.Log(2001, "Hello World!")

	testLogger(logger1)

	// ------------------------------------------------------------------------
	// Configured logger
	// ------------------------------------------------------------------------

	loggerOptions2 := []interface{}{
		&logging.OptionIdMessages{Value: idMessages},
		&logging.OptionIdStatuses{Value: idStatuses},
		&logging.OptionMessageIdTemplate{Value: messageIdTemplate},
		&logging.OptionCallerSkip{Value: callerSkip},
	}
	logger2, err := logging.New(loggerOptions2...)
	if err != nil {
		fmt.Println(err)
	}
	testLogger(logger2)

	// ------------------------------------------------------------------------
	// NewSenzingToolsLogger - for use with senzing-tools commands.
	// ------------------------------------------------------------------------

	loggerOptions3 := []interface{}{
		&logging.OptionCallerSkip{Value: callerSkip},
	}
	logger3, err := logging.NewSenzingToolsLogger(9998, idMessages, loggerOptions3...)
	if err != nil {
		fmt.Println(err)
	}
	testLogger(logger3)

	// ------------------------------------------------------------------------
	// README.md examples
	// ------------------------------------------------------------------------

	var (
		ComponentId = 9999            // See https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-component-ids.md
		IdMessages  = map[int]string{ // Message templates. Example: https://github.com/Senzing/init-database/blob/main/senzingconfig/main.go
			2000: "Today's greeting:  %s",
		}
		callerSkip = 3 // Used to determine "location" information. See https://pkg.go.dev/runtime#Caller
	)

	// Logging options. See https://github.com/Senzing/go-logging/blob/main/logging/main.go
	loggerOptions := []interface{}{
		&logging.OptionCallerSkip{Value: callerSkip},
	}

	logger, err := logging.NewSenzingToolsLogger(ComponentId, IdMessages, loggerOptions...)
	if err != nil {
		fmt.Println(err)
	}

	logger.Log(2000, "Hello, world!")

}
