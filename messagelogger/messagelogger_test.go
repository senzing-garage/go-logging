package messagelogger

// const MessageIdFormat = "senzing-9999%04d"

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

// -- LogMessage --------------------------------------------------------------

// func TestLogMessage(test *testing.T) {
// 	LogMessage(MessageIdFormat, 2000, "Test message", "Variable1", "Variable2")
// }

// -- LogMessageFromError -----------------------------------------------------

// func TestLogMessageFromError(test *testing.T) {
// 	anError := errors.New("This is a new error")
// 	LogMessageFromError(MessageIdFormat, 2001, "Test message", anError, "Variable1", "Variable2")
// }

// -- LogMessageFromErrorUsingMap ---------------------------------------------

// func TestLogMessageFromErrorUsingMap(test *testing.T) {
// 	anError := errors.New("This is a new error")
// 	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA"}`
// 	detailsMap := map[string]interface{}{
// 		"FirstVariable":  "First value",
// 		"SecondVariable": "Second value",
// 		"TestClass":      test,
// 		"JSON":           jsonData,
// 	}
// 	LogMessageFromErrorUsingMap(MessageIdFormat, 2002, "Test message", anError, detailsMap)
// }

// -- LogMessageUsingMap ------------------------------------------------------

// func TestLogMessageUsingMap(test *testing.T) {

// 	jsonData := `{"SOCIAL_HANDLE": "flavorh", "DATE_OF_BIRTH": "4/8/1983", "ADDR_STATE": "LA"}`
// 	detailsMap := map[string]interface{}{
// 		"FirstVariable":  "First value",
// 		"SecondVariable": "Second value",
// 		"TestClass":      test,
// 		"JSON":           jsonData,
// 	}

// 	LogMessageUsingMap(MessageIdFormat, 2003, "Test message", detailsMap)
// }
