package color

import (
	"fmt"
)

// escape is the ASCII escape character
const escape = "\x1b"

// Define color constants
const (
	NONE   = iota // No color
	RED           // Red color (int 1)
	GREEN         // Green color (int 2)
	YELLOW        // Yellow color (int 3)
	BLUE          // Blue color (int 4)
	PURPLE        // Purple color (int 5)
	MINGLE        // Mingle color (int 6)
	WHITE         // White color (int 7)
)

// colorize function takes an integer c as input and
// returns an escape sequence string for that color.
// If c is NONE, it returns the reset sequence.
func colorize(c int) string {
	if c == NONE {
		// Return the reset sequence
		return fmt.Sprintf("%s[%dm", escape, c)
	}

	// Return the escape sequence for color c
	return fmt.Sprintf("%s[3%dm", escape, c)
}

// Format function takes a color c and a string text as input.
// It returns the string text wrapped with the escape sequences
// for color c and the reset sequence.
func Format(c int, text string) string {
	// Wrap text with color c and reset sequence
	return colorize(c) + text + colorize(NONE)
}
