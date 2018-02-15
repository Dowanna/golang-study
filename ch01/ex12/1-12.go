package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		convertedQueryMap := convertQueries(r.URL.Query())

		cycles := convertedQueryMap["cycles"]
		res := convertedQueryMap["res"]
		lissajous(w, cycles, res)
	})
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// こう書くとmap宣言しなくても使えるがnilのため代入時にクラッシュする。どうしてだろう
// func convertQueries(queryMap url.Values) (convertedMap map[string]float64) {
func convertQueries(queryMap url.Values) map[string]float64 {
	var convertedMap = map[string]float64{}
	for k, v := range queryMap {
		if vInt, err := strconv.Atoi(v[0]); err == nil {
			vAsFloat := float64(vInt)
			convertedMap[k] = vAsFloat
		} else {
			if vAsFloat, err := strconv.ParseFloat(v[0], 64); err == nil {
				convertedMap[k] = vAsFloat
			}
		}
	}
	return convertedMap
}

// handler はリクエストされた URL の Path 要素を返します。 func handler(w http.ResponseWriter, r *http.Request) {
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// --- lissajous --- //

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // パレットの最初の色 blackIndex = 1 // パレットの次の色
	blackIndex = 1
)

func lissajous(out io.Writer, cycles float64, res float64) {
	const (
		// cycles  = 5 // 発振器 x が完了する周回の回数 res = 0.001 // 回転の分解能
		// res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	// 画像キャンバスは [-size..+size] の範囲を扱う // アニメーションフレーム数
	// 10ms 単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発振器 y の相対周波数 anim := gif.GIF{LoopCount: nframes}
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: エンコードエラーを無視 }
}
