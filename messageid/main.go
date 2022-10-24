/*
Package logger provides...
*/
package messageid

// ----------------------------------------------------------------------------
// Interfaces
// ----------------------------------------------------------------------------

type MessageIdInterface interface {
	MessageId(messageNumber int, details ...interface{}) (string, error)
	SetIdTemplate(string)
}
