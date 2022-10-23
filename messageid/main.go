// The messageid package...
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messageid

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageIdInterface interface {
	GetMessageId(errorNumber int) (string, error)
}
