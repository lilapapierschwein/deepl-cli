package utils

import (
	"fmt"
	"slices"
	"strings"
)

func formatAvailableLanguages(langsAvail []string) []string {
	getSpecialChars := func() []string {
		specialChars := []string{"å"}
		return specialChars
	}

	getLangCodesAndTexts := func(langsAvail []string) ([]string, []string) {
		var langCodes []string
		var langsFull []string

		for _, l := range langsAvail {
			langCode, langFull, _ := strings.Cut(l, ":")

			langCodes = append(langCodes, strings.Trim(langCode, " "))
			langsFull = append(langsFull, strings.Trim(langFull, " "))
		}

		return langCodes, langsFull
	}
	getMaxLengths := func(langCodes []string, langsFull []string) (int, int) {
		maxLenCode, maxLenFull := -1, -1

		for i := range langCodes {
			lenCode, lenFull := len(langCodes[i]), len(langsFull[i])

			if lenCode > maxLenCode {
				maxLenCode = lenCode
			}
			if lenFull > maxLenFull {
				maxLenFull = lenFull
			}
		}

		return maxLenCode, maxLenFull
	}
	createHeader := func(maxLenCode int, maxLenFull int) string {
		langsHeadText := fmt.Sprintf(
			"%s%s%s\n%s %s",
			"Language",
			strings.Repeat(" ", maxLenFull-len("Language")+1),
			"Code",
			strings.Repeat("-", maxLenFull),
			strings.Repeat("-", maxLenCode),
		)

		return langsHeadText
	}
	getLangsFormatted := func(langCodes []string, langsFull []string, maxLenFull int) []string {
		specialCharacters := getSpecialChars()
		getFillRange := func(maxLenFull int, langFull string, specialChars []string) int {
			fillRange := maxLenFull - len(langFull) + 1

			// account for non counting chars like Norwegian 'å'
			for _, c := range specialChars {
				if strings.Contains(langFull, c) {
					fillRange += 1
				}
			}

			return fillRange
		}

		var langsFormatted []string

		for i := range len(langCodes) {
			langFull := langsFull[i]
			langCodeQuoted := fmt.Sprintf("'%s'", langCodes[i])

			if len(langFull) < maxLenFull {
				fillRange := getFillRange(maxLenFull, langFull, specialCharacters)

				langFull = fmt.Sprintf("%s%s%s", langFull, strings.Repeat(" ", fillRange), langCodeQuoted)
			} else {
				langFull = fmt.Sprintf("%s%s%s", langFull, " ", langCodeQuoted)
			}

			langsFormatted = append(langsFormatted, langFull)
		}

		return langsFormatted
	}

	langCodes, langsFull := getLangCodesAndTexts(langsAvail)
	maxCodeLen, maxLangLen := getMaxLengths(langCodes, langsFull)

	langsFormatted := getLangsFormatted(langCodes, langsFull, maxLangLen)

	langsHeader := createHeader(maxCodeLen, maxLangLen)
	langsAvail = slices.Insert(langsFormatted, 0, langsHeader)

	return langsAvail
}

func getLanguages() []string {
	langsAvail := []string{
		"ar: Arabic",
		"bg: Bulgarian",
		"zh: Chinese",
		"zh-hans: Chinese (simplified)",
		"zh-hant: Chinese (traditional)",
		"cs: Czech",
		"da: Danish",
		"nl: Dutch",
		"en: English",
		"en-gb: English (British)",
		"en-us: English (America)",
		"et: Estonian",
		"fi: Finnish",
		"fr: French",
		"de: German",
		"el: Greek",
		"hu: Hungarian",
		"id: Indonesian",
		"it: Italian",
		"ja: Japanese",
		"ko: Korean",
		"lv: Latvian",
		"lt: Lithuanian",
		"nb: Norwegian (bikmål)",
		"pl: Polish",
		"pt-pt: Portuguese",
		"pt-br: Portuguese (Brazilian)",
		"ro: Romanian",
		"ru: Russian",
		"sk: Slovak",
		"sl: Slovenian",
		"es: Spanish",
		"sv: Swedish",
		"tr: Turkish",
		"rk: Ukranian",
	}

	return langsAvail
}

func GetAvailableLanguages() []string {
	langsAvail := getLanguages()
	return formatAvailableLanguages(langsAvail)
}

func GetLanguageCodes() []string {
	langsAvail := getLanguages()

	var langCodes []string
	for _, l := range langsAvail {
		langCode, _, _ := strings.Cut(l, ":")
		langCodes = append(langCodes, strings.Trim(langCode, " "))
	}

	return langCodes
}
