package asciitxt

import (
	"strings"
)

type Style int

const (
	Standard Style = iota + 1
)

type Config struct {
	Style Style
}

func New(txt string) string {
	return WithConfig(txt, Config{Style: Standard})
}

func WithConfig(txt string, config Config) string {
	if config.Style == 0 {
		config.Style = Standard
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

	var output string

	for k := 0; k < llen; k++ {
		for l := 0; l < len(arr); l++ {
			output += arr[l][k]
		}
		output += "\n"
	}

	return output
}

func getStyleLength(style Style) int {
	switch style {
	case Standard:
		return StandardLength
	default:
		return 0
	}
}

func getStyleLetter(style Style, letter string) []string {
	switch style {
	case Standard:
		return getStandardLetter(letter)
	default:
		panic("invalid style was given")
	}
}