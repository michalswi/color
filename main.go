package color

import (
	"fmt"
)

// escape is the ASCII escape character
const escape = "\x1b"

// Define color constants
const (
	NONE   = iota // No color
	RED           // Red color
	GREEN         // Green color
	YELLOW        // Yellow color
	BLUE          // Blue color
	PURPLE        // Purple color
)

// color function takes an integer c as input and
// returns an escape sequence string for that color.
// If c is NONE, it returns the reset sequence.
func color(c int) string {
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
	return color(c) + text + color(NONE)
}
