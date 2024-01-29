package main

func isValid(s string) bool {
	var stack []rune

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && top != '(') || (char == '}' && top != '{') || (char == ']' && top != '[') {
				return false
			}
		}
	}

	return len(stack) == 0
}
