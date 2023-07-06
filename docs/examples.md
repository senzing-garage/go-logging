# go-logging examples

The following examples can be seen in actual code at
[main.go](../main.go).

In each of the following examples, the following import is assumed:

```go
import "github.com/senzing/go-logging/logging"
```

## Basic

The basic use of senzing/go-logging looks like this:

```go
logger, _ := logging.New()
logger.Log(2001, "A message")
```

Output:

```json
{"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"2001","details":{"1":"A message"}}
```

## Log level

The log level is determined by the message ID number.  Visit
[Logging levels](../README.md#logging-levels)
to see the different ranges for log levels.
Example:

```go
logger, _ := logging.New()
logger.Log(0999, "TRACE level")
logger.Log(1000, "DEBUG level")
logger.Log(2000, "INFO  level")
logger.Log(3000, "WARN  level")
logger.Log(4000, "ERROR level")
logger.Log(5000, "FATAL level")
logger.Log(6000, "PANIC level")
logger.Log(7000, "undefined level")
logger.Log(8000, "undefined level")
```

Output:

```json
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"2000","details":{"1":"INFO  level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"WARN","id":"3000","details":{"1":"WARN  level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"ERROR","id":"4000","details":{"1":"ERROR level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"FATAL","id":"5000","details":{"1":"FATAL level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"PANIC","id":"6000","details":{"1":"PANIC level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"PANIC","id":"7000","details":{"1":"undefined level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"PANIC","id":"8000","details":{"1":"undefined level"}}
 ```

Notice that the TRACE and DEBUG messages were not logged.
That is because the log level was set to INFO.
To set the log level to include TRACE and DEBUG,
add a "logger option" to set the log level.
Example:

```go
loggerOptions := []interface{}{
    &logging.OptionLogLevel{Value: "TRACE"},
}
logger, _ := logging.New(loggerOptions...)
logger.Log(0999, "TRACE level")
logger.Log(1000, "DEBUG level")
logger.Log(2000, "INFO  level")
```

Output:

```json
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"TRACE","id":"999","details":{"1":"TRACE level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"DEBUG","id":"1000","details":{"1":"DEBUG level"}}
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"2000","details":{"1":"INFO  level"}}
```

## Customize the id field

To create a unique identifier, not just an integer,
a [go format string](https://pkg.go.dev/fmt)
can be used as an ID template.
**NOTE:** The identifier string should include `%04d` so that the message number is included in the message ID.
Example:

```go
loggerOptions := []interface{}{
    &logging.OptionMessageIdTemplate{Value: "my-message-%04d"},
}
logger, _ = logging.New(loggerOptions...)
logger.Log(2002, "A message")
```

Output:

```json
 {"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"my-message-2002","details":{"1":"A message"}}
```

## Log additional information

In addition to a message identification integer ("id"), other types can be logged.
Example:

```go
aMap := map[int]string{
    10: "ten",
    20: "twenty",
}

aStruct := struct {
    Name string
    ID   int
}{
    Name: "Robert Smith",
    ID:   123145,
}

logger.Log(2003, "Robert Smith", 12345, aMap, aStruct)
```

Output:

```json
{"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"my-message-2003","details":{"1":"Robert Smith","2":12345,"3":"map[int]string{10:\"ten\", 20:\"twenty\"}","4":"struct { Name string; ID int }{Name:\"Robert Smith\", ID:123145}"}}
 ```

Output pretty printed:

```json
{
    "time": "YYYY-MM-DDThh:mm:ss.nnnnnnZ",
    "level": "INFO",
    "id": "my-message-2003",
    "details": {
        "1": "Robert Smith",
        "2": 12345,
        "3": "map[int]string{10:\"ten\", 20:\"twenty\"}",
        "4": "struct { Name string; ID int }{Name:\"Robert Smith\", ID:123145}"
    }
}
```

The fields submitted in the Log() call are seen in the "details" of the log message.
They will be listed in the order specified in the Log() call.

## Adding a text field

The additional information that is submitted in a Log() call can be used to create a text message.
By mapping message numbers to format strings, the Log() call will create formatted text output.
Example:

```go
idMessages := map[int]string{
     999:  "A test of TRACE.",
    1000: "A test of DEBUG.",
    2000: "A test of INFO.",
    2004: "The favorite number for %s is %d.",
    3000: "A test of WARN.",
    4000: "A test of ERROR.",
    5000: "A test of FATAL.",
    6000: "A test of PANIC.",
 }

loggerOptions := []interface{}{
    &logging.OptionIdMessages{Value: idMessages},
}
logger, _ = logging.New(loggerOptions...)
logger.Log(2004, "Robert Smith", 12345)
```

Output:

```json
{"time":"2023-07-06T19:38:32.350977388Z","level":"INFO","text":"The favorite number for Robert Smith is 12345.","id":"2004","details":{"1":"Robert Smith","2":12345}}
```

Notice that the information used to build the formatted text still remains in the "details" map.
This is by design.

## Logging errors

Go errors can also be logged.
Example:

```go
 err1 := errors.New("error #1")
 err2 := errors.New("error #2")
 logger.Log(2005, err1, err2)
```

Output:

```json
{"time":"YYYY-MM-DDThh:mm:ss.nnnnnnZ","level":"INFO","id":"2005","errors":["error #1","error #2"]}
```

## Guards

go-logging implements "guards",
e.g. IsXxxxx() methods,
to avoid calling a Log() method that
wouldn't print anyway because of the logging level.
For instance, there's no reason to call a DEBUG Log() method when the
logging level is set to INFO.  Guards prevent this.
Example:

```go
if logger.IsDebug() {
    logger.Log(1, "Log only in DEBUG mode", complexProcess())
}
```
