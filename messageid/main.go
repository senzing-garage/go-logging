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
	MessageId(errorNumber int, details ...interface{}) (string, error)
	SetMessageIdTemplate(string)
}
