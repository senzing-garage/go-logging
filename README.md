# go-logging

## Synopsis

The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (<https://pkg.go.dev/log>).

[![GoReportCard example](https://goreportcard.com/badge/github.com/senzing/go-logging)](https://goreportcard.com/report/github.com/senzing/go-logging)
[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-logging.svg)](https://pkg.go.dev/github.com/senzing/go-logging)

## Overview

The Senzing go-logging packages use the message number to coordinate aspects of the log message such as
message identification, message text, status, and logging level.

go-logging also allows different formatting options such as JSON or simply terse messages.

go-logging extends the levels of logging to include:
Trace, Debug, Info, Warn, Error, Fatal, and Panic.

go-logging supports "guards",
e.g. IsXxxxx() methods,
to avoid calling a `Log()` method that
wouldn't print anyway because of the logging level.
For instance, there's no reason to call a DEBUG `Log()` method when the
logging level is set to INFO.  Guards prevent this.
Example:

```go
 if logger.IsDebug() {
  logger.Debugf("%s", complexProcess())
 }
```

The basic use of senzing/go-logging looks like this:

```go
 import "log"
 import "github.com/senzing/go-logging/messagelogger"

 log.SetFlags(0)
 messageLogger, _ := messagelogger.New()
 messageLogger.Log(1)
```

Output:

```console
 INFO 1:
```

The API documentation is available at
[pkg.go.dev/github.com/senzing/go-logging](https://pkg.go.dev/github.com/senzing/go-logging)
