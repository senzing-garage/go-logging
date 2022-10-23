// The messageid package...
//
// The purpose of a logger object is:
//   - ...
//   - ...
//   - ...
package messagetext

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageTextInterface interface {
	MessageText(errorNumber int, details ...interface{}) (string, error)
}
