package hoge

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

func Lissajous(out io.Writer) {
	var palette = []color.Color{color.White, color.Black}

	const (
		whiteIndex = 0 // パレットの最初の色 blackIndex = 1 // パレットの次の色
		blackIndex = 1
	)

	const (
		cycles  = 5 // 発振器 x が完了する周回の回数 res = 0.001 // 回転の分解能
		res     = 0.001
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
