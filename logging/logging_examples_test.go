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

func ExampleNewSenzingLogger() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	componentID := 9999
	idMessages := map[int]string{
		2001: "My message",
	}
	logger, err := NewSenzingLogger(componentID, idMessages)
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001)
	// Output:
}

func ExampleBasicLogging_IsTrace() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("TRACE")
	if err != nil {
		fmt.Println(err)
	}
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

func ExampleBasicLogging_IsDebug() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("DEBUG")
	if err != nil {
		fmt.Println(err)
	}
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

func ExampleBasicLogging_IsInfo() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("INFO")
	if err != nil {
		fmt.Println(err)
	}
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

func ExampleBasicLogging_IsWarn() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("WARN")
	if err != nil {
		fmt.Println(err)
	}
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

func ExampleBasicLogging_IsError() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("ERROR")
	if err != nil {
		fmt.Println(err)
	}
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

func ExampleBasicLogging_IsFatal() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("FATAL")
	if err != nil {
		fmt.Println(err)
	}
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

func ExampleBasicLogging_IsPanic() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	err = logger.SetLogLevel("PANIC")
	if err != nil {
		fmt.Println(err)
	}
	if logger.IsFatal() {
		fmt.Println("FATAL active")
	}
	if logger.IsPanic() {
		fmt.Println("PANIC active")
	}
	// Output:
	// PANIC active
}

func ExampleBasicLogging_Log_new() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	logger, err := New()
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001, "Bob", "Jane") // Note that 2000's are INFO messages.
	// Output:
}

func ExampleBasicLogging_Log_newSenzingLogger() {
	// For more information, visit https://github.com/senzing-garage/go-logging/blob/main/logging/logging_examples_test.go
	componentID := 9999
	idMessages := map[int]string{
		2001: "%s works with %s",
	}
	logger, err := NewSenzingLogger(componentID, idMessages)
	if err != nil {
		fmt.Println(err)
	}
	logger.Log(2001, "Bob", "Jane") // Note that 2000's are INFO messages.
	// Output:
}
