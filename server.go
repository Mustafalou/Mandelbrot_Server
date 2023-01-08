package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

func main() {
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/mandelbrot", handler)
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
func formhandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form.html")
}
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		t, _ := template.ParseFiles("form.html")
		t.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		realMin, _ := strconv.ParseFloat(r.Form.Get("realMin"), 64)
		realMax, _ := strconv.ParseFloat(r.Form.Get("realMax"), 64)
		imagMin, _ := strconv.ParseFloat(r.Form.Get("imagMin"), 64)
		imagMax, _ := strconv.ParseFloat(r.Form.Get("imagMax"), 64)
		iterations, _ := strconv.Atoi(r.Form.Get("iterations"))
		width, _ := strconv.Atoi(r.Form.Get("width"))
		height, _ := strconv.Atoi(r.Form.Get("height"))
		img := mandelbrot(realMin, realMax, imagMin, imagMax, width, height, iterations)
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, img)
	}
}

func mandelbrot(realMin, realMax, imagMin, imagMax float64, width, height, iterations int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	wg := sync.WaitGroup{}
	for py := 0; py < height; py++ {
		wg.Add(1)
		go func(py int) {
			defer wg.Done()
			y := float64(py)/float64(height)*(imagMax-imagMin) + imagMin
			for px := 0; px < width; px++ {
				x := float64(px)/float64(width)*(realMax-realMin) + realMin
				z := complex(x, y)
				img.Set(px, py, computeMandelbrotColor(z, iterations))
			}
		}(py)
	}
	wg.Wait()
	return img
}
func computeMandelbrotColor(z complex128, iterations int) color.Color {
	const contrast = 15

	var v complex128
	for n := uint8(0); n < uint8(iterations); n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
