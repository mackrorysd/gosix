package term

import (
	"regexp"
	"strings"
)

// Package term provides support for ANSI escape codes and related standards or
// conventions to improve the user experience in a terminal or emulator. Details
// found on [Wikipedia].
//
// [Wikipedia] https://en.wikipedia.org/wiki/ANSI_escape_code

const (
	Reset = "0"

	Bold         = "1"
	Faint        = "2"
	Italic       = "3"
	Underline    = "4"
	SlowBlink    = "5"
	RapidBlink   = "6"
	ReverseVideo = "7"
	Conceal      = "8"
	CrossedOut   = "9"

	// Font
	PrimaryFont      = "10"
	AlternativeFont0 = "11"
	AlternativeFont1 = "12"
	AlternativeFont2 = "13"
	AlternativeFont3 = "14"
	AlternativeFont4 = "15"
	AlternativeFont5 = "16"
	AlternativeFont6 = "17"
	AlternativeFont7 = "18"
	AlternativeFont8 = "19"
	Fraktur          = "20"

	DoublyUnderline     = "21"
	NormalIntensity     = "22"
	NotItalic           = "23"
	NotUnderline        = "24"
	NotBlink            = "25"
	ProportionalSpacing = "26"
	NotReverseVideo     = "27"
	Reveal              = "28"
	NotCrossedout       = "29"

	// Foreground colors
	BlackForeground   = "30"
	RedForeground     = "31"
	GreenForeground   = "32"
	YellowForeground  = "33"
	BlueForeground    = "34"
	MagentaForeground = "35"
	CyanForeground    = "36"
	WhiteForeground   = "37"
	RGBForeground     = "38"
	DefaultForeground = "39"

	//Background colors
	BlackBackground   = "40"
	RedBackground     = "41"
	GreenBackground   = "42"
	YellowBackground  = "43"
	BlueBackground    = "44"
	MagentaBackground = "45"
	CyanBackground    = "46"
	WhiteBackground   = "47"
	RGBBackground     = "48"
	DefaultBackground = "49"

	NotProportionalSpacing    = "50"
	Framed                    = "51"
	Encircled                 = "52"
	Overlined                 = "53"
	NotFramed                 = "54"
	NotOverlined              = "55"
	UnderlineColor            = "58"
	DefaultUnderlineColor     = "59"
	IdeogramUnderline         = "60"
	IdeogramDoubleUnderline   = "61"
	IdeogramOverline          = "62"
	IdeogramDoubleOverline    = "63"
	IdeogramStressMarking     = "64"
	NoIdeogramAttributes      = "65"
	Superscript               = "73"
	Subscript                 = "74"
	NotSuperscriptOrSubscript = "75"

	// Bright foreground colors
	BrightBlackForeground   = "90"
	BrightRedForeground     = "91"
	BrightGreenForeground   = "92"
	BrightYellowForeground  = "93"
	BrightBlueForeground    = "94"
	BrightMagentaForeground = "95"
	BrightCyanForeground    = "96"
	BrightWhiteForeground   = "97"

	// Bright background colors
	BrightBlackBackground   = "100"
	BrightRedBackground     = "101"
	BrightGreenBackground   = "102"
	BrightYellowBackground  = "103"
	BrightBlueBackground    = "104"
	BrightMagentaBackground = "105"
	BrightCyanBackground    = "106"
	BrightWhiteBackground   = "107"
)

// SGR implements the SelectGraphicRendition control function for selecting
// character attributes such as color and various styles and effects.
func SGR(attributes ...string) string {
	return "\033[" + strings.Join(attributes, ";") + "m"
}

// ControlCodeRegexp should match any of the supported control codes so that
// they can be detected or removed.
var ControlCodeRegexp = regexp.MustCompile("\033\\[[0-9;]*m")

// StripEscapeCodes removes anything that appears to be a control code.
func StripEscapeCodes(text string) string {
	return ControlCodeRegexp.ReplaceAllString(text, "")
}
