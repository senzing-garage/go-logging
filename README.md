# go-logging

## Synopsis

The Senzing go-logging packages build a composable logging system
that sits on top of Go's log package (<https://pkg.go.dev/log>).

[![GoReportCard example](https://goreportcard.com/badge/github.com/senzing/go-logging)](https://goreportcard.com/report/github.com/senzing/go-logging)
[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-logging.svg)](https://pkg.go.dev/github.com/senzing/go-logging)
![example workflow](https://github.com/github/docs/actions/workflows/main.yml/badge.svg)

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

### Contents

1. [Preamble](#preamble)
    1. [Legend](#legend)
1. [Related artifacts](#related-artifacts)
1. [Expectations](#expectations)
1. [Errors](#errors)
1. [References](#references)

## Preamble

At [Senzing](http://senzing.com),
we strive to create GitHub documentation in a
"[don't make me think](https://github.com/Senzing/knowledge-base/blob/main/WHATIS/dont-make-me-think.md)" style.
For the most part, instructions are copy and paste.
Whenever thinking is needed, it's marked with a "thinking" icon :thinking:.
Whenever customization is needed, it's marked with a "pencil" icon :pencil2:.
If the instructions are not clear, please let us know by opening a new
[Documentation issue](https://github.com/Senzing/template-python/issues/new?template=documentation_request.md)
describing where we can improve.   Now on with the show...

### Legend

1. :thinking: - A "thinker" icon means that a little extra thinking may be required.
   Perhaps there are some choices to be made.
   Perhaps it's an optional step.
1. :pencil2: - A "pencil" icon means that the instructions may need modification before performing.
1. :warning: - A "warning" icon means that something tricky is happening, so pay attention.

## Related artifacts

1. [DockerHub](https://hub.docker.com/r/senzing/xxxxxxxx)
1. [Helm Chart](https://github.com/Senzing/charts/tree/main/charts/xxxxxxxx)

## Expectations

## Errors

1. See [docs/errors.md](docs/errors.md).

## References

- **Space:** This repository and demonstration require 6 GB free disk space.
- **Time:** Budget 40 minutes to get the demonstration up-and-running, depending on CPU and network speeds.
- **Background knowledge:** This repository assumes a working knowledge of:
  - [Docker](https://github.com/Senzing/knowledge-base/blob/main/WHATIS/docker.md)
