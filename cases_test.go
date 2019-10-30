package luhn

var testCases = []struct {
	description string
	input       string
	ok          bool
}{
	{
		"single digit string can not be valid",
		"1",
		false,
	},
	{
		"a single zero is invalid",
		"0",
		false,
	},
	{
		"a simple valid SIN that remains valid if reversed",
		"059",
		true,
	},
	{
		"a simple valid SIN that becomes invalid if reversed",
		"59",
		true,
	},
	{
		"a valid Canadian SIN",
		"055 444 285",
		true,
	},
	{
		"invalid Canadian SIN",
		"055 444 286",
		false,
	},
	{
		"invalid credit card",
		"8273 1232 7352 0569",
		false,
	},
	{
		"valid number with an even number of digits",
		"095 245 88",
		true,
	},
	{
		"valid string with a non-digit added at the end to become invalid",
		"059a",
		false,
	},
	{
		"valid string with punctuation included become invalid",
		"055-444-285",
		false,
	},
	{
		"valid string with symbols included become invalid",
		"055% 444% 285",
		false,
	},
	{
		"single zero with space is invalid",
		" 0",
		false,
	},
	{
		"more than a single zero is valid",
		"0000 0",
		true,
	},
	{
		"string with non-digit is invalid",
		":9",
		false,
	},
}
