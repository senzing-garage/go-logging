package logging

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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
	name                          string
	messageNumber                 int
	options                       []interface{}
	details                       []interface{}
	expectedNew                   string
	expectedNewSenzingToolsLogger string
}{
	{
		name:                          "logging-0001",
		messageNumber:                 1,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   ``,
		expectedNewSenzingToolsLogger: ``,
	},
	{
		name:                          "logging-1001",
		messageNumber:                 1001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   ``,
		expectedNewSenzingToolsLogger: ``,
	},
	{
		name:                          "logging-2001",
		messageNumber:                 2001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"INFO","text":"INFO: Bob works with Jane","id":"2001","location":"In func1() at logging_test.go:167","details":{"1":"Bob","2":"Jane"}}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"INFO","text":"INFO: Bob works with Jane","id":"senzing-99972001","location":"In func1() at logging_test.go:180","details":{"1":"Bob","2":"Jane"}}` + "\n",
	},
	{
		name:                          "logging-3001",
		messageNumber:                 3001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"WARN","text":"WARN: Bob works with Jane","id":"3001","location":"In func1() at logging_test.go:167","details":{"1":"Bob","2":"Jane"}}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"WARN","text":"WARN: Bob works with Jane","id":"senzing-99973001","location":"In func1() at logging_test.go:180","details":{"1":"Bob","2":"Jane"}}` + "\n",
	},
	{
		name:                          "logging-4001",
		messageNumber:                 4001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"4001","location":"In func1() at logging_test.go:167","details":{"1":"Bob","2":"Jane"}}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"ERROR","text":"ERROR: Bob works with Jane","id":"senzing-99974001","location":"In func1() at logging_test.go:180","details":{"1":"Bob","2":"Jane"}}` + "\n",
	},
	{
		name:                          "logging-5001",
		messageNumber:                 5001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"5001","location":"In func1() at logging_test.go:167","details":{"1":"Bob","2":"Jane"}}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"FATAL","text":"FATAL: Bob works with Jane","id":"senzing-99975001","location":"In func1() at logging_test.go:180","details":{"1":"Bob","2":"Jane"}}` + "\n",
	},
	{
		name:                          "logging-6001",
		messageNumber:                 6001,
		options:                       []interface{}{getOptionIdMessages(), getOptionCallerSkip(), getOptionOutput(), getOptionTimeHidden()},
		details:                       []interface{}{"Bob", "Jane"},
		expectedNew:                   `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"6001","location":"In func1() at logging_test.go:167","details":{"1":"Bob","2":"Jane"}}` + "\n",
		expectedNewSenzingToolsLogger: `{"level":"PANIC","text":"PANIC: Bob works with Jane","id":"senzing-99976001","location":"In func1() at logging_test.go:180","details":{"1":"Bob","2":"Jane"}}` + "\n",
	},
}

var (
	componentId  int           = 9997
	outputString *bytes.Buffer = new(bytes.Buffer)
)

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

func getOptionCallerSkip() *OptionCallerSkip {
	return &OptionCallerSkip{
		Value: 3,
	}
}

func getOptionIdMessages() *OptionIdMessages {
	return &OptionIdMessages{
		Value: idMessagesTest,
	}
}

func getOptionOutput() *OptionOutput {
	return &OptionOutput{
		Value: outputString,
	}
}

func getOptionTimeHidden() *OptionTimeHidden {
	return &OptionTimeHidden{
		Value: true,
	}
}

func testError(test *testing.T, testObject LoggingInterface, err error) {
	if err != nil {
		assert.Fail(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

// -- Test New() method ---------------------------------------------------------

func TestLoggingImpl_New(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-New", func(test *testing.T) {
			testObject, err := New(testCase.options...)
			testError(test, testObject, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNew, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}

func TestLoggingImpl_NewSenzingToolsLogger(test *testing.T) {
	outputString.Reset()
	for _, testCase := range testCasesForMessage {
		test.Run(testCase.name+"-NewSenzingToolsLogger", func(test *testing.T) {
			testObject, err := NewSenzingToolsLogger(componentId, idMessagesTest, testCase.options...)
			testError(test, testObject, err)
			testObject.Log(testCase.messageNumber, testCase.details...)
			assert.Equal(test, testCase.expectedNewSenzingToolsLogger, outputString.String(), testCase.name)
			outputString.Reset()
		})
	}
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleLoggingImpl_Log() {
	// For more information, visit https://github.com/Senzing/go-messaging/blob/main/logging/logging_test.go
	// example, err := New()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Print(example.NewJson(2001, "Bob", "Jane", getTimestamp(), getOptionCallerSkip()))
	//Output: {"time":"2000-01-01 00:00:00 +0000 UTC","level":"INFO","id":"senzing-99992001","location":"In ExampleloggingImpl_NewJson() at logging_test.go:205","details":{"1":"Bob","2":"Jane"}}
}
