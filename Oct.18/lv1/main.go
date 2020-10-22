package main

import (
	"fmt"
	"math"
)

func f(x float64, y float64) bool {
	const EPS float64 = 0.01
	return math.Pow(x*x+y*y-1, 3)-x*x-y*y*y < EPS
}

func main() {
	const BUFSIZE int = 100
	const HALFRANGE int = 40
	const PRECX float64 = 20
	const PRECY float64 = 10
	var (
		buf              [BUFSIZE][BUFSIZE]byte
		haveWrittenCharX [BUFSIZE]bool
		haveWrittenCharY [BUFSIZE]bool
	)
	for y := -HALFRANGE; y <= HALFRANGE; y++ {
		for x := -HALFRANGE; x <= HALFRANGE; x++ {
			if f(float64(x)/PRECX, float64(y)/PRECY) {
				buf[BUFSIZE/2+x][BUFSIZE/2+y] = '*'
				haveWrittenCharX[BUFSIZE/2+x] = true
				haveWrittenCharY[BUFSIZE/2+y] = true
			} else {
				buf[BUFSIZE/2+x][BUFSIZE/2+y] = ' '
			}
		}
	}
	for y := -HALFRANGE; y <= HALFRANGE; y++ {
		if !haveWrittenCharY[BUFSIZE/2+y] {
			continue
		}
		for x := -HALFRANGE; x <= HALFRANGE; x++ {
			if !haveWrittenCharX[BUFSIZE/2+x] {
				continue
			}
			fmt.Printf("%c", buf[BUFSIZE/2+x][BUFSIZE/2+y])
		}
		fmt.Printf("\n")
	}
}
