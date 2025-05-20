/*
The [Logging] interface lists the methods available to a Logging object.

For examples of use, see https://github.com/senzing-garage/go-logging/blob/main/logger/logger_test.go
*/
package logging

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/senzing-garage/go-helpers/wraperror"
	"github.com/senzing-garage/go-messaging/messenger"
	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types - interface
// ----------------------------------------------------------------------------

// The Logging interface has methods for creating different
// representations of a message.
type Logging interface {
	GetLogLevel() string                                      // Get the current level of logging.
	Is(logLevelName string) bool                              // Returns true if logLevelName message will be logged.
	IsDebug() bool                                            // Returns true if a DEBUG message will be logged.
	IsError() bool                                            // Returns true if an ERROR message will be logged.
	IsFatal() bool                                            // Returns true if a FATAL message will be logged.
	IsInfo() bool                                             // Returns true if an INFO message will be logged.
	IsPanic() bool                                            // Returns true if a PANIC message will be logged.
	IsTrace() bool                                            // Returns true if a TRACE message will be logged.
	IsWarn() bool                                             // Returns true if a WARN message will be logged.
	JSON(messageNumber int, details ...interface{}) string    // Return a JSON string with the message.
	Log(messageNumber int, details ...interface{})            // Log the message.
	NewError(messageNumber int, details ...interface{}) error // Return an error object with the message.
	SetLogLevel(logLevelName string) error                    // Set the level of logging.
}

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

type ExtractedValues struct {
	callerSkip          int
	componentIdentifier int
	idMessages          map[int]string
	idStatuses          map[int]string
	logLevel            string
	messageIDTemplate   string
	messageFields       []string
	output              io.Writer
	messengerOptions    []interface{}
}

// --- Override values when creating messages ---------------------------------

type MessageCode struct {
	Value string
}

type MessageDetails struct {
	Value interface{}
}

type MessageDuration struct {
	Value int64
}

type MessageID struct {
	Value string
}

type MessageLevel struct {
	Value string
}

type MessageLocation struct {
	Value string
}

type MessageReason struct {
	Value string
}

type MessageStatus struct {
	Value string
}

type MessageText struct {
	Value string
}

type MessageTime struct {
	Value time.Time
}

// --- Options for New() ------------------------------------------------------

type OptionCallerSkip struct {
	Value int
}

type OptionComponentID struct {
	Value int
}
type OptionIDMessages struct {
	Value map[int]string
}

type OptionIDStatuses struct {
	Value map[int]string
}

type OptionLogLevel struct {
	Value string
}

type OptionMessageField struct {
	Value string
}

type OptionMessageFields struct {
	Value []string
}

