package logging

import (
	"fmt"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNew() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001)
	// Output:
}

func ExampleNewSenzingToolsLogger() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	componentId := 9999
	idMessages := map[int]string{
		2001: "My message",
	}
	logger, err := NewSenzingToolsLogger(componentId, idMessages)
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001)
	// Output:
}

func ExampleLoggingImpl_IsTrace() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("TRACE")
	if logger.IsTrace() {
		fmt.Println("TRACE active")
	}
	if logger.IsDebug() {
		fmt.Println("DEBUG active")
	}
	// Output:
	// TRACE active
	// DEBUG active
}

func ExampleLoggingImpl_IsDebug() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("DEBUG")
	if logger.IsTrace() {
		fmt.Println("TRACE active")
	}
	if logger.IsDebug() {
		fmt.Println("DEBUG active")
	}
	if logger.IsInfo() {
		fmt.Println("INFO active")
	}
	// Output:
	// DEBUG active
	// INFO active
}

func ExampleLoggingImpl_IsInfo() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("INFO")
	if logger.IsDebug() {
		fmt.Println("DEBUG active")
	}
	if logger.IsInfo() {
		fmt.Println("INFO active")
	}
	if logger.IsWarn() {
		fmt.Println("WARN active")
	}
	// Output:
	// INFO active
	// WARN active
}

func ExampleLoggingImpl_IsWarn() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("WARN")
	if logger.IsInfo() {
		fmt.Println("INFO active")
	}
	if logger.IsWarn() {
		fmt.Println("WARN active")
	}
	if logger.IsError() {
		fmt.Println("ERROR active")
	}
	// Output:
	// WARN active
	// ERROR active
}

func ExampleLoggingImpl_IsError() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("ERROR")
	if logger.IsWarn() {
		fmt.Println("WARN active")
	}
	if logger.IsError() {
		fmt.Println("ERROR active")
	}
	if logger.IsFatal() {
		fmt.Println("FATAL active")
	}
	// Output:
	// ERROR active
	// FATAL active
}

func ExampleLoggingImpl_IsFatal() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("FATAL")
	if logger.IsError() {
		fmt.Println("ERROR active")
	}
	if logger.IsFatal() {
		fmt.Println("FATAL active")
	}
	if logger.IsPanic() {
		fmt.Println("PANIC active")
	}
	// Output:
	// FATAL active
	// PANIC active
}

func ExampleLoggingImpl_IsPanic() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.SetLogLevel("PANIC")
	if logger.IsFatal() {
		fmt.Println("FATAL active")
	}
	if logger.IsPanic() {
		fmt.Println("PANIC active")
	}
	// Output:
	// PANIC active
}

func ExampleLoggingImpl_Log_new() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001, "Bob", "Jane") // Note that 2000's are INFO messages.
	// Output:
}

func ExampleLoggingImpl_Log_newSenzingToolsLogger() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	componentID := 9999
	idMessages := map[int]string{
		2001: "%s works with %s",
	}
	logger, err := NewSenzingToolsLogger(componentID, idMessages)
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001, "Bob", "Jane") // Note that 2000's are INFO messages.
	// Output:
}
