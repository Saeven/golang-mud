package saevenx

import (
	"regexp"
)

/*
 * Saeven COLOR CODES and VT SEQUENCES
 */
const MOD_CLEAR = "[0m" /* Resets Color (n) */
const MOD_BOLD = "[1m"
const MOD_FAINT = "[2m"
const MOD_UNDERLINE = "[4m"
const MOD_BLINK = "[5m"
const MOD_REVERSE = "[7m"

// Foreground Colors
const FG_BLACK = "[0;30m"   /* (l) */
const FG_RED = "[0;31m"     /* (r) */
const FG_GREEN = "[0;32m"   /* (g) */
const FG_YELLOW = "[0;33m"  /* (y) */
const FG_BLUE = "[0;34m"    /* (b) */
const FG_MAGENTA = "[0;35m" /* (m) */
const FG_CYAN = "[0;36m"    /* (c) */
const FG_WHITE = "[0;37m"   /* (w) */

// Bold Foreground Colors
const FG_B_BLACK = "[1;30m"   /* (L) */
const FG_B_RED = "[1;31m"     /* (R) */
const FG_B_GREEN = "[1;32m"   /* (G) */
const FG_B_YELLOW = "[1;33m"  /* (Y) */
const FG_B_BLUE = "[1;34m"    /* (B) */
const FG_B_MAGENTA = "[1;35m" /* (M) */
const FG_B_CYAN = "[1;36m"    /* (C) */
const FG_B_WHITE = "[1;37m"   /* (W) */

// Background Colors
const BG_BLACK = "[40m"
const BG_RED = "[41m"
const BG_GREEN = "[42m"
const BG_YELLOW = "[43m"
const BG_BLUE = "[44m"
const BG_MAGENTA = "[45m"
const BG_CYAN = "[46m"
const BG_WHITE = "[47m"

const VT_SAVE_CURSOR = "\x1b7"    // Save cursor and attributes
const VT_REST_CURSOR = "\x1b8"    // Restore cursor and attributes
const VT_CLEAR_SET = "\x1b[r"     // Clear scrollable window size
const VT_CLEAR_SCREEN = "\x1b[2J" // Clear screen
const VT_CLEAR_LINE = "\x1b[2K"   // Clear line
const VT_TERM_RESET = "\x1bc"     // Reset terminal completely

var ColorMap = map[string]string{
	"{n": MOD_CLEAR,

	"{l": FG_BLACK,
	"{r": FG_RED,
	"{g": FG_GREEN,
	"{y": FG_YELLOW,
	"{b": FG_BLUE,
	"{m": FG_MAGENTA,
	"{c": FG_CYAN,
	"{w": FG_WHITE,

	"{L": FG_B_BLACK,
	"{R": FG_B_RED,
	"{G": FG_B_GREEN,
	"{Y": FG_B_YELLOW,
	"{B": FG_B_BLUE,
	"{M": FG_B_MAGENTA,
	"{C": FG_B_CYAN,
	"{W": FG_B_WHITE,
}

var colorExpression = regexp.MustCompile("\\{[nlrgybmcwLRGYBMCW]")

func colorCode(message string) string {
	return ColorMap[message]
}

func Colorize(message string) string {
	return colorExpression.ReplaceAllStringFunc(message, colorCode) + MOD_CLEAR
}
