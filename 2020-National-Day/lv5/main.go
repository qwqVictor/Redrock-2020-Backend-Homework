package main

import "fmt"

func isPalindromic(str string) bool {
	length := len(str)
	for i := 0; i < (length >> 1); i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

func main() {
	str := ""
	startends := [2]int{0, 0}
	fmt.Scanln(&str)
	length := len(str)
	for i := 0; i < length; i++ {
		for j := length; j > i; j-- {
			if isPalindromic(str[i:j]) {
				if j-i > startends[1]-startends[0] {
					startends[0] = i
					startends[1] = j
				}
			}
		}
	}
	fmt.Println(startends[1] - startends[0])
	fmt.Println(str[startends[0]:startends[1]])
}
