package ansi

// Moves the cursor n (default 1) cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorUp(n int) {
	write(CSI + itoa(n) + "A")
}

// Moves the cursor n (default 1) cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorDown(n int) {
	write(CSI + itoa(n) + "B")
}

// Moves the cursor n (default 1) cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorForward(n int) {
	write(CSI + itoa(n) + "C")
}

// Moves the cursor n (default 1) cells in the given direction.
// If the cursor is already at the edge of the screen, this has no effect.
func CursorBackward(n int) {
	write(CSI + itoa(n) + "D")
}

// Moves cursor to beginning of the line n (default 1) lines down.
func CursorNextLine(n int) {
	write(CSI + itoa(n) + "E")
}

// Moves cursor to beginning of the line n (default 1) lines up
func CursorPreviousLine(n int) {
	write(CSI + itoa(n) + "F")
}

// Moves the cursor to column n (default 1).
func CursorHorizontalAbsolute(n int) {
	write(CSI + itoa(n) + "G")
}

// Moves the cursor to row n, column m. The values are 1-based.
func CursorPosition(n int, m int) {
	write(CSI + itoa(n) + ";" + itoa(n) + "H")
}

// Reports the cursor position (CPR) by transmitting
// ESC[n;mR, where n is the row and m is the column.
//
// os.Stdin must be empty
func GetCursorPos() {
	panic("not implemented yet (waiting for the termois lib)")

	// TODO
	// set ICANON and ECHO lflags to false
	// ask and parse input
	// return to previous state
}
