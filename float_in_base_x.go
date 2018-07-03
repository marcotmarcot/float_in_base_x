package main

import (
	"flag"
	"fmt"
	"math"
)

var number = flag.Float64("number", 3.141592653589793, "Number to be converted to base.")

func main() {
	flag.Parse()
	for base := 2; base <= 36; base++ {
		fmt.Println(inBase(*number, base))
	}
}

func inBase(number float64, base int) string {
	s := ""
	fbase := float64(base)
	factor := math.Floor(math.Log(number) / math.Log(fbase))
	if factor < 0 {
		s += "0."
		for i := -1; i > int(factor); i-- {
			s += "0"
		}
	}
	power := math.Pow(fbase, factor)
	number /= power
	for i := 0; i < 16; i++ {
		d := byte(number)
		s += digit(d)
		number -= float64(d)
		number *= fbase
		if factor == 0 {
			s += "."
		}
		factor--
	}
	return s
}

func digit(number byte) string {
	if number < 10 {
		return string([]byte{number + 48})
	}
	return string([]byte{number + 55})
}
