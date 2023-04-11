package logging

import (
	"context"

	"github.com/senzing/go-messaging/messenger"
	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// loggingImpl is an type-struct for an implementation of the loggingInterface.
type LoggingImpl struct {
	Ctx       context.Context // Not a preferred practice, but used to simplify Log() calls.
	messenger messenger.MessengerInterface
	logger    *slog.Logger
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Private methods
// ----------------------------------------------------------------------------

func (logging *LoggingImpl) initialize() error {
	var err error = nil

	if logging.Ctx == nil {
		logging.Ctx = context.Background()
	}

	// TODO: Set logging level
	return err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Log method ...

Input
  - messageNumber: A message identifier which indexes into "idMessages".
  - details: Variadic arguments of any type to be added to the message.

Output
  - error
*/
func (logging *LoggingImpl) Log(messageNumber int, details ...interface{}) {
	message, logLevel, details := logging.messenger.NewSlogLevel(messageNumber, details...)
	// slog.Log(logging.Ctx, logLevel, message, details...)
	logging.logger.Log(logging.Ctx, logLevel, message, details...)
}
