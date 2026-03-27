package utils

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type UserTier int

const (
	UserTierNone UserTier = iota
	UserTierFree
	UserTierPro
)

var userTierName = map[UserTier]string{
	UserTierNone: "none",
	UserTierFree: "free",
	UserTierPro:  "pro",
}

func getUserTierNames() []string {
	var userTierNames []string

	for _, n := range userTierName {
		userTierNames = append(userTierNames, n)
	}

	return userTierNames
}

func (u UserTier) String() string {
	switch u {
	case 0:
		return "none"
	case 1:
		return "free"
	case 2:
		return "pro"
	}
	return fmt.Sprintf("UserTier(%q)", int(u))
}

var ErrInvalidUserTier = errors.New(color.RedString("invalid user tier"))

func ParseUserTier(t string) (UserTier, error) {
	switch t {
	case UserTierNone.String():
		return UserTierNone, nil
	case UserTierFree.String():
		return UserTierFree, nil
	case UserTierPro.String():
		return UserTierPro, nil
	}

	userTierNames := getUserTierNames()
	userTierNamesText := fmt.Sprintf("'%s'", userTierNames[0])

	for i := 1; i < len(userTierNames); i++ {
		userTierNamesText = fmt.Sprintf("%s,'%s'", userTierNamesText, userTierNames[i])
	}

	return UserTierNone, fmt.Errorf(
		"%s",
		fmt.Sprintf(
			`%s ('%s'). choose one of (%s).
`,
			"invalid user tier",
			t,
			userTierNamesText,
		),
	)
}

func getUserTierLimits() map[UserTier]int {
	userTierLimits := map[UserTier]int{
		UserTierNone: 1500,
		UserTierFree: 5000,
		UserTierPro:  20000,
	}

	return userTierLimits
}

func GetUserTiersAndLimits() []string {
	var userTierHelp []string
	userTierLimits := getUserTierLimits()

	userTierHelpText := "User Tiers & Limits"
	userTierHelpText = fmt.Sprintf("%s\n%s\n\ntier    limit*\n%s %s", userTierHelpText, strings.Repeat("=", len(userTierHelpText)), strings.Repeat("-", 7), strings.Repeat("-", 17))
	userTierHelp = append(userTierHelp, userTierHelpText)

	for t, l := range userTierLimits {
		fillChar := " "
		fillRange := 1
		tierName := t.String()

		if len(tierName) < 4 {
			fillRange += 4 - len(tierName)
		}

		numberFormatted := formatNumberSep(l)

		tierNameAndLimit := fmt.Sprintf("'%s'%s %s characters", tierName, strings.Repeat(fillChar, fillRange), numberFormatted)
		userTierHelp = append(userTierHelp, tierNameAndLimit)

	}
	userTierHelpNote := fmt.Sprintf(
		`
please note, that you must be logged into deepl in your web browser 
for elevated limits (%s & %s tier) to apply correctly. 

*character limitations for deepl web may be subject change`,
		"free",
		"pro",
	)

	userTierHelp = append(userTierHelp, userTierHelpNote)

	return userTierHelp
}
