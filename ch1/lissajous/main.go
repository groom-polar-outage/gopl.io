// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.Black, color.RGBA{G: 0xff, A: 0xff}, color.RGBA{R: 0xff, A: 0xff}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	redIndex   = 2
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			user_res := request.URL.Query().Get("res")
			user_cycles := request.URL.Query().Get("cycles")
			user_size := request.URL.Query().Get("size")
			user_nframes := request.URL.Query().Get("nframes")
			user_delay := request.URL.Query().Get("delay")
			if len(user_cycles) == 0 {
				user_cycles = "5"
			}
			if len(user_res) == 0 {
				user_res = "0.001"
			}
			if len(user_size) == 0 {
				user_size = "100"
			}
			if len(user_nframes) == 0 {
				user_nframes = "100"
			}
			if len(user_delay) == 0 {
				user_delay = "8"
			}
			res, _ := strconv.ParseFloat(user_res, 64)

			cycles, _ := strconv.Atoi(user_cycles)

			size, _ := strconv.Atoi(user_size)

			nframes, _ := strconv.Atoi(user_nframes)

			delay, _ := strconv.Atoi(user_delay)

			//cycles, err := strconv.Atoi(request.URL.Query().Get("cycles"))
			//if err != nil {
			//	fmt.Fprintf(writer, "%v", err)
			//}
			//res, _ := strconv.ParseFloat(request.URL.Query().Get("res"), 64)
			//size, _ := strconv.Atoi(request.URL.Query().Get("size"))
			//nframes, _ := strconv.Atoi(request.URL.Query().Get("nframes"))
			//delay, _ := strconv.Atoi(request.URL.Query().Get("delay"))
			lissajous(writer, cycles, res, size, nframes, delay)
		})
		//handler := func(w http.ResponseWriter, r *http.Request) {
		//	lissajous(w)
		//}
		//http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8001", nil))
		return
	}
	//!+main
	//lissajous(os.Stdout)
}

//type Params struct {
//	cycles  int
//	res     float64
//	size    int
//	nframes int
//	delay   int
//}

//	func getParamValue(r http.Request) Params {
//		queryParams:= []string{"cycles","res","size","nframes","delay"}
//		for _,param :=range queryParams{
//			v := r.URL.Query().Get(param)
//			p :=Params{param: }
//		}
//	}
func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	//type params struct {
	//	cycles uint8 5   // number of complete x oscillator revolutions
	//	res = 0.001  // angular resolution
	//	size = 100   // image canvas covers [-size..+size]
	//	nframes = 100 // number of animation frames
	//	delay = 8    // delay between frames in 10ms units
	//}
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, int(2*size+1), int(2*size+1))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			var colorIndex uint8
			remainder := int(t/res) % 2
			if remainder == 0 {
				colorIndex = greenIndex
			} else {
				colorIndex = redIndex
			}
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
