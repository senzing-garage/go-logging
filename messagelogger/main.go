/*
The messagelogger package generates messages, logs messages, or creates errors from messages.

For examples of use, see https://github.com/Senzing/go-logging/blob/main/messagelogger/messagelogger_test.go
*/
package messagelogger

import (
	"errors"
	"fmt"
	"sync"

	"github.com/senzing/go-logging/logger"
	"github.com/senzing/go-logging/messagedate"
	"github.com/senzing/go-logging/messagedetails"
	"github.com/senzing/go-logging/messageduration"
	"github.com/senzing/go-logging/messageerrors"
	"github.com/senzing/go-logging/messageformat"
	"github.com/senzing/go-logging/messageid"
	"github.com/senzing/go-logging/messagelevel"
	"github.com/senzing/go-logging/messagelocation"
	"github.com/senzing/go-logging/messagestatus"
	"github.com/senzing/go-logging/messagetext"
	"github.com/senzing/go-logging/messagetime"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The Level type is used to identify the integer is the detail parameters.
// The Level type is used for logging levels. (e.g. LevelInfo, LevelWarn, etc.)
type Level int

/*
The MessageLoggerInterface type defines methods for creating messages, logging messages, and creating errors from messages.
It also has convenience methods for setting and getting the current log level.
*/
type MessageLoggerInterface interface {
	Error(messageNumber int, details ...interface{}) error             // Returns an error type populated with the message.
	GetLogLevel() Level                                                // Gets the logger instance logging level.
	GetLogLevelAsString() string                                       // Gets the logger instance logging level in string representation.
	IsDebug() bool                                                     // Returns true if a DEBUG message will be logged.
	IsError() bool                                                     // Returns true if an ERROR message will be logged.
	IsFatal() bool                                                     // Returns true if a FATAL message will be logged.
	IsInfo() bool                                                      // Returns true if an INFO message will be logged.
	IsPanic() bool                                                     // Returns true if a PANIC message will be logged.
	IsTrace() bool                                                     // Returns true if a TRACE message will be logged.
	IsWarn() bool                                                      // Returns true if a WARN message will be logged.
	Log(messageNumber int, details ...interface{}) error               // Logs the message.
	Message(messageNumber int, details ...interface{}) (string, error) // Returns the message.
	SetLogLevel(level Level) MessageLoggerInterface                    // Sets the logger instance logging level.
	SetLogLevelFromString(levelString string) MessageLoggerInterface   // Sets the logger instance logging level using a string representation.
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// LevelXxxx values are an enumeration of typed integers representing logging levels.
// Must match what's in logger/main.go
const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var (
	lock                   = &sync.Mutex{}
	isSystemLogLevelSet    = false
	systemLogLevel         = LevelInfo
	messageLoggerObservers = []MessageLoggerInterface{}
)

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

/*
The new() function creates a base logger using null and default components.
*/
func new(interfaces ...interface{}) (MessageLoggerInterface, error) {
	var err error = nil

	// Start with default values.

	logLevel := LevelInfo
	result := &MessageLoggerDefault{
		Logger:          &logger.LoggerDefault{},
		MessageDate:     &messagedate.MessageDateNull{},
		MessageDetails:  &messagedetails.MessageDetailsNull{},
		MessageDuration: &messageduration.MessageDurationNull{},
		MessageErrors:   &messageerrors.MessageErrorsNull{},
		MessageFormat:   &messageformat.MessageFormatDefault{},
		MessageId:       &messageid.MessageIdDefault{},
		MessageLocation: &messagelocation.MessageLocationNull{},
		MessageLevel: &messagelevel.MessageLevelDefault{
			DefaultLogLevel: logger.LevelInfo,
		},
		MessageStatus: &messagestatus.MessageStatusNull{},
		MessageText:   &messagetext.MessageTextNull{},
		MessageTime:   &messagetime.MessageTimeNull{},
	}

	// Incorporate parameters.

	var errorsList []interface{}
	if len(interfaces) > 0 {
		for _, value := range interfaces {
			switch typedValue := value.(type) {
			case logger.LoggerInterface:
				result.Logger = typedValue
			case messagedate.MessageDateInterface:
				result.MessageDate = typedValue
			case messagedetails.MessageDetailsInterface:
				result.MessageDetails = typedValue
			case messageduration.MessageDurationInterface:
				result.MessageDuration = typedValue
			case messageerrors.MessageErrorsInterface:
				result.MessageErrors = typedValue
			case messageformat.MessageFormatInterface:
				result.MessageFormat = typedValue
			case messageid.MessageIdInterface:
				result.MessageId = typedValue
			case messagelevel.MessageLevelInterface:
				result.MessageLevel = typedValue
			case messagelocation.MessageLocationInterface:
				result.MessageLocation = typedValue
			case messagestatus.MessageStatusInterface:
				result.MessageStatus = typedValue
			case messagetext.MessageTextInterface:
				result.MessageText = typedValue
			case messagetime.MessageTimeInterface:
				result.MessageTime = typedValue
			case logger.Level:
				logLevelCandidate, ok := value.(logger.Level)
				if ok {
					logLevel = Level(logLevelCandidate)
				}
			case Level:
				logLevelCandidate, ok := value.(Level)
				if ok {
					logLevel = Level(logLevelCandidate)
				}
			default:
				errorsList = append(errorsList, typedValue)
			}
		}
	}
	result.SetLogLevel(logLevel)

	// Report any unknown parameters.

	if len(errorsList) > 0 {
		err = fmt.Errorf("unsupported interfaces: %#v", errorsList)
	}

	// If system logging level set, set this logger to that level and
	// add this messageLogger to the Observers list.
	// Do this in a thread-safe way.

	lock.Lock()
	defer lock.Unlock()

	if isSystemLogLevelSet {
		result.SetLogLevel(systemLogLevel)
	}

	if messageLoggerObservers == nil {
		messageLoggerObservers = make([]MessageLoggerInterface, 0)
	}

	messageLoggerObservers = append(messageLoggerObservers, result)

	return result, err
}

/*
The New function creates a new instance of MessageLoggerDefault.
The default message logger uses default and null subcomponents.
To use non-default subcomponents,
adding parameters to New() can specify the subcomponent desired.
The parameters can be of the following type and in any order:

  - logger.Level
  - logger.LoggerInterface
  - messagedate.MessageDateInterface
  - messagedetails.MessageDetailsInterface
  - messageduration.MessageDurationInterface
  - messageerrors.MessageErrorsInterface
  - messageformat.MessageFormatInterface
  - messageid.MessageIdInterface
  - messagelevel.MessageLevelInterface
  - messagelocation.MessageLocationInterface
  - messagestatus.MessageStatusInterface
  - messagetext.MessageTextInterface
  - messagetime.MessageTimeInterface

If a type is specified multiple times,
the last instance instance of the type specified wins.
*/
func New(interfaces ...interface{}) (MessageLoggerInterface, error) {

	messageDetails := &messagedetails.MessageDetailsDefault{}
	messageErrors := &messageerrors.MessageErrorsDefault{}

	var newInterfaces = []interface{}{
		messageDetails,
		messageErrors,
	}

	// Add other user-supplied interfaces to newInterfaces.

	newInterfaces = append(newInterfaces, interfaces...)

	// Using a Factory Pattern, build the messagelogger.

	return new(newInterfaces...)
}

/*
The NewSenzingLogger function creates a new instance of MessageLoggerInterface
that is tailored to Senzing applications.
Like New(), adding parameters can be used to modify subcomponents.
*/
func NewSenzingLogger(componentIdentifier int, idMessages map[int]string, interfaces ...interface{}) (MessageLoggerInterface, error) {

	// Detect incorrect parameter values.

	if componentIdentifier <= 0 || componentIdentifier >= 10000 {
		err := errors.New("componentIdentifier must be in range 1..9999. See https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-product-ids.md")
		return nil, err
	}

	if idMessages == nil {
		err := errors.New("messages must be a map[int]string")
		return nil, err
	}

	// Defaults

	callerSkip := 1

	// Look for specific flags in details.

	if len(interfaces) > 0 {
		for _, value := range interfaces {
			switch typedValue := value.(type) {
			case messagelocation.CallerSkip:
				callerSkip = int(typedValue)
			}
		}
	}

	// Construct the components of the messagelogger.

	messageDate := &messagedate.MessageDateSenzing{}
	messageDetails := &messagedetails.MessageDetailsSenzing{}
	messageDuration := &messageduration.MessageDurationSenzing{}
	messageErrors := &messageerrors.MessageErrorsSenzing{}
	messageFormat := &messageformat.MessageFormatSenzing{}
	messageId := &messageid.MessageIdSenzing{
		MessageIdTemplate: fmt.Sprintf("senzing-%04d", componentIdentifier) + "%04d",
	}
	messageLevel := &messagelevel.MessageLevelSenzing{
		DefaultLogLevel: logger.LevelInfo,
		IdLevelRanges:   messagelevel.IdLevelRanges,
	}
	messageLocation := &messagelocation.MessageLocationSenzing{
		CallerSkip: callerSkip,
	}
	messageStatus := &messagestatus.MessageStatusSenzing{
		IdStatuses: messagelevel.IdLevelRangesAsString,
	}
	messageText := &messagetext.MessageTextSenzing{
		IdMessages: idMessages,
	}
	messageTime := &messagetime.MessageTimeSenzing{}

	var newInterfaces = []interface{}{
		messageDate,
		messageDetails,
		messageDuration,
		messageErrors,
		messageFormat,
		messageId,
		messageLevel,
		messageLocation,
		messageStatus,
		messageText,
		messageTime,
	}

	// Add other user-supplied interfaces to newInterfaces.

	newInterfaces = append(newInterfaces, interfaces...)

	// Using a Factory Pattern, build the messagelogger.

	return New(newInterfaces...)
}

/*
The NewSenzingApiLogger function creates a new instance of MessageLoggerInterface
that is tailored for the Senzing SDK implementation.
*/
func NewSenzingApiLogger(componentIdentifier int, idMessages map[int]string, idStatuses map[int]string, interfaces ...interface{}) (MessageLoggerInterface, error) {
	messageLevel := &messagelevel.MessageLevelSenzingApi{
		DefaultLogLevel: logger.LevelInfo,
		IdLevelRanges:   messagelevel.IdLevelRanges,
		IdStatuses:      idStatuses,
	}
	messageStatus := &messagestatus.MessageStatusSenzingApi{
		IdStatuses: idStatuses,
	}
	messageLocation := &messagelocation.MessageLocationSenzing{
		CallerSkip: 4,
	}
	var newInterfaces = []interface{}{
		messageLevel,
		messageLocation,
		messageStatus,
	}

	// Add other user-supplied interfaces to newInterfaces.

	newInterfaces = append(newInterfaces, interfaces...)

	// Using a Factory Pattern, build the messagelogger.

	return NewSenzingLogger(componentIdentifier, idMessages, newInterfaces...)
}

// ----------------------------------------------------------------------------
// Functions
// ----------------------------------------------------------------------------

/*
The GetLogLevel will return the current system setting for the log level.
*/
func GetLogLevel() (Level, error) {
	var err error = nil
	if !isSystemLogLevelSet {
		err = fmt.Errorf("system log level not set")
	}
	return systemLogLevel, err
}

/*
The GetLogLevelAsString will return the current system setting for the log level as a string.
*/
func GetLogLevelAsString() (string, error) {
	logLevel, err := GetLogLevel()
	if err != nil {
		return "", err
	}

	levelName, ok := logger.LevelToTextMap[logger.Level(logLevel)]
	if !ok {
		err = fmt.Errorf("logLevel not found")
	}

	return levelName, err
}

/*
The SetLogLevel will set the current system setting for the log level.
*/
func SetLogLevel(level Level) error {
	var err error = nil

	// Make this thread-safe.

	lock.Lock()
	defer lock.Unlock()

	// Propagate level across all messageLoggers.

	isSystemLogLevelSet = true
	systemLogLevel = level
	for _, messageLogger := range messageLoggerObservers {
		messageLogger.SetLogLevel(systemLogLevel)
	}
	return err
}

/*
The SetLogLevelFromString will set the current system setting for the log level.
*/
func SetLogLevelFromString(levelName string) error {

	level, ok := logger.TextToLevelMap[levelName]
	if ok {
		return SetLogLevel(Level(level))
	}

	return fmt.Errorf("unknown log level name: %s", levelName)
}
