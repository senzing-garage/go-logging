package main

import (
	"testing"

	"github.com/senzing-garage/go-logging/logger"
	"github.com/senzing-garage/go-logging/logging"
)

var (
	componentIdentifier   = 9999
	testMessageIDTemplate = "test-%04d"
	idMessagesTest        = map[int]string{
		0001: "Info for %s",
		1000: "Warning for %s",
		2000: "Error for %s",
	}
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestMain(test *testing.T) {
	_ = test
	main()
}

// ----------------------------------------------------------------------------
// Test interface functions for New
// ----------------------------------------------------------------------------

func TestNew(test *testing.T) {
	_ = test
	logger, _ := logging.New()
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingLogger
// ----------------------------------------------------------------------------

func TestNewSenzingLogger(test *testing.T) {
	_ = test
	logger, _ := logging.NewSenzingLogger(testMessageIDTemplate, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingLoggerAtErrorLevel(test *testing.T) {
	_ = test
	loggerOptions := []interface{}{
		&logging.OptionLogLevel{Value: logger.LevelErrorName},
	}
	logger, _ := logging.NewSenzingLogger(testMessageIDTemplate, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingSdkLogger
// ----------------------------------------------------------------------------

func TestNewSenzingSdkLogger(test *testing.T) {
	_ = test
	logger, _ := logging.NewSenzingSdkLogger(componentIdentifier, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingSdkLoggerAtErrorLevel(test *testing.T) {
	_ = test
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

func TestNewSenzingToolsLogger(test *testing.T) {
	_ = test
	logger, _ := logging.NewSenzingToolsLogger(componentIdentifier, idMessagesTest)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingToolsLoggerAtErrorLevel(test *testing.T) {
	_ = test
	loggerOptions := []interface{}{
		&logging.OptionLogLevel{Value: logger.LevelErrorName},
	}
	logger, _ := logging.NewSenzingToolsLogger(componentIdentifier, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}
