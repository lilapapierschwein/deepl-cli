package cmd

import (
	"deepl-cli/utils"
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"strings"
)

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

var translateCMD = &cobra.Command{
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
