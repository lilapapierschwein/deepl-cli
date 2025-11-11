package utils

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
