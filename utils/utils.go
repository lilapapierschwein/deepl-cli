package utils

import "strings"

func IsInArray(val string, arr []string) bool {
	var isInArr bool = false

	for _, v := range arr {
		if v == val {
			isInArr = true
			break
		}
		continue
	}

	return isInArr
}

func getUrlEscapes() map[string]string {
	urlEscapes := map[string]string{
		"/":  "\\/",
		"\n": "%0D%0A",
	}

	return urlEscapes
}

func UrlEscape(text string) string {
	urlEscapes := getUrlEscapes()
	stringEscaped := text

	for c, e := range urlEscapes {
		stringEscaped = strings.ReplaceAll(stringEscaped, c, e)
	}

	return stringEscaped
}
