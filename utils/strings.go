package utils

import "strings"

func NotExistReturnDefault(pre string, def string) string {
	if pre == "" {
		return def
	}

	return pre
}

func IncludeString(arr []string, s string) bool {
	for _, v := range arr {

		if strings.Contains(v, s) {
			return true
		}
	}

	return false
}
