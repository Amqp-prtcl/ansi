package ansi

const (
	clearCursorToEnd     = CSI + "0J"
	clearStartToCursor   = CSI + "1J"
	clearScreen          = CSI + "2J"
	clearScreenAndScroll = CSI + "3J"
)

func ClearCursorToEnd() {
	write(clearCursorToEnd)
}

func ClearStartToCursor() {
	write(clearStartToCursor)
}

func ClearScreen() {
	write(clearScreen)
}

func ClearScreenAndScroll() {
	write(clearScreenAndScroll)
}

// Reports the cursor position (CPR) by transmitting
// ESC[n;mR, where n is the row and m is the column.
//
// os.Stdin must be empty
func GetCursorPos() {
	panic("not implemented yet (waiting for the termois lib)")
}
