# go-logging

## Synopsis

The `go-logging` packages build a logging system
upon Go's experimental `slog` package (<https://pkg.go.dev/golang.org/x/exp/slog>).

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-logging.svg)](https://pkg.go.dev/github.com/senzing/go-logging)
[![GoReportCard example](https://goreportcard.com/badge/github.com/senzing/go-logging)](https://goreportcard.com/report/github.com/senzing/go-logging)
[![go-test.yaml](https://github.com/Senzing/go-logging/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-logging/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/go-logging/blob/main/LICENSE)

## Overview

The Senzing `go-logging` packages use the message number to coordinate aspects of the log message such as
message identification, message text, and logging level.

### Logging levels

`go-logging` extends the levels of logging to include:
TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.

The message number determines the log record level.
The ranges are:

| Level     | Range     | Use                                                   |
|-----------|-----------|-------------------------------------------------------|
| **TRACE** | 0000-0999 | Entry/Exit tracing                                    |
| **DEBUG** | 1000-1999 | Values seen during processing                         |
| **INFO**  | 2000-2999 | Process steps achieved                                |
| **WARN**  | 3000-3999 | Unexpected situations, but processing was successful  |
| **ERROR** | 4000-4999 | Unexpected situations, processing was not successful  |
| **FATAL** | 5000-5999 | The process needs to shutdown                         |
| **PANIC** | 6000-6999 | The underlying system is at issue                     |
|           | 8000-8999 | Reserved for observer messages                        |

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
 import "github.com/senzing/go-logging/logging"

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

## References

- [API documentation](https://pkg.go.dev/github.com/senzing/go-logging)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
