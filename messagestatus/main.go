// The messagestatus package is a set of methods logging messages.
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagestatus

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageStatusInterface interface {
	CalculateMessageStatus(errorNumber int, text string) (string, error)
}
