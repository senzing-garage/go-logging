/*
# Synopsis

The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (https://pkg.go.dev/log).

# Overview

The Senzing go-logging packages use the message number to coordinate aspects of the log message such as
message identification, message text, status, and logging level.

go-logging also allows different formatting options such as JSON or simply terse messages.

go-logging extends the levels of logging to include:
Trace, Debug, Info, Warn, Error, Fatal, and Panic.

go-logging implements "guards",
e.g. IsXxxxx() methods,
to avoid calling a Log() method that
wouldn't print anyway because of the logging level.
For instance, there's no reason to call a DEBUG Log() method when the
logging level is set to INFO.  Guards prevent this.
Example:

	if messageLogger.IsDebug() {
		messageLogger.Log(1, "Log only in DEBUG mode", complexProcess())
	}

The basic use of senzing/go-logging looks like this:

	import "log"
	import "github.com/senzing/go-logging/messagelogger"

	log.SetFlags(0)
	messageLogger, _ := messagelogger.New()
	messageLogger.Log(1)

Output:

	INFO 1:

# Examples

The following examples can be seen in actual code at
https://github.com/Senzing/go-logging/blob/main/main.go

In each of the following examples, the following imports are assumed:

	import "log"
	import "github.com/senzing/go-logging/messagelogger"

-- Configure log --------------------------------------------------------------

The Go "log" package can be independently configured.
Example:

	log.SetFlags(log.LstdFlags)
	messageLogger.Log(2)

Output:

	YYYY/MM/DD HH:MM:SS INFO 2:

-- Customize the id field -----------------------------------------------------

To create a unique identifier, not just an integer,
a go format string
(https://pkg.go.dev/fmt)
can be used as an ID template.
Example:

	log.SetFlags(0)
	messageId := &messageid.MessageIdTemplated{
		IdTemplate: "senzing-9999%04d",
	}
	messageLogger, _ = messagelogger.New(messageId)
	messagelogger.Log(3)

Output:

	INFO senzing-99990003:

-- Log additional information -------------------------------------------------

In addition to a message identification integer, other types can be logged.
Example:

	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}
	messageLogger.Log(4, "Robert Smith", 12345, aMap)

Output:

	INFO senzing-99990004: [map[1:Robert Smith 2:12345 3:map[int]string{10:"ten", 20:"twenty"}]]

The fields submitted in the Log() call are seen in a map in the log message.
They will be listed in the order specified in the Log() call.

-- Adding a text field --------------------------------------------------------

The additional information that is submitted in a Log() call can be used to create a text message.
By mapping message numbers to format strings, the Log() call will create formatted text output.
Example:

	messageText := &messagetext.MessageTextTemplated{
		IdMessage: map[int]string{
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

Output:

	INFO senzing-99990005: The favorite number for Robert Smith is 12345. [map[1:Robert Smith 2:12345 3:map[int]string{10:"ten", 20:"twenty"}]]

Notice that the information used to build the formatted text still remains in the map.
This is by design.

-- Log level ------------------------------------------------------------------

A log level can be specified anywhere after the first parameter (the message number parameter).
Example:

	import "github.com/senzing/go-logging/logger"

	messageLogger.Log(6, "Robert Smith", 12345, aMap, logger.LevelError)

Output:

	ERROR senzing-99990006: Person number #12345 is Robert Smith. [map[1:Robert Smith 2:12345 3:map[int]string{10:"ten", 20:"twenty"} 4:4]]

The logging level can be automated by identifying a MessageLogLevel of type MessageLogLevelInterface.
Example:

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

Output:

	INFO senzing-99990999: A test of INFO.
	WARN senzing-99991000: A test of WARN.
	ERROR senzing-99992000: A test of ERROR.

-- Status ---------------------------------------------------------------------

A status field can be added to the log message by
identifying a MessageStatus of type MessageStatusInterface.
One method is to identify ranges of messages ids and their
corresponding statuses.
Example:

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

Output:

	INFO senzing-99990999: (INFO) A test of INFO.
	WARN senzing-99991000: (WARN) A test of WARN.
	ERROR senzing-99992000: (ERROR) A test of ERROR.

Status can also be individually assigned to message numbers.
Example:

	messageStatus2 := &messagestatus.MessageStatusById{
		IdStatus: map[int]string{
			999:  "Foo",
			1000: "Bar",
			2000: "Baz",
		},
	}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageId, messageText, messageStatus2)
	messageLogger.Log(999)
	messageLogger.Log(1000)
	messageLogger.Log(2000)

Output:

	INFO senzing-99990999: (Foo) A test of INFO.
	WARN senzing-99991000: (Bar) A test of WARN.
	ERROR senzing-99992000: (Baz) A test of ERROR.

-- Logging errors -------------------------------------------------------------

Go errors can also be logged.
Example:

	err1 := errors.New("error #1")
	err2 := errors.New("error #2")
	messageLogger.Log(10, err1, err2)


	Output:

	INFO senzing-99990010: Example errors. [map[1:error #1 2:error #2]]

-- Formatting -----------------------------------------------------------------

The format of the log message can be modified by choosing a different message format.
Example:

	messageFormat := &messageformat.MessageFormatJson{}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageFormat, messageId, messageText, messageStatus)
	messageLogger.Log(1000)
	messageLogger.Log(11, "Robert Smith", 12345, aMap, err1, err2)

Output:

	WARN {"id":"senzing-99991000","status":"WARN","text":"A test of WARN."}
	INFO {"id":"senzing-99990011","status":"INFO","text":"Robert Smith has a score of 12345.","errors":[{"text":"error #1"},{"text":"error #2"}],"details":{"1":"Robert Smith","2":12345,"3":"map[int]string{10:\"ten\", 20:\"twenty\"}"}}
*/
package main
