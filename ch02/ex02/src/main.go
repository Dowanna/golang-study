package main

import (
	"fmt"
	"goLight/ch02/ex02/src/tempconv"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		t := getOnlyNumbers(arg)

		// length
		if strings.HasSuffix(arg, "m") || strings.HasSuffix(arg, "ft") {
			m := tempconv.Meter(t)
			f := tempconv.Feet(t)
			fmt.Printf("%s = %s, %s = %s\n", m, tempconv.MToFt(m), f, tempconv.FtToM(f))
			return
		}

		// weight
		if strings.HasSuffix(arg, "kg") || strings.HasSuffix(arg, "lb") {
			kg := tempconv.Kilogram(t)
			lb := tempconv.Pound(t)
			fmt.Printf("%s = %s, %s = %s\n", kg, tempconv.LbToKg(lb), lb, tempconv.KgToLb(kg))
			return
		}

		// temp
		if strings.HasSuffix(arg, "c") || strings.HasSuffix(arg, "f") {
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
			return
		}
	}
}

// remove alphabets, and return float
func getOnlyNumbers(s string) float64 {
	r, _ := regexp.Compile("[^0-9\\.]+")

	convs, e := strconv.ParseFloat(r.ReplaceAllString(s, ""), 64)
	if e != nil {
		fmt.Printf("%v", e)
	}

	return float64(convs)
}
