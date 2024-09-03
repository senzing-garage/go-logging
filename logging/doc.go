/*
Package logging is used to create loggers with guard methods.

The [guard] methods can be used to avoid calls to logging functions
that would be ignored due to logging level.
This avoids long computations for log messages that would be discarded anyway.

[guard]: https://en.wikipedia.org/wiki/Guard_(computer_science)
*/
package logging
