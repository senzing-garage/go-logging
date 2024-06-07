# go-logging

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## Synopsis

The `go-logging` packages build a logging system
upon Go's experimental `slog` package (<https://pkg.go.dev/golang.org/x/exp/slog>).

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/go-logging.svg)](https://pkg.go.dev/github.com/senzing-garage/go-logging)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/go-logging)](https://goreportcard.com/report/github.com/senzing-garage/go-logging)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/go-logging/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/go-logging/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/go-logging/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/go-logging/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/go-logging/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/go-logging/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/go-logging/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/go-logging/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/go-logging/actions/workflows/go-test-windows.yaml)

## Overview

The Senzing `go-logging` packages use the message number to coordinate aspects of the log message such as
message identification, message text, and logging level.

### Logging levels

`go-logging` extends the levels of logging to include:
TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.

The message number determines the log record level.
The ranges are:

| Level     | Range     | Use                                                   | Comments                    |
|-----------|-----------|-------------------------------------------------------|-----------------------------|
| **TRACE** | 0000-0999 | Entry/Exit tracing                                    | May contain sensitive data. |
| **DEBUG** | 1000-1999 | Values seen during processing                         | May contain sensitive data. |
| **INFO**  | 2000-2999 | Process steps achieved                                |                             |
| **WARN**  | 3000-3999 | Unexpected situations, but processing was successful  |                             |
| **ERROR** | 4000-4999 | Unexpected situations, processing was not successful  |                             |
| **FATAL** | 5000-5999 | The process needs to shutdown                         |                             |
| **PANIC** | 6000-6999 | The underlying system is at issue                     |                             |
|           | 8000-8999 | Reserved for observer messages                        |                             |

**WARNING:** `TRACE` and `DEBUG` are meant for problem determination and should not be used in normal processing
as the log records, by convention, may contain sensitive data.

### Guards

`go-logging` supports "guards",
e.g. IsXxxxx() methods,
to avoid calling a `Log()` method that
wouldn't print anyway because of the logging level.
For instance, there's no reason to call a DEBUG `Log()` method when the
logging level is set to INFO.  Guards prevent this.
Example:

```go
 if logger.IsDebug() {
  logger.Log(1001, complexProcess())
 }
```

## Use

The basic use of senzing/go-logging looks like this:

```go
 import "github.com/senzing-garage/go-logging/logging"

 logger, _ := logging.New()
 logger.Log(2001, "Hello world!")
```

Output:

```console
{"time":"YYYY-MM-DDThh:mm:ss.nnnnnnnnn-00:00","level":"INFO","id":"2001","details":{"1":"Hello World!"}}
```

### Message format

Although not all fields may be present for an individual message,
a complete message has these fields:

```json
{
    "time": "YYYY-MM-DDThh:mm:ss.nnnnnnnnn-00:00",
    "level": "INFO",
    "text": "Sent SQL in /var/tmp/tmpfile.sql to database sqlite3://na:xxxxx@/tmp/sqlite/G2C.db",
    "id": "senzing-65032002",
    "status":  "status_message",
    "duration":  "",
    "location": "In processDatabase() at senzingschema.go:129",
    "errors": ["unknown value in foo",  "bar has no value"],
    "details": {
        "1": "/var/tmp/tmpfile.sql",
        "2": "sqlite3://na:xxxxx@/tmp/sqlite/G2C.db"
    }
}
```

### Logging output

By default, logging goes to `STDERR`.
Since `go-logging` is built upon the
[log](https://pkg.go.dev/log)
package,
this can be modified using
[log.SetOutput()](https://pkg.go.dev/log#SetOutput).

Examples:

1. To have the output go a file, it would be something like this:

    ```go
    import (
        "log"
        "os"
        "io"
    )

    aFile, err := os.Open("/path/to/a/logfile")
    log.SetOutput(io.Writer(aFile))
    ```

1. To have the output go to STDERR and a file, it would be something like this:

    ```go
    import (
        "log"
        "os"
        "io"
    )

    aFile, err := os.Open("/path/to/a/logfile")
    log.SetOutput(io.MultiWriter(os.Stderr, aFile))
    ```

### Use with senzing-tools

In the suite of
[senzing-tools](https://github.com/senzing-garage/senzing-tools),
logging is created by:

```go
import (
    "fmt"
    "github.com/senzing-garage/go-logging/logging"
)

var (
    ComponentId = 9999            // See https://github.com/senzing-garage/knowledge-base/blob/main/lists/senzing-component-ids.md
    IdMessages  = map[int]string{ // Message templates. Example: https://github.com/senzing-garage/init-database/blob/main/senzingconfig/main.go
        2000: "Today's greeting:  %s",
        4000: "Here's what happened: %s",
    }
    callerSkip = 3                // Used to determine "location" information. See https://pkg.go.dev/runtime#Caller
)

// Logging options. See https://github.com/senzing-garage/go-logging/blob/main/logging/main.go
loggerOptions := []interface{}{
    &logging.OptionCallerSkip{Value: callerSkip},
}

// Create a logger from a factory.
logger, err := logging.NewSenzingLogger(ComponentId, IdMessages, loggerOptions...)
if err != nil {
    fmt.Println(err)
}

// Write log record.
logger.Log(2000, "Hello, world!")

// Create an error
err = logger.NewError(4000, "A bad thing")
fmt.Printf("The error: %v\n", err)
```

Example output:

```console
{"time":"YYYY-MM-DDThh:mm:ss.nnZ","level":"INFO","text":"Today's greeting:  Hello, world!","id":"senzing-99992000","location":"In main() at main.go:137","details":{"1":"Hello, world!"}}
The error: {"time":"YYYY-MM-DDThh:mm:ss.nnZ","level":"ERROR","id":"senzing-99994000","text":"Here's what happened: A bad thing","location":"In main() at main.go:140","details":{"1":"A bad thing"}}
```

## References

- [API documentation](https://pkg.go.dev/github.com/senzing-garage/go-logging)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
