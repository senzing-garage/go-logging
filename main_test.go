package main

import (
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/logging"
)

var componentIdentifier = 9999

var idMessagesTest = map[int]string{
	0001: "Info for %s",
	1000: "Warning for %s",
	2000: "Error for %s",
}

var idStatusesTest = map[int]string{
	0001: "Status for 0001",
	1000: "Status for 1000",
}

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestMain(testing *testing.T) {
	main()
}

// ----------------------------------------------------------------------------
// Test interface functions for New - names begin with "Test"
// ----------------------------------------------------------------------------

func TestNew(t *testing.T) {
	logger, _ := logging.New()
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingLogger - names begin with "Test"
// ----------------------------------------------------------------------------

func TestNewSenzingSdkLogger(t *testing.T) {
	logger, _ := logging.NewSenzingSdkLogger(componentIdentifier, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingLoggerAtErrorLevel(t *testing.T) {

	loggerOptions := []interface{}{
		&logging.OptionLogLevel{Value: logger.LevelErrorName},
	}

	logger, _ := logging.NewSenzingSdkLogger(componentIdentifier, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingApiLogger - names begin with "Test"
// ----------------------------------------------------------------------------

// func TestNewSenzingApiLogger(t *testing.T) {

// 	logger, _ := logging.NewSenzingApiLogger(componentIdentifier, idMessagesTest, idStatusesTest)
// 	logger.Log(1, "Mary")
// 	logger.Log(1000, "Jane")
// 	logger.Log(2000, "Bob")
// }

// func TestNewSenzingApiLoggerAtErrorLevel(t *testing.T) {
// 	logger, _ := logging.NewSenzingApiLogger(componentIdentifier, idMessagesTest, idStatusesTest, logger.LevelError)
// 	logger.Log(1, "Mary")
// 	logger.Log(1000, "Jane")
// 	logger.Log(2000, "Bob")
// }
