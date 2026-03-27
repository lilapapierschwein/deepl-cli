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

// const userTier = utils.UserTierFree

var showLangs bool
var showTiers bool
var rootCMD = &cobra.Command{
	Use:   "deepl",
	Short: "interact with deepl via cli",
	Long:  "interact with deepl via cli",
	Run: func(cmd *cobra.Command, args []string) {
		var langs []string = utils.GetAvailableLanguages()
		var tiers []string = utils.GetUserTiersAndLimits()

		if showLangs {
			for _, l := range langs {
				fmt.Println(l)
			}
		} else if showTiers {
			for _, t := range tiers {
				fmt.Println(t)
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
			"%s", fmt.Sprintf(
				`invalid language code on source ('%s').
run 'deepl -L,--languages' to get a list of available languages and corresponding codes.
`,
				from,
			),
		)
	}
	if !utils.IsInArray(to, opts) {
		return fmt.Errorf(
			"%s", fmt.Sprintf(
				`invalid language code on target ('%s').
run 'deepl -L,--languages' to get a list of available languages and corresponding codes.
`,
				to,
			),
		)
	}

	return nil
}

var from string
var to string
var userTier string

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

		usrTier, err := utils.ParseUserTier(userTier)
		if err != nil {
			return err
		}

		textFormatted := utils.FormatText(text, usrTier)

		target := []string{deeplUrl, "#", from, "/", to, "/", textFormatted}
		var targetUrl string = strings.Join(target, "")

		browser.OpenURL(targetUrl)

		return nil
	},
}

// TODO: update usage strings
func init() {
	rootCMD.Flags().BoolVarP(&showLangs, "languages", "L", false, "Display available languages")
	rootCMD.Flags().BoolVarP(&showTiers, "user-tiers", "T", false, "Display user tiers and limits")
	webCMD.Flags().StringVarP(&from, "from", "F", "en", "Language to translate from")
	webCMD.Flags().StringVarP(&to, "to", "T", "de", "Language to translate to")
	webCMD.Flags().StringVarP(&userTier, "user-tier", "U", "free", "Language to translate to")
	rootCMD.AddCommand(webCMD)
}
