package main

import (
	"testing"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagelogger"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestMain(testing *testing.T) {
	main()
}

func TestNewSenzingLogger(t *testing.T) {

	// func NewSenzingLogger(productIdentifier int, idMessages map[int]string, interfaces ...interface{}) (MessageLoggerInterface, error) {

	productIdentifier := 9999

	idMessages := map[int]string{
		0001: "Info for %s",
		1000: "Warning for %s",
		2000: "Error for %s",
	}

	logger, _ := messagelogger.NewSenzingLogger(productIdentifier, idMessages)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}

func TestNewSenzingLoggerWithOnlyError(t *testing.T) {

	// func NewSenzingLogger(productIdentifier int, idMessages map[int]string, interfaces ...interface{}) (MessageLoggerInterface, error) {

	productIdentifier := 9999

	idMessages := map[int]string{
		0001: "Info for %s",
		1000: "Warning for %s",
		2000: "Error for %s",
	}

	logger, _ := messagelogger.NewSenzingLogger(productIdentifier, idMessages, logger.LevelError)
	logger.Log(1, "Mary")
	logger.Log(1000, "Jane")
	logger.Log(2000, "Bob")
}
