package main

import (
	"bytes"
	"testing"

	"github.com/senzing-garage/go-logging/logger"
	"github.com/senzing-garage/go-logging/logging"
	"github.com/stretchr/testify/assert"
)

const (
	componentIdentifier = 9999
)

var idMessagesTest = map[int]string{ //nolint
	1:    "Info for %s",
	1000: "Warning for %s",
	2000: "Error for %s",
}

func TestMain(test *testing.T) {
	test.Parallel()
	main()
}

// ----------------------------------------------------------------------------
// Test interface functions for New
// ----------------------------------------------------------------------------

func TestNewLogger(test *testing.T) {
	test.Parallel()

	expected := `{"level":"INFO","id":"2000"}` + "\n"
	outputString := new(bytes.Buffer)
	logger, _ := logging.New(optionOutput(outputString), optionTimeHidden())
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
	assert.Equal(test, expected, outputString.String())
}

// ----------------------------------------------------------------------------
// Test interface functions for NewSenzingLogger
// ----------------------------------------------------------------------------

func TestNewSenzingLogger(test *testing.T) {
	test.Parallel()

	expected := `{"level":"INFO","id":"SZTL99992000"}` + "\n"
	outputString := new(bytes.Buffer)
	logger, _ := logging.NewSenzingLogger(
		componentIdentifier,
		idMessagesTest,
		optionOutput(outputString),
		optionTimeHidden(),
	)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
	assert.Equal(test, expected, outputString.String())
}

func TestNewSenzingLoggerAtErrorLevel(test *testing.T) {
	test.Parallel()

	expected := ""
	outputString := new(bytes.Buffer)
	loggerOptions := []interface{}{
		logging.OptionLogLevel{Value: logger.LevelErrorName},
		optionOutput(outputString),
		logging.OptionTimeHidden{Value: true},
	}
	logger, _ := logging.NewSenzingLogger(componentIdentifier, idMessagesTest, loggerOptions...)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
	assert.Equal(test, expected, outputString.String())
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func optionOutput(outputString *bytes.Buffer) logging.OptionOutput {
	return logging.OptionOutput{
		Value: outputString,
	}
}

func optionTimeHidden() logging.OptionTimeHidden {
	return logging.OptionTimeHidden{
		Value: true,
	}
}
