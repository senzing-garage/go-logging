package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/senzing-garage/go-logging/logger"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/stretchr/testify/assert"
)

var (
	componentIdentifier = 9999
	idMessagesTest      = map[int]string{
		0001: "Info for %s",
		1000: "Warning for %s",
		2000: "Error for %s",
	}
	outputString = new(bytes.Buffer) // *bytes.Buffer
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
	expected := `{"level":"INFO","id":"2000"}` + "\n"
	outputString.Reset()
	logger, _ := logging.New(optionOutput(), optionTimeHidden())
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
	assert.Equal(test, expected, outputString.String())
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingLogger
// ----------------------------------------------------------------------------

func TestNewSenzingLogger(test *testing.T) {
	_ = test
	expected := `{"level":"INFO","id":"SZTL99992000"}` + "\n"
	outputString.Reset()
	logger, _ := logging.NewSenzingLogger(componentIdentifier, idMessagesTest, optionOutput(), optionTimeHidden())
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
	assert.Equal(test, expected, outputString.String())
}

func TestNewSenzingLoggerAtErrorLevel(test *testing.T) {
	_ = test
	expected := ""
	outputString.Reset()
	loggerOptions := []interface{}{
		logging.OptionLogLevel{Value: logger.LevelErrorName},
		logging.OptionOutput{Value: outputString},
		logging.OptionTimeHidden{Value: true},
	}
	logger, _ := logging.NewSenzingLogger(componentIdentifier, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
	assert.Equal(test, expected, outputString.String())

	fmt.Println(">>>> ", outputString)

}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func optionOutput() logging.OptionOutput {
	return logging.OptionOutput{
		Value: outputString,
	}
}

func optionTimeHidden() logging.OptionTimeHidden {
	return logging.OptionTimeHidden{
		Value: true,
	}
}
