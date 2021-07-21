package stacksandqueues

func BalancedParenthesis(A string) int {
	stack := make([]rune, 0)
	for _, v := range A {
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			return 0
		}
	}
	if len(stack) == 0 {
		return 1
	}
	return 0
}
