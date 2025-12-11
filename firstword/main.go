package piscine

func FirstWord(s string) string {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	start := i
	for i < len(s) && s[i] != ' ' {
		i++
	}

	return s[start:i] + "\n"
}
