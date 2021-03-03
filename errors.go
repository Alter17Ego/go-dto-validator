package validator

import "errors"

const (
	STRUCT_ERROR num_error = iota
	ISREQUIRED
	INVALID_INT
	INVALID_EMAIL
	INVALID_TYPE
	INVALID_STRING
	MIN
	MAX
	STARTS_WITH
	ENDS_WITH
	VALUE_OF
	INVALID_CONVERSION_INT
	INVALID_ALPHABET
	INVALID_ALPHANUMERIC
)

func (num num_error) String() string {
	error := [...]string{
		"Not a struct type or memory address of struct!",
		"Is required!",
		"Expected number value!",
		"Email format invalid",
		"Invalid type of tag!",
		"Expected string value!",
		"Should be greater than",
		"Should be less than",
		"Should starts with",
		"Should ends with",
		"Missing value, value should be one of",
		"Invalid integer conversion for number type",
		"Alphabet format invalid!",
		"Alphanumeric format invalid!",
	}
	return error[num]
}

func myerr(num num_error, args ...interface{}) error {

	new_err := errors.New(num.String())
	struct_name := ""
	msg := ""
	if len(args) == 1 {
		struct_name = " ( " + args[0].(string) + " )"
		msg = struct_name
	}
	if len(args) == 2 {
		struct_name = " ( " + args[0].(string) + " )"
		struct_value := " " + args[1].(string) + " "
		msg = struct_value + struct_name
	}
	new_err = errors.New(num.String() + msg)

	return new_err
}
