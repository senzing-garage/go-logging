package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/senzing-garage/go-logging/logging"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var (
	callerSkip = 3
	err1       = errors.New("example error #1")
	err2       = errors.New("example error #2")
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
	messageReason           = logging.MessageReason{Value: "The reason is..."}
	optionCallerSkip        = logging.OptionCallerSkip{Value: callerSkip}
	optionIDMessages        = logging.OptionIDMessages{Value: idMessages}
	optionIDStatuses        = logging.OptionIDStatuses{Value: idStatuses}
	optionMessageID         = "my-id-%04d"
	optionMessageIDTemplate = logging.OptionMessageIDTemplate{Value: optionMessageID}
)

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
	testError(err)
	logger1.Log(2001)
	testLogger("Simple logger", logger1)

	// ------------------------------------------------------------------------
	// Configured logger
	// ------------------------------------------------------------------------

	loggerOptions2 := []interface{}{
		optionIDMessages,
		optionIDStatuses,
		optionMessageIDTemplate,
		optionCallerSkip,
		logging.OptionMessageFields{Value: []string{"id", "text", "reason"}},
	}
	logger2, err := logging.New(loggerOptions2...)
	testError(err)
	testLogger("Configured logger", logger2)

	// ------------------------------------------------------------------------
	// NewSenzingLogger - for use with senzing-tools commands.
	// ------------------------------------------------------------------------

	logger3, err := logging.NewSenzingLogger(9998, idMessages)
	testError(err)
	testLogger("SenzingLogger", logger3)

	// ------------------------------------------------------------------------
	// README.md examples
	// ------------------------------------------------------------------------

	var (
		ComponentID = 9999            // See https://github.com/senzing-garage/knowledge-base/blob/main/lists/senzing-component-ids.md
		IDMessages  = map[int]string{ // Message templates. Example: https://github.com/senzing-garage/init-database/blob/main/senzingconfig/main.go
			2000: "Today's greeting:  %s",
			4000: "Here's what happened: %s",
		}
		callerSkip = 3 // Used to determine "location" information. See https://pkg.go.dev/runtime#Caller
	)

	printBanner("README.md examples")

	// Logging options. See https://github.com/senzing-garage/go-logging/blob/main/logging/main.go

	loggerOptions4 := []interface{}{
		logging.OptionCallerSkip{Value: callerSkip},
	}
	logger4, err := logging.NewSenzingLogger(ComponentID, IDMessages, loggerOptions4...)
	testError(err)
	logger4.Log(2000, "Hello, world!")
	err = logger4.NewError(4000, "A bad thing")
	fmt.Printf("The error: %v\n", err)

	// ------------------------------------------------------------------------
	// docs/examples.md examples
	// ------------------------------------------------------------------------

	printBanner("docs/examples.md examples")

	// logger5

	logger5, _ := logging.New()
	logger5.Log(2001, "A message")

	// logger6

	loggerOptions6 := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
	}
	logger6, _ := logging.New(loggerOptions6...)
	logger6.Log(999, "TRACE level")
	logger6.Log(1000, "DEBUG level")
	logger6.Log(2000, "INFO  level")
	logger6.Log(3000, "WARN  level")
	logger6.Log(4000, "ERROR level")
	logger6.Log(5000, "FATAL level")
	logger6.Log(6000, "PANIC level")
	logger6.Log(7000, "undefined level")
	logger6.Log(8000, "undefined level")

	// logger7

	loggerOptions7 := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
		logging.OptionLogLevel{Value: "TRACE"},
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

	// logger8

	loggerOptions8 := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
		logging.OptionMessageIDTemplate{Value: "my-message-%04d"},
	}
	logger8, _ := logging.New(loggerOptions8...)
	logger8.Log(2002, "A message")

	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}

	aStruct := struct {
		Name string
		ID   int
	}{
		Name: "Robert Smith",
		ID:   123145,
	}

	logger8.Log(2003, "Robert Smith", 12345, aMap, aStruct)

	// logger9

	idMessages := map[int]string{
		999:  "A test of TRACE.",
		1000: "A test of DEBUG.",
		2000: "A test of INFO.",
		2004: "The favorite number for %s is %d.",
		3000: "A test of WARN.",
		4000: "A test of ERROR.",
		5000: "A test of FATAL.",
		6000: "A test of PANIC.",
	}
	loggerOptions9 := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "text"}},
		logging.OptionIDMessages{Value: idMessages},
	}
	logger9, _ := logging.New(loggerOptions9...)
	logger9.Log(2004, "Robert Smith", 12345)

	// logger10

	err1 := errors.New("error #1")
	err2 := errors.New("error #2")

	loggerOptions10 := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
		logging.OptionIDMessages{Value: idMessages},
	}
	logger10, _ := logging.New(loggerOptions10...)
	logger10.Log(2005, err1, err2)

	// Epilog.

	printBanner("Done")
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func printBanner(banner string) {
	fmt.Printf("\n%s\n", strings.Repeat("-", 80))
	fmt.Printf("-- %s\n", banner)
	fmt.Printf("%s\n\n", strings.Repeat("-", 80))
}

func testError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func testLogger(banner string, logger logging.Logging) {
	printBanner(banner)

	// Test logging.

	logger.Log(0000, "TRACE level", messageReason, err1, err2)
	logger.Log(1000, "DEBUG level", messageReason, err1, err2)
	logger.Log(2000, "INFO level", messageReason, err1, err2)
	logger.Log(3000, "WARN level", messageReason, err1, err2)
	logger.Log(4000, "ERROR level", messageReason, err1, err2)
	logger.Log(5000, "FATAL level", messageReason, err1, err2)
	logger.Log(6000, "PANIC level", messageReason, err1, err2)
}
