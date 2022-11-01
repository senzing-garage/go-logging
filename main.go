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
	"github.com/senzing/go-logging/messagelogger"
	"github.com/senzing/go-logging/messageloglevel"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func complexProcess() string {
	time.Sleep(1000 * time.Second)
	return "slept"
}

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {

	// ------------------------------------------------------------------------
	// The following demonstrates the high-level messagelogger calls.
	// ------------------------------------------------------------------------

	fmt.Printf("\n--- Test 1: - Overview -----------------------------------------------------------\n\n")

	log.SetFlags(0)
	messageLogger, _ := messagelogger.New()
	messageLogger.Log(1)

	fmt.Printf("\n\n--- Test 2: - Configure log ----------------------------------------------------\n\n")

	log.SetFlags(log.LstdFlags)
	messageLogger.Log(2)

	fmt.Printf("\n\n--- Test 3: - Customize the id field -------------------------------------------\n\n")

	log.SetFlags(0)
	messageId := &messageid.MessageIdTemplated{
		MessageIdTemplate: "senzing-9999%04d",
	}
	messageLogger, _ = messagelogger.New(messageId)
	messageLogger.Log(3)

	fmt.Printf("\n\n--- Test 4: - Log additional information ---------------------------------------\n\n")

	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}
	messageLogger.Log(4, "Robert Smith", 12345, aMap)

	fmt.Printf("\n\n--- Test 5: - Adding a text field ----------------------------------------------\n\n")

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

	fmt.Printf("\n\n--- Test 6: - Log level --------------------------------------------------------\n\n")

	messageLogger.Log(6, "Robert Smith", 12345, aMap, logger.LevelError)

	fmt.Printf("\n\n--- Test 7: - Log level --------------------------------------------------------\n\n")

	messageLogLevel := &messageloglevel.MessageLogLevelByIdRange{
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

	fmt.Printf("\n\n--- Test 8: - Status -----------------------------------------------------------\n\n")

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

	fmt.Printf("\n\n--- Test 9: - Status -----------------------------------------------------------\n\n")

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

	fmt.Printf("\n\n--- Test 10: - Logging errors --------------------------------------------------\n\n")

	err1 := errors.New("error #1")
	err2 := errors.New("error #2")
	messageLogger.Log(10, err1, err2)

	fmt.Printf("\n\n--- Test 11: - Formatting ------------------------------------------------------\n\n")

	messageFormat := &messageformat.MessageFormatJson{}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageFormat, messageId, messageText, messageStatus)
	messageLogger.Log(1000)
	messageLogger.Log(11, "Robert Smith", 12345, aMap, err1, err2)

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
