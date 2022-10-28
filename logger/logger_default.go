/*
The logger package is a decorator over Go's log package.

The LoggerDefault implementation  provides a layer over go's log to
add Trace, Debug, Info, Warn, and Error levels.

It also implements IsXxxx() functions
that can be used as guards
(https://en.wikipedia.org/wiki/Guard_(computer_science))
*/
package logger

import (
	"fmt"
	"log"
	"strings"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

/*
The LoggerDefault type is for logging messages based on the following levels:
TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.
*/
type LoggerDefault struct {
	level   Level
	isDebug bool
	isError bool
	isFatal bool
	isInfo  bool
	isPanic bool
	isTrace bool
	isWarn  bool
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (logger *LoggerDefault) printf(debugLevelName string, format string, v ...interface{}) LoggerInterface {
	var message string
	calldepth := 3
	if format == noFormat {
		v := append(v, 0)
		copy(v[1:], v[0:])
		v[0] = debugLevelName + " "
		message = fmt.Sprint(v...)
	} else {
		message = fmt.Sprintf(debugLevelName+" "+format, v...)
	}
	log.Output(calldepth, message)
	return loggerInstance
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

func (logger *LoggerDefault) Debug(v ...interface{}) LoggerInterface {
	if logger.isDebug {
		logger.printf(LevelDebugName, noFormat, v...)
	}
	return logger
}

func (logger *LoggerDefault) Debugf(format string, v ...interface{}) LoggerInterface {
	if logger.isDebug {
		logger.printf(LevelDebugName, format, v...)
	}
	return logger
}

func (logger *LoggerDefault) Error(v ...interface{}) LoggerInterface {
	if logger.isError {
		logger.printf(LevelErrorName, noFormat, v...)
	}
	return logger
}

func (logger *LoggerDefault) Errorf(format string, v ...interface{}) LoggerInterface {
	if logger.isError {
		logger.printf(LevelErrorName, format, v...)
	}
	return logger
}

func (logger *LoggerDefault) Fatal(v ...interface{}) LoggerInterface {
	if logger.isFatal {
		logger.printf(LevelFatalName, noFormat, v...)
		log.Fatal("")
	}
	return logger
}

func (logger *LoggerDefault) Fatalf(format string, v ...interface{}) LoggerInterface {
	if logger.isFatal {
		logger.printf(LevelFatalName, format, v...)
		log.Fatal("")
	}
	return logger
}

func (logger *LoggerDefault) GetLogLevel() Level {
	return logger.level
}

func (logger *LoggerDefault) GetLogLevelAsString() string {
	return levelToTextMap[logger.level]
}

func (logger *LoggerDefault) Info(v ...interface{}) LoggerInterface {
	if logger.isInfo {
		logger.printf(LevelInfoName, noFormat, v...)
	}
	return logger
}

func (logger *LoggerDefault) Infof(format string, v ...interface{}) LoggerInterface {
	if logger.isInfo {
		logger.printf(LevelInfoName, format, v...)
	}
	return logger
}

func (logger *LoggerDefault) IsPanic() bool {
	return logger.isPanic
}

func (logger *LoggerDefault) IsFatal() bool {
	return logger.isFatal
}

func (logger *LoggerDefault) IsError() bool {
	return logger.isError
}

func (logger *LoggerDefault) IsWarn() bool {
	return logger.isWarn
}

func (logger *LoggerDefault) IsInfo() bool {
	return logger.isInfo
}

func (logger *LoggerDefault) IsDebug() bool {
	return logger.isDebug
}

func (logger *LoggerDefault) IsTrace() bool {
	return logger.isTrace
}

func (logger *LoggerDefault) Panic(v ...interface{}) LoggerInterface {
	if logger.isPanic {
		logger.printf(LevelPanicName, noFormat, v...)
		log.Panic("")
	}
	return logger
}

func (logger *LoggerDefault) Panicf(format string, v ...interface{}) LoggerInterface {
	if logger.isPanic {
		logger.printf(LevelPanicName, format, v...)
		log.Panic("")
	}
	return logger
}

func (logger *LoggerDefault) SetLogLevel(level Level) LoggerInterface {
	logger.level = level
	logger.isPanic = level <= LevelPanic
	logger.isFatal = level <= LevelFatal
	logger.isError = level <= LevelError
	logger.isWarn = level <= LevelWarn
	logger.isInfo = level <= LevelInfo
	logger.isDebug = level <= LevelDebug
	logger.isTrace = level <= LevelTrace
	return logger
}

func (logger *LoggerDefault) SetLogLevelFromString(levelString string) LoggerInterface {
	upperLevelString := strings.ToUpper(levelString)
	level, ok := textToLevelMap[upperLevelString]
	if !ok {
		level = LevelPanic
	}
	logger.SetLogLevel(level)
	return logger
}

func (logger *LoggerDefault) Trace(v ...interface{}) LoggerInterface {
	if logger.isTrace {
		logger.printf(LevelTraceName, noFormat, v...)
	}
	return logger
}

func (logger *LoggerDefault) Tracef(format string, v ...interface{}) LoggerInterface {
	if logger.isTrace {
		logger.printf(LevelTraceName, format, v...)
	}
	return logger
}

func (logger *LoggerDefault) Warn(v ...interface{}) LoggerInterface {
	if logger.isWarn {
		logger.printf(LevelWarnName, noFormat, v...)
	}
	return logger
}

func (logger *LoggerDefault) Warnf(format string, v ...interface{}) LoggerInterface {
	if logger.isWarn {
		logger.printf(LevelWarnName, format, v...)
	}
	return logger
}
