package utils

import "fmt"

const prog string = "deepl"
const version string = "0.1.0"

func GetVersion() string {
	return fmt.Sprintf("%s %s", prog, version)
}
