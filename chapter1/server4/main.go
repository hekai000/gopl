// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	lconf := lconfig{
		cycles: 5,
		res:    0.001,
		freq:   rand.Float64() * 3.0,
		size:   100,
		frames: 64,
		delay:  8,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		confs := r.URL.Query()
		for i, c := range confs {
			switch i {
			case "cycles":
				lconf.cycles, _ = strconv.ParseFloat(c[0], 64)
			case "freq":
				lconf.freq, _ = strconv.ParseFloat(c[0], 64)
			case "res":
				lconf.res, _ = strconv.ParseFloat(c[0], 64)
			case "size":
				lconf.size, _ = strconv.Atoi(c[0])
			case "frames":
				lconf.frames, _ = strconv.Atoi(c[0])
			case "delay":
				lconf.delay, _ = strconv.Atoi(c[0])
			}
		}

		lissajous(w, lconf)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type lconfig struct {
	cycles float64
	res    float64
	freq   float64
	size   int
	frames int
	delay  int
}

func lissajous(out io.Writer, set lconfig) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	//freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: set.frames}
	phase := 0.0 // phase difference
	for i := 0; i < set.frames; i++ {
		rect := image.Rect(0, 0, 2*set.size+1, 2*set.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < set.cycles*2*math.Pi; t += set.res {
			x := math.Sin(t)
			y := math.Sin(t*set.freq + phase)
			img.SetColorIndex(set.size+int(x*float64(set.size)+0.5), set.size+int(y*float64(set.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, set.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-handler
