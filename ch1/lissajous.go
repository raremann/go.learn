// lissajous generates GIF animations
package main

// Namespace is the last part of part; references to color.White are really image.color.White, etc
import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// Composite literal; instatiate any of Go's composite types, here a slice
var palette = []color.Color{color.White, color.Black}

// named constants, values are set at compile time.
const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillations
		res     = 0.001 // angular resolution
		size    = 100
		nframes = 64 // number of animation frames
		delay   = 8  // delay between frames
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator

	// Composite literal that instatiates a struct of type gif.GIF
	// Sets LoopCount field of struct to nframe; all other fields are set to zero for there type
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		// 201 x 201 image
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// initialized to the 0th component of palette
		img := image.NewPaletted(rect, palette)
		// Set specific pixels of the current image to black
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5),
				size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
	//TODO: handle encoding errors
}
