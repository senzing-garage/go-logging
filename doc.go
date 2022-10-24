/*
# Synopsis

The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (https://pkg.go.dev/log).

# Overview

The basic use of senzing/go-logging looks like this:

	import "log"
	import "github.com/senzing/go-logging/messagelogger"

	log.SetFlags(0)
	messagelogger.Log(1)

Output:

	INFO {"id":"0001"}

# Details

In each of the following examples, remember to include the following imports:

	import "log"
	import "github.com/senzing/go-logging/messagelogger"

## Configure log

The "log" can be independently configured.
Example:

	log.SetFlags(log.LstdFlags)
	messagelogger.Log(2)

Output:

	YYYY/MM/DD HH:MM:SS INFO {"id":"0002"}

## Customize the "id" field

To create a unique identifier, not just an integer,
a go format string
(https://pkg.go.dev/fmt)
can be used as an ID template.
Example:

	log.SetFlags(0)
	messagelogger.GetMessageLogger().SetIdTemplate("senzing-9999%04d")
	messagelogger.Log(3)

Output:

	INFO {"id":"senzing-99990003"}

## Log additional information

In addition to a message identification integer, other types can be logged.
Example:

	aMap := map[int]string{
		10: "ten",
		20: "twenty",
	}
	messagelogger.Log(4, "Robert Smith", 12345, aMap)

Output:

	INFO {"id":"senzing-99990004","details":{"1":"Robert Smith","2":12345,"3":"(map[int]string)map[10:ten 20:twenty]"}}

The fields submitted in the *.Log()* call are seen in the "details"	field of the log message.
They will be listed in the order specified in the *.Log()* call.

## Using additional information in a text field

The additional information that is submitted in a *.Log()* call can be used to create a text message.
A map that maps error numbers to format-templates needs to be created and identified to the logger.
Example:

	var textTemplates = map[int]string{
		5:    "The favorite number for %s is %d",
		999:  "A test of INFO",
		1000: "A test of WARN",
		2000: "A test of ERROR",
	}
	messagelogger.GetMessageLogger().SetTextTemplates(textTemplates)
	messagelogger.Log(5, "Robert Smith", 12345, aMap)

Output:

	INFO {"id":"senzing-99990005","text":"The favorite number for Robert Smith is 12345","details":{"1":"Robert Smith","2":12345,"3":"(map[int]string)map[10:ten 20:twenty]"}}

Notice that the information used to build the "text" still remains in the "details".
This is by design.

## Log level

The logging level can be automated by identifying a MessageLogLevel of type MessageLogLevelInterface.

	messagelogger.GetMessageLogger().MessageLogLevel = &messageloglevel.MessageLogLevelSenzingApi{}
	messagelogger.Log(999)
	messagelogger.Log(1000)

Output:

	INFO {"id":"senzing-99990999","text":"A test of INFO"}
	WARN {"id":"senzing-99991000","text":"A test of WARN"}

## Status

A status field can be added to the log message by
by identifying a MessageStatus of type MessageStatusInterface.

	messagelogger.GetMessageLogger().MessageStatus = &messagestatus.MessageStatusById{}
	messagelogger.Log(999)
	messagelogger.Log(1000)

Output:

	INFO {"id":"senzing-99990999","status":"INFO","text":"A test of INFO"}
	WARN {"id":"senzing-99991000","status":"WARN","text":"A test of WARN"}

## Log errors

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
