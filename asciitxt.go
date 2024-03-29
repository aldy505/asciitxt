// ASCIITXT is a simple utility for creating ASCII texts.
// The text that can be created is limited to a certain input as the maintainers
// have to input all the text manually.
package asciitxt

import (
	"strings"
)

type Style int

const (
	StyleStandard Style = iota + 1
)

type Config struct {
	Style Style
}

// Generates a multi-line ASCII string with the default style.
// As simple as that.
func New(txt string) string {
	return WithConfig(txt, Config{Style: StyleStandard})
}

// Generates a multi-line ASCII string from a defined Config.
// The config itself could be empty. It'll look for its' default value.
// But if you don't intend to put any config whatsoever, please use New().
//
// Please note that this function will panic if the config.Style
// is invalid. Find the styles from the constants with Style prefix.
func WithConfig(txt string, config Config) string {
	if config.Style == 0 {
		config.Style = StyleStandard
	}

	if txt == "" {
		return ""
	}

	letters := strings.Split(txt, "")
	var arr [][]string
	llen := getStyleLength(config.Style)

	for i := 0; i < len(letters); i++ {
		var temp []string
		letter := getStyleLetter(config.Style, letters[i])
		for j := 0; j < len(letter); j++ {
			temp = append(temp, letter[j])
		}
		arr = append(arr, temp)
	}

	var output strings.Builder

	for k := 0; k < llen; k++ {
		for l := 0; l < len(arr); l++ {
			output.WriteString(arr[l][k])
		}
		output.WriteString("\n")
	}

	return output.String()
}

// Get a depth (vertical) length for a certain style.
func getStyleLength(style Style) int {
	switch style {
	case StyleStandard:
		return standardLength
	default:
		return 0
	}
}

// Get a letter by a certain style.
// It returns a panic if the style input is invalid.
func getStyleLetter(style Style, letter string) []string {
	switch style {
	case StyleStandard:
		return getStandardLetter(letter)
	default:
		panic("invalid style was given")
	}
}
