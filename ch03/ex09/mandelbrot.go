package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	const (
		// 	xmin, xmax, ymin, ymax = -2, +2, -2, +2
		width, height = 1024, 1024
	)

	params := map[string]float64{
		"xmin": -2,
		"xmax": 2,
		"ymin": -2,
		"ymax": 2,
		"zoom": 1,
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		// クエリのパラメータをparamsに格納
		for qkey, qvalue := range queries {
			for param, _ := range params {
				if qkey == param {
					value, _ := strconv.ParseFloat(qvalue[0], 64)
					params[qkey] = value
				}
			}
		}

		zoom := params["zoom"]
		ymax := params["ymax"] / zoom
		ymin := params["ymin"] / zoom
		xmax := params["xmax"] / zoom
		xmin := params["xmin"] / zoom

		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < height; px++ {
				x := float64(px)/height*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}

		// png.Encode(os.Stdout, img)
		png.Encode(w, img)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 255 - contrast*n, 0, 255}
		}
	}
	return color.Black
}
