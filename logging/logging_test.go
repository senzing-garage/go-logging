package logging

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
)

var idMessagesTest = map[int]string{
	0001: "TRACE: %s works with %s",
	1001: "DEBUG: %s works with %s",
	2001: "INFO: %s works with %s",
	3001: "WARN: %s works with %s",
	4001: "ERROR: %s works with %s",
	5001: "FATAL: %s works with %s",
	6001: "PANIC: %s works with %s",
}

var testCasesForMessage = []struct {
	name                string
	messageNumber       int
	options             []interface{}
	details             []interface{}
	expectedMessageJson string
	expectedMessageSlog []interface{}
	expectedText        string
	expectedSlogLevel   slog.Level
}{
	{
		name:                "logging-0001",
		messageNumber:       1,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"TRACE","id":"senzing-99990001","text":"TRACE: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99990001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "TRACE: Bob works with Jane",
		expectedSlogLevel:   LevelTraceSlog,
	},
	{
		name:                "logging-1001",
		messageNumber:       1001,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"DEBUG","id":"senzing-99991001","text":"DEBUG: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99991001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "DEBUG: Bob works with Jane",
		expectedSlogLevel:   LevelDebugSlog,
	},
	{
		name:                "logging-2001",
		messageNumber:       2001,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"INFO","id":"senzing-99992001","text":"INFO: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99992001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "INFO: Bob works with Jane",
		expectedSlogLevel:   LevelInfoSlog,
	},
	{
		name:                "logging-3001",
		messageNumber:       3001,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"WARN","id":"senzing-99993001","text":"WARN: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99993001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "WARN: Bob works with Jane",
		expectedSlogLevel:   LevelWarnSlog,
	},
	{
		name:                "logging-4001",
		messageNumber:       4001,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"ERROR","id":"senzing-99994001","text":"ERROR: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99994001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "ERROR: Bob works with Jane",
		expectedSlogLevel:   LevelErrorSlog,
	},
	{
		name:                "logging-5001",
		messageNumber:       5001,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"FATAL","id":"senzing-99995001","text":"FATAL: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99995001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "FATAL: Bob works with Jane",
		expectedSlogLevel:   LevelFatalSlog,
	},
	{
		name:                "logging-6001",
		messageNumber:       6001,
		options:             []interface{}{getOptionIdMessages(), getOptionCallerSkip()},
		details:             []interface{}{"Bob", "Jane", getTimestamp()},
		expectedMessageJson: `{"time":"2000-01-01 00:00:00 +0000 UTC","level":"PANIC","id":"senzing-99996001","text":"PANIC: Bob works with Jane","location":"In func1() at logging_test.go:173","details":{"1":"Bob","2":"Jane"}}`,
		expectedMessageSlog: []interface{}{"id", "senzing-99996001", "location", "In func1() at logging_test.go:186", "details", map[string]interface{}{"1": "Bob", "2": "Jane"}},
		expectedText:        "PANIC: Bob works with Jane",
		expectedSlogLevel:   LevelPanicSlog,
	},
}

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Internal functions - names begin with lowercase letter
// ----------------------------------------------------------------------------

func testError(test *testing.T, testObject LoggingInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

func getOptionIdMessages() *OptionIdMessages {
	return &OptionIdMessages{
		Value: idMessagesTest,
	}
}

func getOptionCallerSkip() *OptionCallerSkip {
	return &OptionCallerSkip{
		Value: 2,
	}
}

func getTimestamp() *MessageTime {
	return &MessageTime{
		Value: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

// -- Test New() method ---------------------------------------------------------

func TestLoggingImpl_NewJson(test *testing.T) {
	for _, testCase := range testCasesForMessage {
		if len(testCase.expectedMessageJson) > 0 {
			test.Run(testCase.name+"-NewJson", func(test *testing.T) {
				testObject, err := New(testCase.options...)
				testError(test, testObject, err)
				actual := testObject.NewJson(testCase.messageNumber, testCase.details...)
				assert.Equal(test, testCase.expectedMessageJson, actual, testCase.name)
			})
		}
	}
}

func TestLoggingImpl_NewSlogLevel(test *testing.T) {
	for _, testCase := range testCasesForMessage {
		if len(testCase.expectedMessageSlog) > 0 {
			test.Run(testCase.name+"-NewSlog", func(test *testing.T) {
				testObject, err := New(testCase.options...)
				testError(test, testObject, err)
				message, slogLevel, actual := testObject.NewSlogLevel(testCase.messageNumber, testCase.details...)
				assert.Equal(test, testCase.expectedText, message, testCase.name)
				assert.Equal(test, testCase.expectedSlogLevel, slogLevel, testCase.name)
				assert.Equal(test, testCase.expectedMessageSlog, actual, testCase.name)
			})
		}
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleLoggingImpl_NewJson() {
	// For more information, visit https://github.com/Senzing/go-messaging/blob/main/logging/logging_test.go
	example, err := New()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(example.NewJson(2001, "Bob", "Jane", getTimestamp(), getOptionCallerSkip()))
	//Output: {"time":"2000-01-01 00:00:00 +0000 UTC","level":"INFO","id":"senzing-99992001","location":"In ExampleloggingImpl_NewJson() at logging_test.go:205","details":{"1":"Bob","2":"Jane"}}
}
