package logging

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/senzing/go-messaging/messenger"
	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

// The loggingInterface interface has methods for creating different
// representations of a message.
type LoggingInterface interface {
	Log(messageNumber int, details ...interface{})
	GetLogLevel() string
	SetLogLevel(logLevelName string) error
}

// --- Override values when creating messages ---------------------------------

type MessageDetails struct {
	Value interface{}
}

type MessageDuration struct {
	Value int64
}

type MessageId struct {
	Value string
}

type MessageLevel struct {
	Value string
}

type MessageLocation struct {
	Value string
}

type MessageStatus struct {
	Value string
}

type MessageText struct {
	Value interface{}
}

type MessageTime struct {
	Value time.Time
}

// --- Options for New() ------------------------------------------------------

type OptionCallerSkip struct {
	Value int
}

type OptionIdMessages struct {
	Value map[int]string
}

type OptionIdStatuses struct {
	Value map[int]string
}

type OptionLogLevel struct {
	Value string
}

type OptionMessageIdTemplate struct {
	Value string
}

type OptionSenzingComponentId struct {
	Value int
}

type OptionTimeHidden struct {
	Value bool
}

type OptionOutput struct {
	Value io.Writer
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Log levels as integers.
// Compatible with golang.org/x/exp/slog.
const (
	LevelTraceInt int = -8
	LevelDebugInt int = -4
	LevelInfoInt  int = 0
	LevelWarnInt  int = 4
	LevelErrorInt int = 8
	LevelFatalInt int = 12
	LevelPanicInt int = 16
)

// Strings representing the supported logging levels.
const (
	LevelDebugName = "DEBUG"
	LevelErrorName = "ERROR"
	LevelFatalName = "FATAL"
	LevelInfoName  = "INFO"
	LevelPanicName = "PANIC"
	LevelTraceName = "TRACE"
	LevelWarnName  = "WARN"
)

// Existing and new log levels used with slog.Level.
const (
	LevelDebugSlog = slog.LevelDebug
	LevelErrorSlog = slog.LevelError
	LevelFatalSlog = slog.Level(LevelFatalInt)
	LevelInfoSlog  = slog.LevelInfo
	LevelPanicSlog = slog.Level(LevelPanicInt)
	LevelTraceSlog = slog.Level(LevelTraceInt)
	LevelWarnSlog  = slog.LevelWarn
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message ID Low-bound for message levels
// i.e. a message in range 0 - 999 is a TRACE message.
var IdLevelRangesAsString = map[int]string{
	0000: LevelTraceName,
	1000: LevelDebugName,
	2000: LevelInfoName,
	3000: LevelWarnName,
	4000: LevelErrorName,
	5000: LevelFatalName,
	6000: LevelPanicName,
}

// Map from slog.Level to string representation.
var LevelToTextMap = map[slog.Level]string{
	LevelDebugSlog: LevelDebugName,
	LevelErrorSlog: LevelErrorName,
	LevelFatalSlog: LevelFatalName,
	LevelInfoSlog:  LevelInfoName,
	LevelPanicSlog: LevelPanicName,
	LevelTraceSlog: LevelTraceName,
	LevelWarnSlog:  LevelWarnName,
}

// Map from string representation to Log level as typed integer.
var TextToLevelMap = map[string]slog.Level{
	LevelDebugName: LevelDebugSlog,
	LevelErrorName: LevelErrorSlog,
	LevelFatalName: LevelFatalSlog,
	LevelInfoName:  LevelInfoSlog,
	LevelPanicName: LevelPanicSlog,
	LevelTraceName: LevelTraceSlog,
	LevelWarnName:  LevelWarnSlog,
}

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The New function creates a new instance of loggingInterface.
Adding options can be used to modify subcomponents.

Input
  - options: A list of options (usually having type OptionXxxxx) used to configur the logger.

Output
  - A logger
  - error
*/
func New(options ...interface{}) (LoggingInterface, error) {
	var err error = nil
	var result LoggingInterface = nil

	// Default values.

	var (
		callerSkip          int            = 0
		idMessages          map[int]string = map[int]string{}
		idStatuses          map[int]string = map[int]string{}
		logLevel            string         = LevelInfoName
		messageIdTemplate   string         = "%d"
		componentIdentifier int            = 9999
		output              io.Writer      = os.Stderr
	)

	// Process options.

	for _, value := range options {
		switch typedValue := value.(type) {
		case *OptionCallerSkip:
			callerSkip = typedValue.Value
		case *OptionIdMessages:
			idMessages = typedValue.Value
		case *OptionIdStatuses:
			idStatuses = typedValue.Value
		case *OptionLogLevel:
			logLevel = typedValue.Value
		case *OptionMessageIdTemplate:
			messageIdTemplate = typedValue.Value
		case *OptionOutput:
			output = typedValue.Value
		case *OptionSenzingComponentId:
			componentIdentifier = typedValue.Value
			messageIdTemplate = fmt.Sprintf("senzing-%04d", componentIdentifier) + "%04d"
		}
	}

	// Detect incorrect option values.

	if componentIdentifier <= 0 || componentIdentifier >= 10000 {
		err := errors.New("componentIdentifier must be in range 1..9999. See https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-product-ids.md")
		return result, err
	}

	if idMessages == nil {
		err := errors.New("messages must be a map[int]string, not nil")
		return result, err
	}

	if idStatuses == nil {
		err := errors.New("statuses must be a map[int]string, not nil")
		return result, err
	}

	slogLevel, ok := TextToLevelMap[logLevel]
	if !ok {
		err := fmt.Errorf("unknown error level: %s", logLevel)
		return result, err
	}
	var slogLeveler = new(slog.LevelVar)
	slogLeveler.Set(slogLevel)

	// Create messenger.

	messengerOptions := []interface{}{
		&messenger.OptionIdMessages{Value: idMessages},
		&messenger.OptionIdStatuses{Value: idStatuses},
		&messenger.OptionMessageIdTemplate{Value: messageIdTemplate},
		&messenger.OptionCallerSkip{Value: callerSkip},
	}

	messenger, err := messenger.New(messengerOptions...)
	if err != nil {
		return result, err
	}

	// Create logger.

	logger := slog.New(SlogHandlerOptions(slogLeveler, options...).NewJSONHandler(output))

	// Create LoggingInterface.

	loggingImpl := &LoggingImpl{
		logger:    logger,
		messenger: messenger,
		leveler:   slogLeveler,
	}

	err = loggingImpl.initialize()
	if err != nil {
		return result, err
	}

	result = loggingImpl
	return result, err
}

/*
The NewSenzingSdkLogger function creates a new instance of loggingInterface
specifically for use with g2-sdk-go-* packages.

Input
  - componentId: See list at https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-product-ids.md
  - idMessage: A map of integer to string message templates.
  - options: Variadic arguments listing the options (usually having type OptionXxxxx) used to configure the logger.

Output
  - A logger
  - error
*/
func NewSenzingSdkLogger(componentId int, idMessages map[int]string, options ...interface{}) (LoggingInterface, error) {
	loggerOptions := []interface{}{
		&OptionIdMessages{Value: idMessages},
		&OptionSenzingComponentId{Value: componentId},
	}
	loggerOptions = append(loggerOptions, options...)
	return New(loggerOptions...)
}

/*
The NewSenzingToolsLogger function creates a new instance of loggingInterface
specifically for use with senzing-tools.

Input
  - componentId: See list at https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-product-ids.md
  - idMessage: A map of integer to string message templates.
  - options: Variadic arguments listing the options (usually having type OptionXxxxx) used to configure the logger.

Output
  - A logger
  - error
*/
func NewSenzingToolsLogger(componentId int, idMessages map[int]string, options ...interface{}) (LoggingInterface, error) {
	loggerOptions := []interface{}{
		&OptionIdMessages{Value: idMessages},
		&OptionSenzingComponentId{Value: componentId},
	}
	loggerOptions = append(loggerOptions, options...)
	return New(loggerOptions...)
}

/*
The SlogHandlerOptions function returns a slog handler that includes TRACE, FATAL, and PANIC.
TODO: Move to Senzing/go-logging.
See: https://go.googlesource.com/exp/+/refs/heads/master/slog/example_custom_levels_test.go
*/
func SlogHandlerOptions(leveler slog.Leveler, options ...interface{}) *slog.HandlerOptions {
	if leveler == nil {
		leveler = LevelInfoSlog
	}

	// Default values.

	var (
		timeHidden bool = false
	)

	// Process options.

	for _, value := range options {
		switch typedValue := value.(type) {
		case *OptionTimeHidden:
			timeHidden = typedValue.Value
		}
	}

	handlerOptions := &slog.HandlerOptions{
		Level: leveler,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := ""
				switch typedValue := a.Value.Any().(type) {
				case string:
					level = typedValue
				case slog.Level:
					level = typedValue.String()
				}
				switch level {
				case "DEBUG-4":
					a.Value = slog.StringValue(LevelTraceName)
				case "ERROR+4":
					a.Value = slog.StringValue(LevelFatalName)
				case "ERROR+8":
					a.Value = slog.StringValue(LevelPanicName)
				}
			}
			if a.Key == slog.MessageKey {
				a.Key = "text"
				if a.Value.Any().(string) == "" {
					return slog.Attr{}
				}
			}
			if a.Key == slog.TimeKey {
				if timeHidden {
					return slog.Attr{}
				}
			}
			return a
		},
	}
	return handlerOptions
}
