package cmd

import (
	"deepl-cli/utils"
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"os"
)

const deeplUrl string = "https://deepl.com/en/translator"

var displayVersion bool
var showLangs bool
var showTiers bool
var rootCMD = &cobra.Command{
	Use:   "deepl",
	Short: "interact with deepl via cli",
	Long:  "interact with deepl via cli",
	Run: func(cmd *cobra.Command, args []string) {
		var langs []string = utils.GetAvailableLanguages()
		var tiers []string = utils.GetUserTiersAndLimits()

		if displayVersion {
			deeplVersion := utils.GetVersion()
			fmt.Println(deeplVersion)
		} else if showLangs {
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

func init() {
	rootCMD.Flags().BoolVarP(&displayVersion, "version", "V", false, "Display version and exit.")
	rootCMD.Flags().BoolVarP(&showLangs, "languages", "L", false, "Display available languages.")
	rootCMD.Flags().BoolVarP(&showTiers, "user-tiers", "T", false, "Display user tiers and limits.")
	translateCMD.Flags().StringVarP(&from, "from", "F", "en", "Language to translate from")
	translateCMD.Flags().StringVarP(&to, "to", "T", "de", "Language to translate to")
	translateCMD.Flags().StringVarP(&userTier, "user-tier", "U", "free", "Language to translate to")
	rootCMD.AddCommand(translateCMD)
}
