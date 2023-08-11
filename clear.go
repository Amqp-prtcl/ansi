package ansi

const (
	clearCursorToEnd   = CSI + "0J"
	clearStartToCursor = CSI + "1J"
	clearScreen        = CSI + "2J"
	clearScreenHard    = CSI + "3J"

	clearLineEnd   = CSI + "0K"
	clearLineStart = CSI + "1K"
	clearLine      = CSI + "2K"
)

// Clears from cursor to end of screen.
func ClearCursorToEnd() {
	write(clearCursorToEnd)
}

// Clears from cursor to beginning of the screen.
func ClearStartToCursor() {
	write(clearStartToCursor)
}

// Clears the entire screen (and moves cursor to upper left on DOS ANSI.SYS).
func ClearScreen() {
	write(clearScreen)
}

// Clears the entire screen and delete all
// lines saved in the scrollback buffer
func ClearScreenHard() {
	write(clearScreenHard)
}

// Clears from cursor to beginning of the line.
func ClearLineCorusorToEnd() {
	write(clearLineEnd)
}

// Clears from cursor to the end of the line.
func ClearLineStartToCursor() {
	write(clearLineStart)
}

// Clears entire line.
func ClearLine() {
	write(clearLine)
}

// Clears the screen and goes to screen position 1;1
func ClearScreenAndGoHome() {
	CursorPosition(1, 1)
	ClearScreen()
}

// Same as ClearScreenAndGoHome but also clears scroll buffer
func ClearScreenAndGoHomeHard() {
	CursorPosition(1, 1)
	ClearScreenHard()
}
