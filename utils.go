package ansi

import "os"

func write(str string) (i int, err error) {
	return os.Stdin.WriteString(str)
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}

	var res string
	var neg bool
	if v < 0 {
		neg = true
		v = -v
	}

	for v != 0 {
		res = string((rune)(v%10)+'0') + res
	}
	if neg {
		res = "-" + res
	}
	return res
}
