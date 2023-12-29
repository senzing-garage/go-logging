/*
The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (https://pkg.go.dev/log).

# Overview

The Senzing go-logging packages use the message number to coordinate aspects of the log message such as
message identification, message text, status, and logging level.

go-logging extends the levels of logging to include:
Trace, Debug, Info, Warn, Error, Fatal, and Panic.

go-logging implements "guards",
e.g. IsXxxxx() methods,
to avoid calling a Log() method that
wouldn't print anyway because of the logging level.
For instance, there's no reason to call a DEBUG Log() method when the
logging level is set to INFO.  Guards prevent this.
Example:

	if logger.IsDebug() {
		logger.Log(1, "Log only in DEBUG mode", complexProcess())
	}

The basic use of senzing/go-logging looks like this:

	import "github.com/senzing-garage/go-logging/logging"

	logger, _ := logging.New()
	logger.Log(2000, "A message")

Output:

	{"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"2000","details":{"1":"A message"}}

# Examples

For examples, visit:

- [Examples](https://github.com/senzing-garage/go-logging/blob/main/docs/examples.md)
- [main.go](https://github.com/senzing-garage/go-logging/blob/main/main.go)
*/
package main
