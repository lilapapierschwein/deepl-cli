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
		var langs []string = utils.GetAvailableLanguages()

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
				`invalid language code on source ('%s')
Run 'deepl -L,--languages' to get a list of available languages and codes
`,
				from,
			),
		)
	}
	if !utils.IsInArray(to, opts) {
		return fmt.Errorf(
			fmt.Sprintf(
				`invalid language code on target ('%s')
Run 'deepl -L,--languages' to get a list of available languages and codes
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
		langOpts := utils.GetLanguageCodes()
		err := checkLangs(from, to, langOpts)
		if err != nil {
			return err
		}

		text := args[0]
		textEscaped := utils.UrlEscape(text)

		target := []string{deeplUrl, "#", from, "/", to, "/", textEscaped}
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
