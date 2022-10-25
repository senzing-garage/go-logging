/*
Package messageid customizes message identifieers.
*/
package messageid

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type MessageIdInterface interface {
	MessageId(messageNumber int, details ...interface{}) (string, error)
	SetIdTemplate(string)
}
