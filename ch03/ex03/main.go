package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle * 2), math.Cos(angle)
var min, max = 0.0, 0.0

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill; white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, color := corner(i+1, j+1)

			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Printf("</svg>")
}

func corner(i, j int) (float64, float64, string) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	max = math.Max(z, max)
	min = math.Min(z, min)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	// 凄まじく無駄なコードと自覚しつつも頭が回らずこうなってしまった・・・
	fmt.Fprintf(os.Stdout, "colorcode: %s\n", getColor(z))
	return sx, sy, getColor(z)
}

// 動作せず・・・どうしても完成しませんでした・・・
func getColor(z float64) string {
	// max値を255とする
	var ratio = 255 / max
	fmt.Fprintf(os.Stdout, "max as int64: %v\n", max)

	// 255との差分、0との差分を計算する
	var redFactor = int64(z * ratio)
	var blueFactor = int64(255 - (z * ratio))
	fmt.Fprintf(os.Stdout, "hexColor as int64: %v\n", redFactor)

	// hexに変える
	var redHex = []byte(strconv.FormatInt(redFactor, 16))
	var blueHex = []byte(strconv.FormatInt(blueFactor, 16))
	fmt.Fprintf(os.Stdout, "hexColor as byte: %.08b\n", redHex)

	return string(redHex[:]) + "00" + string(blueHex[:])
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
