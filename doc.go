/*
# Synopsis

The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (https://pkg.go.dev/log).

# Overview

The basic use of senzing/go-logging looks like this:

	import "log"
	import "github.com/senzing/go-logging/messagelogger"

	log.SetFlags(0)
	messageLogger, _ := messagelogger.New()
	messageLogger.Log(1)

Output:

	INFO 1:

# Examples

In each of the following examples, remember to include the following imports:

	import "log"
	import "github.com/senzing/go-logging/messagelogger"

-- Configure log --------------------------------------------------------------

The "log" can be independently configured.
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

	INFO senzing-99990004: [map[1:"Robert Smith" 2:12345 3:map[int]string{10:"ten", 20:"twenty"}]]

The fields submitted in the *.Log()* call are seen in the "details"	field of the log message.
They will be listed in the order specified in the *.Log()* call.

-- Adding a text field --------------------------------------------------------

The additional information that is submitted in a *.Log()* call can be used to create a text message.
A map that maps error numbers to format-templates needs to be created and identified to the logger.
Example:

	messageText := &messagetext.MessageTextTemplated{
		TextTemplates: map[int]string{
			5:    "The favorite number for %s is %d",
			999:  "A test of INFO",
			1000: "A test of WARN",
			2000: "A test of ERROR",
		},
	}

	messageLogger, _ = messagelogger.New(messageId, messageText)
	messageLogger.Log(5, "Robert Smith", 12345, aMap)

Output:

	INFO senzing-99990005: The favorite number for Robert Smith is 12345 [map[1:"Robert Smith" 2:12345 3:map[int]string{10:"ten", 20:"twenty"}]]

Notice that the information used to build the "text" still remains in the "details".
This is by design.

-- Log level ------------------------------------------------------------------

A log level can be specified anywhere after the first parameter (the message number parameter).

	messageLogger.Log(6, "Robert Smith", 12345, aMap, logger.LevelError)

Output:

	ERROR senzing-99990006: The favorite number for Robert Smith is 12345 [map[1:"Robert Smith" 2:12345 3:map[int]string{10:"ten", 20:"twenty"} 4:4]

The logging level can be automated by identifying a MessageLogLevel of type MessageLogLevelInterface.

	messageLogLevel := &messageloglevel.MessageLogLevelById{
		LogRanges: map[int]logger.Level{
			1000: logger.LevelInfo,
			2000: logger.LevelWarn,
			3000: logger.LevelError,
			4000: logger.LevelDebug,
			5000: logger.LevelTrace,
			6000: logger.LevelFatal,
			7000: logger.LevelPanic,
		},
	}
	messageLogger, _ = messagelogger.New(messageLogLevel, messageId, messageText)

	messageLogger.Log(999)
	messageLogger.Log(1000)

Output:

	INFO senzing-99990999: A test of INFO
	WARN senzing-99991000: A test of WARN

-- Status ---------------------------------------------------------------------

A status field can be added to the log message by
by identifying a MessageStatus of type MessageStatusInterface.

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
	messageLogger.Log(1001)

Output:

	INFO {"id":"senzing-99990999","status":"INFO","text":"A test of INFO"}
	WARN {"id":"senzing-99991000","status":"WARN","text":"A test of WARN"}

-- Log errors------------------------------------------------------------------

Go errors can also be logged.
They will show up in the "errors" field.
Example:

	err1 := errors.New("Error #1")
	err2 := errors.New("Error #2")
	messagelogger.Log(2000, "Message", err1, err2)

Output:

	ERROR {"details":{"1":"Message"},"errors":[{"text":"Error #1"},{"text":"Error #2"}],"id":"senzing-99992000","status":"ERROR","text":"A test of ERROR"}
*/
package main
