package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		a   float32
		b   float32
		opt byte
		ans float32
	)
	fmt.Scanf("%f %c %f", &a, &opt, &b)
	switch opt {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		if b == 0 {
			fmt.Fprintf(os.Stderr, "Divided by 0!")
			return
		} else {
			ans = a / b
		}
	default:
		fmt.Fprintf(os.Stderr, "Invalid operation: %c", opt)
	}
	if ans == float32(int(ans)) {
		fmt.Printf("%d\n", int(ans))
	} else {
		fmt.Printf("%f\n", ans)
	}
	return
}
