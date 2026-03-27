package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

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

func formatNumberSep(num int) string {
	p := message.NewPrinter(language.English)
	numberFormatted := p.Sprintf("%d", num)

	return numberFormatted
}
