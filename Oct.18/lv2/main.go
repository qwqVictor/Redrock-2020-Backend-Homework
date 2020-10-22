package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// print answer
// if it's an integer, print it without dot.
func printAns(ans float32) {
	if ans == float32(int(ans)) {
		fmt.Printf("%d\n", int(ans))
	}else {
		fmt.Printf("%f\n", ans)
	}
}

// input data
func input() (float32, float32, byte, error) {
	var (
		a   float32
		b   float32
		opt byte
	)
	_, err := fmt.Scanf("%f %c %f", &a, &opt, &b)
	return a, b, opt, err
}

// run calculation
func calc(a float32, b float32, opt byte) (float32, error) {
	var (
		ans float32
		err error
	)
	switch opt {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		if b == 0 {
			err = errors.New("divided by 0")
			return 0, err
		} else {
			ans = a / b
		}
	default:
		err = errors.New("invalid operation:" + string(opt))
		return 0, err
	}
	return ans, nil
}

func main() {
	for {
		a, b, opt, err := input()
		if (err != nil) {
			return
		}
		ans, err := calc(a, b, opt)
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error());
			syscall.Exit(1)
		}
		printAns(ans)
	}
	return
}
