package slogconfig

import (
	"io"

	"golang.org/x/exp/slog"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// Reference: https://go.googlesource.com/exp/+/refs/heads/master/slog/example_custom_levels_test.go
func SenzingJsonHandler(writer io.Writer) *slog.JSONHandler {
	result := slog.HandlerOptions{
		ReplaceAttr: func(groups []string, slogAttribute slog.Attr) slog.Attr {
			if slogAttribute.Key == slog.LevelKey {

				// Handle custom level values.
				level := slogAttribute.Value.Any().(slog.Level)
				// This could also look up the name from a map or other structure, but
				// TODO: For maximum performance, the string values should be constants.
				switch {
				case level < LevelDebug:
					slogAttribute.Value = slog.StringValue("TRACE")
				case level < LevelInfo:
					slogAttribute.Value = slog.StringValue("DEBUG")
				case level < LevelWarn:
					slogAttribute.Value = slog.StringValue("INFO")
				case level < LevelError:
					slogAttribute.Value = slog.StringValue("WARN")
				case level < LevelFatal:
					slogAttribute.Value = slog.StringValue("ERROR")
				case level < LevelPanic:
					slogAttribute.Value = slog.StringValue("FATAL")
				default:
					slogAttribute.Value = slog.StringValue("PANIC")
				}
			}
			return slogAttribute
		},
	}.NewJSONHandler(writer)

	return result
}
