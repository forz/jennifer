package main

var Keywords = []string{"break", "default", "func", "select", "defer", "go", "struct", "chan", "else", "goto", "const", "fallthrough", "range", "type", "continue", "var"}

// "return", "map", "switch", "for", "interface", "case" and "if" are special cases
// "import" and "package" are handled automatically, so not needed.

var Types = []string{"bool", "byte", "complex64", "complex128", "error", "float32", "float64", "int", "int8", "int16", "int32", "int64", "rune", "string", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr"}

var Constants = []string{"true", "false", "iota"}

var Zero = []string{"nil"}

var Errs = []string{"err"}

var Identifiers = append(append(append(append([]string{}, Types...), Constants...), Zero...), Errs...)

var Functions = []string{"append", "cap", "close", "complex", "copy", "delete", "imag", "len", "make", "new", "panic", "print", "println", "real", "recover"}

var Groups = []struct {
	Name      string
	Desc      string
	List      bool
	Open      string
	Close     string
	Separator string
}{
	{
		Name:      "Parens",
		Desc:      "parenthesis",
		List:      false,
		Open:      "(",
		Close:     ")",
		Separator: "",
	},
	{
		Name:      "List",
		Desc:      "a comma separated list",
		List:      true,
		Open:      "",
		Close:     "",
		Separator: ",",
	},
	{
		Name:      "Values",
		Desc:      "curly braces containing a comma separated list",
		List:      true,
		Open:      "{",
		Close:     "}",
		Separator: ",",
	},
	{
		Name:      "Index",
		Desc:      "square brackets containing a colon separated list",
		List:      true,
		Open:      "[",
		Close:     "]",
		Separator: ":",
	},
	{
		Name:      "Block",
		Desc:      "curly braces containing a statement list",
		List:      true,
		Open:      "{",
		Close:     "}",
		Separator: "\n",
	},
	{
		Name:      "Call",
		Desc:      "parenthesis containing a comma separated list",
		List:      true,
		Open:      "(",
		Close:     ")",
		Separator: ",",
	},
	{
		Name:      "Params",
		Desc:      "parenthesis containing a comma separated list",
		List:      true,
		Open:      "(",
		Close:     ")",
		Separator: ",",
	},
	{
		Name:      "CaseBlock",
		Desc:      "a statement list preceeded by a colon",
		List:      true,
		Open:      ":",
		Close:     "",
		Separator: "\n",
	},
	{
		Name:      "Assert",
		Desc:      "a type assertion",
		List:      false,
		Open:      ".(",
		Close:     ")",
		Separator: "",
	},
	{
		Name:      "Map",
		Desc:      "the map keyword, followed by square brackets",
		List:      false,
		Open:      "map[",
		Close:     "]",
		Separator: "",
	},
	{
		Name:      "If",
		Desc:      "the if keyword, followed by a semicolon separated list",
		List:      true,
		Open:      "if ",
		Close:     "",
		Separator: ";",
	},
	{
		Name:      "Return",
		Desc:      "the return keyword, followed by a comma separated list",
		List:      true,
		Open:      "return ",
		Close:     "",
		Separator: ",",
	},
	{
		Name:      "For",
		Desc:      "the for keyword, followed by a semicolon separated list",
		List:      true,
		Open:      "for ",
		Close:     "",
		Separator: ";",
	},
	{
		Name:      "Switch",
		Desc:      "the switch keyword, followed by a semicolon separated list",
		List:      true,
		Open:      "switch ",
		Close:     "",
		Separator: ";",
	},
	{
		Name:      "Interface",
		Desc:      "the interface keyword, followed by curly braces containing a statement list",
		List:      true,
		Open:      "interface{",
		Close:     "}",
		Separator: "\n",
	},
	{
		Name:      "Case",
		Desc:      "the case keyword, followed by a comma separated list",
		List:      true,
		Open:      "case ",
		Close:     "",
		Separator: ",",
	},
}
