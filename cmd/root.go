package cmd

import (
	"deepl-cli/utils"
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const deeplUrl string = "https://deepl.com/en/translator"

var showLangs bool
var rootCMD = &cobra.Command{
	Use:   "deepl",
	Short: "interact with deepl via cli",
	Long:  "interact with deepl via cli",
	Run: func(cmd *cobra.Command, args []string) {
		langs := []string{
			"Arabic: ar",
			"Bulgarian: bg",
			"Chinese: zh",
			"Chinese (simplified): zh-hans",
			"Chinese (traditional): zh-hant",
			"Czech: cs",
			"Danish: da",
			"Dutch: nl",
			"English: en",
			"English (British): en-gb",
			"English (America): en-us",
			"Estonian: et",
			"Finnish: fi",
			"French: fr",
			"German: de",
			"Greek: el",
			"Hungarian: hu",
			"Indonesian: id",
			"Italian: it",
			"Japanese: ja",
			"Korean: ko",
			"Latvian: lv",
			"Lithuanian: lt",
			"Norwegian (bikmål): nb",
			"Polish: pl",
			"Portuguese: pt-pt",
			"Portuguese (Brazilian): pt-br",
			"Romanian: ro",
			"Russian: ru",
			"Slovak: sk",
			"Slovenian: sl",
			"Spanish: es",
			"Swedish: sv",
			"Turkish: tr",
			"Ukranian: uk",
		}
		if showLangs {
			for _, v := range langs {
				fmt.Println(v)
			}
		} else {
			browser.OpenURL(deeplUrl)
		}
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		// fmt.Fprintf(os.Stderr, "error: '%s'\n", err)
		os.Exit(1)
	}
}

func checkLangs(from string, to string, opts []string) error {
	if !utils.IsInArray(from, opts) {
		return fmt.Errorf(
			fmt.Sprintf(
				`'%s' not in language options
Run 'deepl -L' to see available language codes
`,
				from,
			),
		)
	}
	if !utils.IsInArray(to, opts) {
		return fmt.Errorf(
			fmt.Sprintf(
				`'%s' not in language options
Run 'deepl -L' to see available language codes
`,
				to,
			),
		)
	}

	return nil
}

var from string
var to string
var webCMD = &cobra.Command{
	Use:   "translate",
	Short: "translate string",
	Long:  "translate string",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		langOpts := []string{
			"ar",
			"bg",
			"zh",
			"zh-hans",
			"zh-hant",
			"cs",
			"da",
			"nl",
			"en",
			"en-gb",
			"en-us",
			"et",
			"fi",
			"fr",
			"de",
			"el",
			"hu",
			"id",
			"it",
			"ja",
			"ko",
			"lv",
			"lt",
			"nb",
			"pl",
			"pt-pt",
			"pt-br",
			"ro",
			"ru",
			"sk",
			"sl",
			"es",
			"sv",
			"tr",
			"uk",
		}
		err := checkLangs(from, to, langOpts)
		if err != nil {
			return err
		}

		target := []string{deeplUrl, "#", from, "/", to, "/", args[0]}
		var targetUrl string = strings.Join(target, "")
		browser.OpenURL(targetUrl)

		return nil
	},
}

// TODO: update usage strings
func init() {
	rootCMD.Flags().BoolVarP(&showLangs, "languages", "L", false, "Display available languages")
	webCMD.Flags().StringVarP(&from, "from", "F", "en", "Language to translate from")
	webCMD.Flags().StringVarP(&to, "to", "T", "de", "Language to translate to")
	rootCMD.AddCommand(webCMD)
}
