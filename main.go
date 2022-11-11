/*
go-logging/main.go implements examples.
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelevel"
	"github.com/senzing/go-logging/messagelocation"
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var globalLogger messagelogger.MessageLoggerInterface

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func complexProcess() string {
	time.Sleep(1000 * time.Second)
	return "slept"
}

func complexProcess2() string {
	defer logDuration(1, time.Now())
	time.Sleep(2 * time.Second)
	return "slept"
}

// func startTime() time.Time {
// 	return time.Now()
// }

func logDuration(id int, start time.Time) {
	globalLogger.Log(id, time.Since(start))

}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	fmt.Printf("\n\n-------------------------------------------------------------------------------")
	fmt.Printf("\n--- High-level logger tests ---------------------------------------------------")
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

	// ------------------------------------------------------------------------
	// The following demonstrates the high-level messagelogger calls.
	// ------------------------------------------------------------------------

	fmt.Printf("\n--- Test 1: - Overview --------------------------------------------------------\n\n")

	log.SetFlags(0)
	messageLogger, _ := messagelogger.New()
	messageLogger.Log(1)

	fmt.Printf("\n\n--- Test 2: - Configure log ---------------------------------------------------\n\n")

	log.SetFlags(log.LstdFlags)
	messageLogger.Log(2)

	fmt.Printf("\n\n--- Test 3: - Customize the id field ------------------------------------------\n\n")

	log.SetFlags(0)
	messageId := &messageid.MessageIdTemplated{
		MessageIdTemplate: "senzing-9999%04d",
	}
	messageLogger, _ = messagelogger.New(messageId)
	messageLogger.Log(3)

	fmt.Printf("\n\n--- Test 4: - Log additional information --------------------------------------\n\n")

	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}
	messageLogger.Log(4, "Robert Smith", 12345, aMap)

	fmt.Printf("\n\n--- Test 5: - Adding a text field ---------------------------------------------\n\n")

	messageText := &messagetext.MessageTextTemplated{
		IdMessages: map[int]string{
			5:    "The favorite number for %s is %d.",
			6:    "Person number #%[2]d is %[1]s.",
			10:   "Example errors.",
			11:   "%s has a score of %d.",
			999:  "A test of INFO.",
			1000: "A test of WARN.",
			2000: "A test of ERROR.",
		},
	}
	messageLogger, _ = messagelogger.New(messageId, messageText)
	messageLogger.Log(5, "Robert Smith", 12345, aMap)

	fmt.Printf("\n\n--- Test 6: - Log level -------------------------------------------------------\n\n")

	messageLogger.Log(6, "Robert Smith", 12345, aMap, logger.LevelError)

	fmt.Printf("\n\n--- Test 7: - Log level -------------------------------------------------------\n\n")

	messageLogLevel := &messagelevel.MessageLevelByIdRange{
		IdRanges: map[int]logger.Level{
			0000: logger.LevelInfo,
			1000: logger.LevelWarn,
			2000: logger.LevelError,
			3000: logger.LevelDebug,
			4000: logger.LevelTrace,
			5000: logger.LevelFatal,
			6000: logger.LevelPanic,
		},
	}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageId, messageText)
	messageLogger.Log(999)
	messageLogger.Log(1000)
	messageLogger.Log(2000)

	fmt.Printf("\n\n--- Test 8: - Status ----------------------------------------------------------\n\n")

	messageStatus := &messagestatus.MessageStatusByIdRange{
		IdRanges: map[int]string{
			0000: logger.LevelInfoName,
			1000: logger.LevelWarnName,
			2000: logger.LevelErrorName,
			3000: logger.LevelDebugName,
			4000: logger.LevelTraceName,
			5000: logger.LevelFatalName,
			6000: logger.LevelPanicName,
		},
	}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageId, messageText, messageStatus)
	messageLogger.Log(999)
	messageLogger.Log(1000)
	messageLogger.Log(2000)

	fmt.Printf("\n\n--- Test 9: - Status ----------------------------------------------------------\n\n")

	messageStatus2 := &messagestatus.MessageStatusById{
		IdStatuses: map[int]string{
			999:  "Foo",
			1000: "Bar",
			2000: "Baz",
		},
	}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageId, messageText, messageStatus2)
	messageLogger.Log(999)
	messageLogger.Log(1000)
	messageLogger.Log(2000)

	fmt.Printf("\n\n--- Test 10: - Logging errors -------------------------------------------------\n\n")

	err1 := errors.New("error #1")
	err2 := errors.New("error #2")
	messageLogger.Log(10, err1, err2)

	fmt.Printf("\n\n--- Test 11: - Formatting -----------------------------------------------------\n\n")

	messageFormat := &messageformat.MessageFormatJson{}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageFormat, messageId, messageText, messageStatus)
	messageLogger.Log(1000)
	messageLogger.Log(11, "Robert Smith", 12345, aMap, err1, err2)

	// ------------------------------------------------------------------------
	// The following demonstrates the system-wide logger calls.
	// ------------------------------------------------------------------------

	fmt.Printf("\n\n--- Test 12: - System loggers -------------------------------------------------\n\n")

	messageLogger1, _ := messagelogger.New(logger.LevelInfo)
	messageLogger2, _ := messagelogger.New(logger.LevelWarn)
	messageLogger3, _ := messagelogger.New(logger.LevelError)

	fmt.Println("------ Before")
	messageLogger1.Log(2001)
	messageLogger2.Log(2002)
	messageLogger3.Log(2003)

	fmt.Println("------ After")

	messagelogger.SetLogLevel(messagelogger.LevelInfo)

	messageLogger1.Log(2001)
	messageLogger2.Log(2002)
	messageLogger3.Log(2003)

	fmt.Println("------ New loggers")

	messageLogger4, _ := messagelogger.New(logger.LevelError)
	messageLogger4.Log(2004)

	fmt.Println("------ Inspection")

	logLevel, _ := messagelogger.GetLogLevel()
	logLevelName, _ := messagelogger.GetLogLevelAsString()

	fmt.Printf("System log level: %d\n", logLevel)
	fmt.Printf("System log level name: %s\n", logLevelName)

	// ------------------------------------------------------------------------
	// The following demonstrates the low-level logger calls for
	// Trace, Debug, Info, Warn, and Error.
	// ------------------------------------------------------------------------

	fmt.Printf("\n\n-------------------------------------------------------------------------------")
	fmt.Printf("\n--- Low-level logger tests ----------------------------------------------------")
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

	fmt.Printf("\n--- Test Trace - should be empty ----------------------------------------------\n")
	logger.Trace("trace prints")
	logger.Tracef("trace A: %s B: %s C: %d", "aaa", "bbb", 35)

	fmt.Printf("\n--- Test Debug - should be empty ----------------------------------------------\n")
	logger.Debug("debug prints")
	logger.Debugf("debug A: %s B: %s C: %d", "aaa", "bbb", 35)

	fmt.Printf("\n--- Test Info -----------------------------------------------------------------\n")
	logger.Info("info prints")
	logger.Infof("info A: %s B: %s C: %d", "aaa", "bbb", 35)

	fmt.Printf("\n--- Test Warn -----------------------------------------------------------------\n")
	logger.Warn("warn prints")
	logger.Warnf("warn A: %s B: %s C: %d", "aaa", "bbb", 35)

	fmt.Printf("\n--- Test Error ----------------------------------------------------------------\n")
	logger.Error("error prints")
	logger.Errorf("error A: %s B: %s C: %d", "aaa", "bbb", 35)

	// Note: The first Fatal or Panic issued will exit the program.

	fmt.Printf("\n--- Test Fatal - should be empty ----------------------------------------------\n")
	//	logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	fmt.Printf("\n--- Test Panic - should be empty ----------------------------------------------\n")
	//	logger.Fatal("fatal prints")
	//	logger.Fatalf("fatal A: %s B: %s C: %d", "aaa", "bbb", 35)

	// Avoid long running logging when appropriate.

	fmt.Printf("\n--- Test IsDebug - should be empty --------------------------------------------\n")
	if logger.IsDebug() {
		logger.Debugf("%s", complexProcess())
	}

	fmt.Printf("\n--- Test Varadic --------------------------------------------------------------\n")
	_, err := time.LoadLocation("bob")
	logger.Info("Should be error: ", err)

	// ------------------------------------------------------------------------
	// Senzing loggong.
	// ------------------------------------------------------------------------

	fmt.Printf("\n\n-------------------------------------------------------------------------------")
	fmt.Printf("\n--- Senzing logger tests ------------------------------------------------------")
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

	messageids := map[int]string{
		1: "Example duration",
	}
	globalLogger, _ = messagelogger.NewSenzingLogger(9999, messageids, messagelocation.CallerSkip(4))
	complexProcess2()

	fmt.Printf("\n\n-------------------------------------------------------------------------------")
	fmt.Printf("\n--- End -----------------------------------------------------------------------")
	fmt.Printf("\n-------------------------------------------------------------------------------\n\n")

}
