package ANSITerminal

const (
	cRed     = "\x1B[0;31m"
	cGreen   = "\x1B[0;32m"
	cYellow  = "\x1B[0;33m"
	cBlue    = "\x1B[0;34m"
	cMagenta = "\x1B[0;35m"
	cCyan    = "\x1B[0;36m"
	cWhite   = "\x1B[0;37m"
	cGrey    = "\x1B[0;90m"
	cBlack   = "\x1B[0;30m"
	cDefault = "\x1B[39m"
)

func Red(str string) string {
	return cRed + str + cDefault
}

func Green(str string) string {
	return cGreen + str + cDefault
}

func Yellow(str string) string {
	return cYellow + str + cDefault
}

func Blue(str string) string {
	return cBlue + str + cDefault
}

func Magenta(str string) string {
	return cMagenta + str + cDefault
}

func Cyan(str string) string {
	return cCyan + str + cDefault
}

func White(str string) string {
	return cWhite + str + cDefault
}

func Grey(str string) string {
	return cGrey + str + cDefault
}

func Black(str string) string {
	return cBlack + str + cDefault
}
