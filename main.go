package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messageloglevel"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

const MessageIdFormat = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func complexProcess() string {
	time.Sleep(1000 * time.Second)
	return "slept"
}

func boilerplateLogging(aLogger messagelogger.MessageLoggerInterface) {
	aLogger.Log(0, "Custom message")
	aLogger.Log(1000, programName, buildVersion, buildIteration)
	aLogger.Log(2000, programName, buildVersion, buildIteration)
	aLogger.Log(3000, programName, buildVersion, buildIteration)
	aLogger.Log(4000, programName, buildVersion, buildIteration)
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	// Configure the "log" standard library.

	log.SetFlags(log.Llongfile | log.Ldate | log.Lmicroseconds | log.LUTC)

	// ------------------------------------------------------------------------
	// The following demonstrates the high-level messagelogger calls for
	// LogMessage, LogMessageUsingMap, and LogMessageFromError.
	// ------------------------------------------------------------------------

	// --- Simple case with default MessageFormat, no Messages, no MessageLevel

	fmt.Printf("\n\n--- Test 1: No customization ---------------------------------------------------\n\n")
	boilerplateLogging(messagelogger.GetMessageLogger())

	fmt.Printf("\n\n--- Test 2: Add customized id --------------------------------------------------\n\n")
	messagelogger.GetMessageLogger().SetIdTemplate("senzing-9999%04d")
	boilerplateLogging(messagelogger.GetMessageLogger())

	fmt.Printf("\n\n--- Test 3: Add text -----------------------------------------------------------\n\n")

	var messageTemplates = map[int]string{
		0:    "No variable substitution",
		1000: "Program name: %s;",
		2000: "Program name: %s; Build version %s;",
		3000: "Program name: %s; Build version %s; Build iterations %s;",
		4000: "Program name: %s; Build version %s; Build iterations %s; Unknown: %s",
	}
	messagelogger.GetMessageLogger().SetTextTemplates(messageTemplates)
	boilerplateLogging(messagelogger.GetMessageLogger())

	fmt.Printf("\n\n--- Test 4: Add log levels -----------------------------------------------------\n\n")

	messagelogger.GetMessageLogger().MessageLogLevel = &messageloglevel.MessageLogLevelSenzingApi{}
	boilerplateLogging(messagelogger.GetMessageLogger())

	fmt.Printf("\n\n--- Test 5: Add status ---------------------------------------------------------\n\n")

	messagelogger.GetMessageLogger().MessageStatus = &messagestatus.MessageStatusSenzingApi{}
	boilerplateLogging(messagelogger.GetMessageLogger())

	fmt.Printf("\n\n--- Test 6: Add logging golang errors ------------------------------------------\n\n")

	error_1 := errors.New("first error")
	error_2 := errors.New("second error")

	messagelogger.Log(0, "Custom message", error_1)
	messagelogger.Log(1000, programName, buildVersion, buildIteration, error_1)
	messagelogger.Log(2000, programName, buildVersion, buildIteration, error_1, "Just some text", error_2)

	fmt.Printf("\n\n--- Test 7: Using Maps ---------------------------------------------------------\n\n")

	messageTemplates[1001] = "Using maps"

	var detailsMap = map[string]string{
		"Husband": "Bob",
		"Wife":    "Mary",
	}

	var detailsMap2 = map[string]string{
		"Son":      "Bobbie",
		"Daughter": "Jane",
	}

	messagelogger.Log(1000, detailsMap, detailsMap2)
	messagelogger.Log(1001, detailsMap, detailsMap2)

	// ------------------------------------------------------------------------
	// Test a custom logger
	// ------------------------------------------------------------------------

	fmt.Printf("\n\n--------------------------------------------------------------------------------")
	fmt.Printf("\n--- Custom logger --------------------------------------------------------------")
	fmt.Printf("\n--------------------------------------------------------------------------------")

	log.SetFlags(0)

	terseMessageLogger := &messagelogger.MessageLoggerImpl{
		Logger:        &logger.LoggerImpl{},
		MessageFormat: &messageformat.MessageFormatTerse{},
	}

	terseMessageLogger.SetLogLevel(messagelogger.LevelDebug)

	fmt.Printf("\n\n--- Test 11: Original logger ---------------------------------------------------\n\n")
	boilerplateLogging(terseMessageLogger)

	fmt.Printf("\n\n--- Test 12: Add customized id -------------------------------------------------\n\n")

	terseMessageLogger.MessageId = &messageid.MessageIdDefault{IdTemplate: "test-%04d"}
	boilerplateLogging(terseMessageLogger)

	fmt.Printf("\n\n--- Test 13: Add text ----------------------------------------------------------\n\n")

	terseMessageLogger.MessageText = &messagetext.MessageTextDefault{TextTemplates: messageTemplates}
	boilerplateLogging(terseMessageLogger)

	fmt.Printf("\n\n--- Test 14: Add log levels ----------------------------------------------------\n\n")

	terseMessageLogger.MessageLogLevel = &messageloglevel.MessageLogLevelSenzingApi{}
	boilerplateLogging(terseMessageLogger)

	fmt.Printf("\n\n--- Test 15: Add status --------------------------------------------------------\n\n")

	terseMessageLogger.MessageStatus = &messagestatus.MessageStatusSenzingApi{}
	boilerplateLogging(terseMessageLogger)

	fmt.Printf("\n\n--- Test 16: Add logging golang errors -----------------------------------------\n\n")

	terseMessageLogger.Log(1000, programName, buildVersion, buildIteration, error_1)
	terseMessageLogger.Log(2000, programName, buildVersion, buildIteration, error_1, "Just some text", error_2)

	// ------------------------------------------------------------------------
	// The following demonstrates the low-level logger calls for
	// Trace, Debug, Info, Warn, and Error.
	// ------------------------------------------------------------------------

	fmt.Printf("\n\n--------------------------------------------------------------------------------")
	fmt.Printf("\n--- Low-level logger tests -----------------------------------------------------")
	fmt.Printf("\n--------------------------------------------------------------------------------\n\n")

	log.Println("Test Trace")
	logger.Trace("trace prints")
	logger.Tracef("trace A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Debug")
	logger.Debug("debug prints")
	logger.Debugf("debug A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Info")
	logger.Info("info prints")
	logger.Infof("info A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Warn")
	logger.Warn("warn prints")
	logger.Warnf("warn A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Error")
	logger.Error("error prints")
	logger.Errorf("error A: %s B: %s C: %d", "aaa", "bbb", 35)

	// Avoid long running logging when appropriate.

	if logger.IsDebug() {
		logger.Debugf("%s", complexProcess())
	}

	// Note:  the first Fatal or Panic issued will exit the program.

	log.Println("Test Fatal")
	//	logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test Panic")
	//		logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	log.Println("Test varadic")
	_, err := time.LoadLocation("bob")
	logger.Info("Should be error: ", err)

	log.Println("End")

}
