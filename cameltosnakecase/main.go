package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	isCamel := true
	for i, r := range s {
		if r < 'A' || (r > 'z' && r < 'a') || r > 'z' {
			isCamel = false
			break
		}
		if i > 0 {
			Prev := rune(s[i-1])
			if Prev >= 'A' && Prev <= 'Z' && r >= 'A' && r <= 'Z' {
				isCamel = false
				break
			}
		}
	}
	if !isCamel {
		return s
	}
	result := ""
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i != 0 {
				result += "-"
			}
		}
		result += string(r)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
