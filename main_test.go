package main

import (
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/logging"
)

var (
	componentIdentifier   = 9999
	testMessageIdTemplate = "test-%04d"

	idMessagesTest = map[int]string{
		0001: "Info for %s",
		1000: "Warning for %s",
		2000: "Error for %s",
	}

	idStatusesTest = map[int]string{
		0001: "Status for 0001",
		1000: "Status for 1000",
	}
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestMain(testing *testing.T) {
	main()
}

// ----------------------------------------------------------------------------
// Test interface functions for New
// ----------------------------------------------------------------------------

func TestNew(t *testing.T) {
	logger, _ := logging.New()
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingLogger
// ----------------------------------------------------------------------------

func TestNewSenzingLogger(t *testing.T) {
	logger, _ := logging.NewSenzingLogger(testMessageIdTemplate, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingLoggerAtErrorLevel(t *testing.T) {

	loggerOptions := []interface{}{
		&logging.OptionLogLevel{Value: logger.LevelErrorName},
	}

	logger, _ := logging.NewSenzingLogger(testMessageIdTemplate, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingSdkLogger
// ----------------------------------------------------------------------------

func TestNewSenzingSdkLogger(t *testing.T) {
	logger, _ := logging.NewSenzingSdkLogger(componentIdentifier, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingSdkLoggerAtErrorLevel(t *testing.T) {

	loggerOptions := []interface{}{
		&logging.OptionLogLevel{Value: logger.LevelErrorName},
	}

	logger, _ := logging.NewSenzingSdkLogger(componentIdentifier, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingToolsLogger
// ----------------------------------------------------------------------------

func TestNewSenzingToolsLogger(t *testing.T) {
	logger, _ := logging.NewSenzingToolsLogger(componentIdentifier, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingToolsLoggerAtErrorLevel(t *testing.T) {

	loggerOptions := []interface{}{
		&logging.OptionLogLevel{Value: logger.LevelErrorName},
	}

	logger, _ := logging.NewSenzingToolsLogger(componentIdentifier, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}
