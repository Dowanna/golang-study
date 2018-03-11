package main

import (
	"flag"
	"fmt"
	"goLight/ch02/ex01/src/tempconv"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var tFlag = flag.String("t", "", "input the temp you want to convert with suffix c, f or k")

const (
	cel = "c"
	fah = "f"
	kel = "k"
)

func main() {
	flag.Parse()
	for k, v := range convertTemp(*tFlag) {
		printTemp(k, v)
	}
}

// remove suffixes, change to float, and convert
func convertTemp(t string) map[string]float64 {
	m := make(map[string]float64)

	if strings.HasSuffix(strings.ToLower(t), cel) {
		m[fah] = float64(tempconv.CToF(tempconv.Celsius(getOnlyNumbers(t))))
		m[kel] = float64(tempconv.CToK(tempconv.Celsius(getOnlyNumbers(t))))
	} else if strings.HasSuffix(strings.ToLower(t), fah) {
		m[cel] = float64(tempconv.FToC(tempconv.Fahrenheit(getOnlyNumbers(t))))
		m[kel] = float64(tempconv.FToK(tempconv.Fahrenheit(getOnlyNumbers(t))))
	} else if strings.HasSuffix(strings.ToLower(t), kel) {
		m[cel] = float64(tempconv.KToC(tempconv.Kelvin(getOnlyNumbers(t))))
		m[fah] = float64(tempconv.KToF(tempconv.Kelvin(getOnlyNumbers(t))))
	} else {
		os.Exit(1)
	}
	return m
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

func printTemp(s string, f float64) {
	fmt.Printf(" = %.2f%s\n", f, s)
}
