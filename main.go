package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/senzing-garage/go-logging/logging"
)

const (
	horizontalRuleLength = 80
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var (
	err1       = errors.New("example error #1")
	err2       = errors.New("example error #2")
	idMessages = map[int]string{ //nolint
		0:    "TRACE has %s",
		1000: "DEBUG has %s",
		2000: "INFO has %s",
		3000: "WARN has %s",
		4000: "ERROR has %s",
		5000: "FATAL has %s",
		6000: "PANIC has %s",
	}
	idStatuses = map[int]string{ //nolint
		2000: "SUCCESS",
		4000: "FAILURE",
		6000: "DISASTER",
	}
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
	logger01()
	logger02()
	logger03()

	// ------------------------------------------------------------------------
	// README.md examples
	// ------------------------------------------------------------------------

	printBanner("README.md examples")
	logger04()

	// ------------------------------------------------------------------------
	// docs/examples.md examples
	// ------------------------------------------------------------------------

	printBanner("docs/examples.md examples")

	logger05()
	logger06()
	logger07()
	logger08()
	logger09()
	logger10()

	printBanner("Done")
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func logger01() {
	logger, err := logging.New()
	testError(err)
	logger.Log(2001)
	testLogger1("Simple logger", logger)
}

func logger02() {
	callerSkip := 3
	optionCallerSkip := logging.OptionCallerSkip{Value: callerSkip}
	optionIDMessages := logging.OptionIDMessages{Value: idMessages}
	optionIDStatuses := logging.OptionIDStatuses{Value: idStatuses}
	optionMessageIDTemplate := logging.OptionMessageIDTemplate{Value: "my-id-%04d"}

	loggerOptions := []interface{}{
		optionIDMessages,
		optionIDStatuses,
		optionMessageIDTemplate,
		optionCallerSkip,
		logging.OptionMessageFields{Value: []string{"id", "text", "reason"}},
	}
	logger, err := logging.New(loggerOptions...)
	testError(err)
	testLogger1("Configured logger", logger)
}

func logger03() {
	logger, err := logging.NewSenzingLogger(9998, idMessages)
	testError(err)
	testLogger1("SenzingLogger", logger)
}

func logger04() {
	var (
		ComponentID = 9999            // See https://github.com/senzing-garage/knowledge-base/blob/main/lists/senzing-component-ids.md
		IDMessages  = map[int]string{ // Message templates. Example: https://github.com/senzing-garage/init-database/blob/main/senzingconfig/main.go
			2000: "Today's greeting:  %s",
			4000: "Here's what happened: %s",
		}
		callerSkip = 3 // Used to determine "location" information. See https://pkg.go.dev/runtime#Caller
	)

	loggerOptions := []interface{}{
		logging.OptionCallerSkip{Value: callerSkip},
	}
	logger, err := logging.NewSenzingLogger(ComponentID, IDMessages, loggerOptions...)
	testError(err)
	logger.Log(2000, "Hello, world!")
	err = logger.NewError(4000, "A bad thing")
	outputf("The error: %v\n", err)
}

func logger05() {
	logger, _ := logging.New()
	logger.Log(2001, "A message")
}

func logger06() {
	loggerOptions := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
	}
	logger, _ := logging.New(loggerOptions...)
	testLogger2("Logger test 6", logger)
}

func logger07() {
	loggerOptions := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
		logging.OptionLogLevel{Value: "TRACE"},
	}
	logger, _ := logging.New(loggerOptions...)
	testLogger2("Logger test 7", logger)
}

func logger08() {
	loggerOptions := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
		logging.OptionMessageIDTemplate{Value: "my-message-%04d"},
	}
	logger, _ := logging.New(loggerOptions...)
	logger.Log(2002, "A message")

	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}

	aStruct := struct {
		Name string
		ID   int
	}{
		Name: "Robert Smith",
		ID:   123145, //nolint
	}

	logger.Log(2003, "Robert Smith", 12345, aMap, aStruct)
}

func logger09() {
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
	loggerOptions := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "text"}},
		logging.OptionIDMessages{Value: idMessages},
	}
	logger, _ := logging.New(loggerOptions...)
	logger.Log(2004, "Robert Smith", 12345)
}

func logger10() {
	loggerOptions := []interface{}{
		logging.OptionMessageFields{Value: []string{"id", "details"}},
		logging.OptionIDMessages{Value: idMessages},
	}
	logger, _ := logging.New(loggerOptions...)
	logger.Log(2005, err1, err2)
}

func outputf(format string, message ...any) {
	fmt.Printf(format, message...) //nolint
}

func outputln(message ...any) {
	fmt.Println(message...) //nolint
}

func printBanner(banner string) {
	outputf("\n%s\n", strings.Repeat("-", horizontalRuleLength))
	outputf("-- %s\n", banner)
	outputf("%s\n\n", strings.Repeat("-", horizontalRuleLength))
}

func testError(err error) {
	if err != nil {
		outputln(err)
	}
}

func testLogger1(banner string, logger logging.Logging) {
	messageReason := logging.MessageReason{Value: "The reason is..."}

	printBanner(banner)

	// Test logging.

	logger.Log(0, "TRACE level", messageReason, err1, err2)
	logger.Log(1000, "DEBUG level", messageReason, err1, err2)
	logger.Log(2000, "INFO level", messageReason, err1, err2)
	logger.Log(3000, "WARN level", messageReason, err1, err2)
	logger.Log(4000, "ERROR level", messageReason, err1, err2)
	logger.Log(5000, "FATAL level", messageReason, err1, err2)
	logger.Log(6000, "PANIC level", messageReason, err1, err2)
}

func testLogger2(banner string, logger logging.Logging) {
	printBanner(banner)

	logger.Log(999, "TRACE level")
	logger.Log(1000, "DEBUG level")
	logger.Log(2000, "INFO  level")
	logger.Log(3000, "WARN  level")
	logger.Log(4000, "ERROR level")
	logger.Log(5000, "FATAL level")
	logger.Log(6000, "PANIC level")
	logger.Log(7000, "undefined level")
	logger.Log(8000, "undefined level")
}
