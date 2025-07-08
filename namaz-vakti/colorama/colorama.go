package colorama

import "fmt"

var Reset string = "\x1b[0m"

var Fore = struct {
	Reset         string
	Black         string
	Red           string
	Green         string
	Yellow        string
	Blue          string
	Magenta       string
	Cyan          string
	White         string
	Gray          string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
	RGB           func(r, g, b int) string
	Bit8          func(n int) string
}{
	Reset:         "\x1b[39m",
	Black:         "\x1b[30m",
	Red:           "\x1b[31m",
	Green:         "\x1b[32m",
	Yellow:        "\x1b[33m",
	Blue:          "\x1b[34m",
	Magenta:       "\x1b[35m",
	Cyan:          "\x1b[36m",
	White:         "\x1b[37m",
	Gray:          "\x1b[90m",
	BrightRed:     "\x1b[91m",
	BrightGreen:   "\x1b[92m",
	BrightYellow:  "\x1b[93m",
	BrightBlue:    "\x1b[94m",
	BrightMagenta: "\x1b[95m",
	BrightCyan:    "\x1b[96m",
	BrightWhite:   "\x1b[97m",
	RGB: func(r, g, b int) string {
		return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
	},
	Bit8: func(n int) string {
		return fmt.Sprintf("\x1b[38;5;%dm", n)
	},
}

var Back = struct {
	Reset         string
	Black         string
	Red           string
	Green         string
	Yellow        string
	Blue          string
	Magenta       string
	Cyan          string
	White         string
	Gray          string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
	RGB           func(r, g, b int) string
	Bit8          func(n int) string
}{
	Reset:         "\x1b[49m",
	Black:         "\x1b[40m",
	Red:           "\x1b[41m",
	Green:         "\x1b[42m",
	Yellow:        "\x1b[43m",
	Blue:          "\x1b[44m",
	Magenta:       "\x1b[45m",
	Cyan:          "\x1b[46m",
	White:         "\x1b[47m",
	Gray:          "\x1b[100m",
	BrightRed:     "\x1b[101m",
	BrightGreen:   "\x1b[102m",
	BrightYellow:  "\x1b[103m",
	BrightBlue:    "\x1b[104m",
	BrightMagenta: "\x1b[105m",
	BrightCyan:    "\x1b[106m",
	BrightWhite:   "\x1b[107m",
	RGB: func(r, g, b int) string {
		return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
	},
	Bit8: func(n int) string {
		return fmt.Sprintf("\x1b[48;5;%dm", n)
	},
}

var Style = struct {
	Bold          string
	Italic        string
	Underline     string
	Blink         string
	RapidBlink    string
	Reverse       string
	Hidden        string
	Strike        string
	NotBold       string
	NotItalic     string
	NotUnderline  string
	NotBlink      string
	NotReverse    string
	NotHidden     string
	NotStrike     string
	Overline      string
	NotOverline   string
}{
	Bold:         "\x1b[1m",
	Italic:       "\x1b[3m",
	Underline:    "\x1b[4m",
	Blink:        "\x1b[5m",
	RapidBlink:   "\x1b[6m",
	Reverse:      "\x1b[7m",
	Hidden:       "\x1b[8m",
	Strike:       "\x1b[9m",
	NotBold:      "\x1b[22m",
	NotItalic:    "\x1b[23m",
	NotUnderline: "\x1b[24m",
	NotBlink:     "\x1b[25m",
	NotReverse:   "\x1b[27m",
	NotHidden:    "\x1b[28m",
	NotStrike:    "\x1b[29m",
	Overline:     "\x1b[53m",
	NotOverline:  "\x1b[55m",
}
