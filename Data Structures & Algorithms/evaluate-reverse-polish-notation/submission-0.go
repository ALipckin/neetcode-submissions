func evalRPN(tokens []string) int {
	stack := []int{}
	var ops = map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}
	for _, s := range tokens {
		switch s{
			case "-", "+", "*", "/":
				f, _ := ops[s]
				r := stack[len(stack)-1]
				l := stack[len(stack)-2]
				res := f(l, r)
				stack = append(stack[:len(stack)-2], res)
			default:
				num, _ := strconv.Atoi(s)
				stack = append(stack, num)
		}
	}
	return int(stack[len(stack)-1])
}