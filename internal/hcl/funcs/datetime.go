package funcs

import (
	"time"

	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
	"github.com/zclconf/go-cty/cty/function/stdlib"
)

// MockTimestampFunc constructs a function that returns a string representation of a static timestamp.
// We keep this as a static value so that it is deterministic when generating cost estimates.
var MockTimestampFunc = function.New(&function.Spec{
	Params: []function.Parameter{},
	Type:   function.StaticReturnType(cty.String),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		return cty.StringVal(time.Date(2023, 3, 14, 15, 9, 29, 0, time.UTC).UTC().Format(time.RFC3339)), nil
	},
})

// TimeAddFunc constructs a function that adds a duration to a timestamp, returning a new timestamp.
var TimeAddFunc = function.New(&function.Spec{
	Params: []function.Parameter{
		{
			Name: "timestamp",
			Type: cty.String,
		},
		{
			Name: "duration",
			Type: cty.String,
		},
	},
	Type: function.StaticReturnType(cty.String),
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		ts, err := time.Parse(time.RFC3339, args[0].AsString())
		if err != nil {
			return cty.UnknownVal(cty.String), err
		}
		duration, err := time.ParseDuration(args[1].AsString())
		if err != nil {
			return cty.UnknownVal(cty.String), err
		}

		return cty.StringVal(ts.Add(duration).Format(time.RFC3339)), nil
	},
})

// Timestamp returns a string representation of a static timestamp.
//
// In the Terraform language, timestamps are conventionally represented as
// strings using RFC 3339 "Date and Time format" syntax, and so timestamp
// returns a string in this format.
func Timestamp() (cty.Value, error) {
	return MockTimestampFunc.Call([]cty.Value{})
}

// TimeAdd adds a duration to a timestamp, returning a new timestamp.
//
// In the Terraform language, timestamps are conventionally represented as
// strings using RFC 3339 "Date and Time format" syntax. Timeadd requires
// the timestamp argument to be a string conforming to this syntax.
//
// `duration` is a string representation of a time difference, consisting of
// sequences of number and unit pairs, like `"1.5h"` or `1h30m`. The accepted
// units are `ns`, `us` (or `µs`), `"ms"`, `"s"`, `"m"`, and `"h"`. The first
// number may be negative to indicate a negative duration, like `"-2h5m"`.
//
// The result is a string, also in RFC 3339 format, representing the result
// of adding the given direction to the given timestamp.
func TimeAdd(timestamp cty.Value, duration cty.Value) (cty.Value, error) {
	return TimeAddFunc.Call([]cty.Value{timestamp, duration})
}

// FormatDateFunc is a wrapper around the stdlib.FormatDateFunc function that
// returns the current date if the input date is invalid. This is useful in cases
// where the date is invalid because of mocking/incomplete data because Diginfra
// cannot infer the values from the IaC.
var FormatDateFunc = function.New(&function.Spec{
	Description:  stdlib.FormatDateFunc.Description(),
	Params:       stdlib.FormatDateFunc.Params(),
	Type:         function.StaticReturnType(cty.String),
	RefineResult: refineNonNull,
	Impl: func(args []cty.Value, retType cty.Type) (cty.Value, error) {
		val, err := stdlib.FormatDateFunc.Call(args)
		var argsError function.ArgError
		if errors.As(err, &argsError) {
			// we have a problem with the passed in date, lets try and keep to the original
			// formatting arg and use the current date as the second arg.
			if argsError.Index == 1 {
				args[1] = cty.StringVal(time.Now().UTC().Format(time.RFC3339))
				val, err = stdlib.FormatDateFunc.Call(args)
				if err == nil {
					return val, nil
				}
			}
		}

		// if we can't infer the error then return the current date as an RFC 3339
		// string.
		if err != nil {
			return cty.StringVal(time.Now().UTC().Format(time.RFC3339)), nil
		}

		return val, nil
	},
})
