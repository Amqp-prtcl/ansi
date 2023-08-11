package ansi

const (
	saveScreenPos = CSI + "s"
	loadScreenPos = CSI + "u"

	showCursor = CSI + "?25h"
	hideCursor = CSI + "?25l"

	enableReportFocus  = CSI + "?1004h"
	disableReportFocus = CSI + "?1004l"

	enableAlterBuf  = CSI + "?1049h"
	disableAlterBuf = CSI + "?1049l"
)

// Saves the cursor position/state in SCO console mode.
// In vertical split screen mode, instead used to set
// (as CSI n ; n s) or reset left and right margins.[32]
func SaveCursorPosition() {
	write(saveScreenPos)
}

// Restores the cursor position/state in SCO console mode.
func LoadCursorPosition() {
	write(loadScreenPos)
}

func ShowCursor() {
	write(showCursor)
}

func HideCursor() {
	write(hideCursor)
}

// Enable reporting focus. Reports whenever terminal emulator
// enters or exits focus as ESC [I and ESC [O, respectively.
func EnableReportFocus() {
	write(enableReportFocus)
}

func DisableReportFocus() {
	write(disableReportFocus)
}

// Enable alternative screen buffer
func EnableAlternateBuffer() {
	write(enableAlterBuf)
}

// Disable alternative screen buffer
func DisableAlternateBuffer() {
	write(disableAlterBuf)
}
