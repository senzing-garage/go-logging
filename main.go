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
	// NewSenzingLogger - for use generally.
	// ------------------------------------------------------------------------

	logger4, err := logging.NewSenzingLogger("my-unique-%04d", idMessages)
	if err != nil {
		fmt.Println(err)
	}
	testLogger(logger4)

	// ------------------------------------------------------------------------
	// README.md examples
	// ------------------------------------------------------------------------

	var (
		ComponentId = 9999            // See https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-component-ids.md
		IdMessages  = map[int]string{ // Message templates. Example: https://github.com/Senzing/init-database/blob/main/senzingconfig/main.go
			2000: "Today's greeting:  %s",
			4000: "Here's what happened: %s",
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

	err = logger.NewError(4000, "A bad thing")
	fmt.Printf("The error: %v\n", err)

	// ------------------------------------------------------------------------
	// doc.go examples
	// ------------------------------------------------------------------------

	logger5, _ := logging.New()
	logger5.Log(2000, "A message")

	logger6, _ := logging.New()
	logger6.Log(999, "TRACE level")
	logger6.Log(1000, "DEBUG level")
	logger6.Log(2000, "INFO  level")
	logger6.Log(3000, "WARN  level")
	logger6.Log(4000, "ERROR level")
	logger6.Log(5000, "FATAL level")
	logger6.Log(6000, "PANIC level")
	logger6.Log(7000, "undefined level")
	logger6.Log(8000, "undefined level")

	loggerOptions7 := []interface{}{
		&logging.OptionLogLevel{Value: "TRACE"},
	}
	logger7, _ := logging.New(loggerOptions7...)
	logger7.Log(999, "TRACE level")
	logger7.Log(1000, "DEBUG level")
	logger7.Log(2000, "INFO  level")
	logger7.Log(3000, "WARN  level")
	logger7.Log(4000, "ERROR level")
	logger7.Log(5000, "FATAL level")
	logger7.Log(6000, "PANIC level")
	logger7.Log(7000, "undefined level")
	logger7.Log(8000, "undefined level")

	loggerOptions8 := []interface{}{
		&logging.OptionMessageIdTemplate{Value: "my-message-%04d"},
	}
	logger8, _ := logging.New(loggerOptions8...)
	logger8.Log(2001, "A message")
	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}
	logger8.Log(2002, "Robert Smith", 12345, aMap)

	idMessages := map[int]string{
		999:  "A test of TRACE.",
		1000: "A test of DEBUG.",
		2000: "A test of INFO.",
		2003: "The favorite number for %s is %d.",
		3000: "A test of WARN.",
		4000: "A test of ERROR.",
		5000: "A test of FATAL.",
		6000: "A test of PANIC.",
	}
	loggerOptions9 := []interface{}{
		&logging.OptionIdMessages{Value: idMessages},
	}
	logger9, _ := logging.New(loggerOptions9...)
	logger9.Log(2003, "Robert Smith", 12345, aMap)

	err1 := errors.New("error #1")
	err2 := errors.New("error #2")
	logger9.Log(2010, err1, err2)

}
