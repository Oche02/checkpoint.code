package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, r := range s {
		repeat := 1
		if r >= 'a' && r <= 'z' {
			repeat = int(r-'a') + 1
		}
		for i := 0; i < repeat; i++ {
			result += string(r)
		}
	}
	return result
}