type OptionMessageIDTemplate struct {
	Value string
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

const (
	componentIdentifier = 9999
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Message ID Low-bound for message levels
// i.e. a message in range 0 - 999 is a TRACE message.
var IDLevelRangesAsString = map[int]string{ //nolint
	0:    LevelTraceName,
	1000: LevelDebugName,
	2000: LevelInfoName,
	3000: LevelWarnName,
	4000: LevelErrorName,
	5000: LevelFatalName,
	6000: LevelPanicName,
}

// Map from slog.Level to string representation.
var LevelToTextMap = map[slog.Level]string{ //nolint
	LevelDebugSlog: LevelDebugName,
	LevelErrorSlog: LevelErrorName,
	LevelFatalSlog: LevelFatalName,
	LevelInfoSlog:  LevelInfoName,
	LevelPanicSlog: LevelPanicName,
	LevelTraceSlog: LevelTraceName,
	LevelWarnSlog:  LevelWarnName,
}

// Map from string representation to Log level as typed integer.
var TextToLevelMap = map[string]slog.Level{ //nolint
	LevelDebugName: LevelDebugSlog,
	LevelErrorName: LevelErrorSlog,
	LevelFatalName: LevelFatalSlog,
	LevelInfoName:  LevelInfoSlog,
	LevelPanicName: LevelPanicSlog,
	LevelTraceName: LevelTraceSlog,
	LevelWarnName:  LevelWarnSlog,
}

// Order is important in AllMessageFields. Should match order in go-messaging/messenger/main.go's MessageFormat.
var AllMessageFields = []string{ //nolint
	"time",
	"id",
	"text",
	"code",
	"reason",
	"status",
	"duration",
	"location",
	"errors",
	"details",
}

var errForPackage = errors.New("logging")

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
The IsValidLogLevelName function checks the logLevelName to verify it is one of
"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC".

Input
  - logLevelName: A name to be tested.

Output
  - boolean: True if name in "TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", or "PANIC".
*/
func IsValidLogLevelName(logLevelName string) bool {
	_, ok := TextToLevelMap[logLevelName]

	return ok
}

/*
The New function creates a new instance of loggingInterface.
Adding options can be used to modify subcomponents.

Input
  - options: A list of options (usually having type OptionXxxxx) used to configure the logger.

Output
  - A logger
  - error
*/
func New(options ...interface{}) (Logging, error) {
	var (
		err    error
		result Logging
	)

	extractedValues := &ExtractedValues{
		callerSkip:          0,
		componentIdentifier: componentIdentifier,
		idMessages:          map[int]string{},
		idStatuses:          map[int]string{},
		logLevel:            LevelInfoName,
		messageIDTemplate:   "%d",
		messengerOptions:    []interface{}{},
		output:              os.Stderr,
	}
	extractFromOptions(extractedValues, options)

	err = verifyOptions(extractedValues)
	if err != nil {
		return result, err
	}

	constructOptions(extractedValues)

	// Create messenger.

	messenger, err := messenger.New(extractedValues.messengerOptions...)
	if err != nil {
		return result, wraperror.Errorf(err, "New")
	}

	slogLevel, ok := TextToLevelMap[extractedValues.logLevel]
	if !ok {
		return result, wraperror.Errorf(errForPackage, "unknown error level: %s", extractedValues.logLevel)
	}

	slogLeveler := new(slog.LevelVar)
	slogLeveler.Set(slogLevel)

	// Create logger.

	logger := slog.New(slog.NewJSONHandler(extractedValues.output, SlogHandlerOptions(slogLeveler, options...)))

	// Create LoggingInterface.

	loggingImpl := &BasicLogging{
		logger:    logger,
		messenger: messenger,
		leveler:   slogLeveler,
	}

	loggingImpl.initialize()

	return loggingImpl, nil
}

/*
The NewSenzingLogger function creates a new instance of Logging
specifically for use with senzing-tools.

Input
  - componentId: See list at https://github.com/senzing-garage/knowledge-base/blob/main/lists/senzing-product-ids.md
  - idMessage: A map of integer to string message templates.
  - options: Variadic arguments listing the options (usually having type OptionXxxxx) used to configure the logger.

Output
  - A logger
  - error
*/
func NewSenzingLogger(componentID int, idMessages map[int]string, options ...interface{}) (Logging, error) {
	optionMessageID := fmt.Sprintf("SZTL%04d", componentID) + "%04d"
	loggerOptions := []interface{}{
		OptionComponentID{Value: componentID},
		OptionIDMessages{Value: idMessages},
		OptionMessageFields{Value: []string{"id", "reason"}},
		OptionMessageIDTemplate{Value: optionMessageID},
	}
	loggerOptions = append(loggerOptions, options...)

	return New(loggerOptions...)
}

/*
The SlogHandlerOptions function returns a slog handler that includes TRACE, FATAL, and PANIC.
See: https://go.googlesource.com/exp/+/refs/heads/master/slog/example_custom_levels_test.go
*/
func SlogHandlerOptions(leveler slog.Leveler, options ...interface{}) *slog.HandlerOptions {
	if leveler == nil {
		leveler = LevelInfoSlog
	}

	// Default values.

	timeHidden := false

	// Process options.

	for _, value := range options {
		switch typedValue := value.(type) {
		case OptionTimeHidden:
			timeHidden = typedValue.Value
		default:
		}
	}

	handlerOptions := &slog.HandlerOptions{
		Level: leveler,
		ReplaceAttr: func(groups []string, slogAttr slog.Attr) slog.Attr {
			_ = groups
			if slogAttr.Key == slog.LevelKey {
				level := ""
				switch typedValue := slogAttr.Value.Any().(type) {
				case string:
					level = typedValue
				case slog.Level:
					level = typedValue.String()
				}
				switch level {
				case "DEBUG-4":
					slogAttr.Value = slog.StringValue(LevelTraceName)
				case "ERROR+4":
					slogAttr.Value = slog.StringValue(LevelFatalName)
				case "ERROR+8":
					slogAttr.Value = slog.StringValue(LevelPanicName)
				}
			}
			if slogAttr.Key == slog.MessageKey {
				slogAttr.Key = "text"
				aValue, isOK := slogAttr.Value.Any().(string)
				if isOK && (aValue == "") {
					return slog.Attr{}
				}
			}
			if slogAttr.Key == slog.TimeKey {
				if timeHidden {
					return slog.Attr{}
				}
				slogAttr.Value = slog.TimeValue(slogAttr.Value.Time().UTC())
			}

			return slogAttr
		},
	}

	return handlerOptions
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func constructOptions(extractedValues *ExtractedValues) {
	extractedValues.messengerOptions = append(
		extractedValues.messengerOptions,
		messenger.OptionIDMessages{Value: extractedValues.idMessages},
	)
	extractedValues.messengerOptions = append(
		extractedValues.messengerOptions,
		messenger.OptionIDStatuses{Value: extractedValues.idStatuses},
	)
	extractedValues.messengerOptions = append(
		extractedValues.messengerOptions,
		messenger.OptionMessageIDTemplate{Value: extractedValues.messageIDTemplate},
	)

	if extractedValues.callerSkip != 0 {
		extractedValues.messengerOptions = append(
			extractedValues.messengerOptions,
			messenger.OptionCallerSkip{Value: extractedValues.callerSkip},
		)
	}

	if extractedValues.messageFields != nil {
		extractedValues.messengerOptions = append(
			extractedValues.messengerOptions,
			messenger.OptionMessageFields{Value: extractedValues.messageFields},
		)
	}
}

func extractFromOptions(extracted *ExtractedValues, options []interface{}) {
	for _, value := range options {
		switch typedValue := value.(type) {
		case OptionCallerSkip:
			extracted.callerSkip = typedValue.Value
		case OptionComponentID:
			extracted.componentIdentifier = typedValue.Value
			extracted.messageIDTemplate = fmt.Sprintf("senzing-%04d", extracted.componentIdentifier) + "%04d"
		case OptionIDMessages:
			extracted.idMessages = typedValue.Value
		case OptionIDStatuses:
			extracted.idStatuses = typedValue.Value
		case OptionLogLevel:
			extracted.logLevel = typedValue.Value
		case OptionMessageField:
			extracted.messengerOptions = append(extracted.messengerOptions, messenger.OptionMessageField{Value: typedValue.Value})
		case OptionMessageFields:
			extracted.messageFields = typedValue.Value
		case OptionMessageIDTemplate:
			extracted.messageIDTemplate = typedValue.Value
		case OptionOutput:
			extracted.output = typedValue.Value
		}
	}
}

func verifyOptions(extractedValues *ExtractedValues) error {
	if extractedValues.componentIdentifier <= 0 || extractedValues.componentIdentifier > 9999 {
		return wraperror.Errorf(
			errForPackage,
			"componentIdentifier %d must be in range 1..9999. See https://github.com/senzing-garage/knowledge-base/blob/main/lists/senzing-product-ids.md",
			extractedValues.componentIdentifier,
		)
	}

	return nil
}
