/*
Package messageid produces customized message identifiers.
*/
package messageid

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageIdInterface interface {
	MessageId(messageNumber int, details ...interface{}) (string, error)
	SetIdTemplate(string)
}
