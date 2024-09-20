package utils

func NotExistReturnDefault(pre string, def string) string {
	if pre == "" {
		return def
	}

	return pre
}
