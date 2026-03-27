package utils

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func getMaxTextLen(userTier UserTier) int {
	tierLimits := getUserTierLimits()
	maxTextLen := tierLimits[userTier]

	return maxTextLen
}

func limitTextLength(text string, maxTextLen int, userTierName string) string {
	getLimitWarning := func(text string, newTextLen int) string {
		textLen := formatNumberSep(len(text))
		actualTextLen := formatNumberSep(newTextLen)

		warningText := fmt.Sprintf("Warning: text was cut from %s to %s characters to fit %s-tier limit!", textLen, actualTextLen, userTierName)
		return color.YellowString(warningText)
	}

	textLimited := text

	if len(text) > maxTextLen {
		textLimited = text[:maxTextLen-1]
		println(getLimitWarning(text, len(textLimited)))
	}
	return textLimited
}

func getUrlEscapes() map[string]string {
	urlEscapes := map[string]string{
		"/":  "\\/",
		"\n": "%0D%0A",
	}

	return urlEscapes
}

func urlEscape(text string) string {
	urlEscapes := getUrlEscapes()
	stringEscaped := text

	for c, e := range urlEscapes {
		stringEscaped = strings.ReplaceAll(stringEscaped, c, e)
	}

	return stringEscaped
}

func FormatText(text string, userTier UserTier) string {
	maxTextLen := getMaxTextLen(userTier)

	formattedText := limitTextLength(text, maxTextLen, userTier.String())
	formattedText = urlEscape(formattedText)

	return formattedText
}
