// mandelbrot package by Erik Wrenholt 2015-11-28

package mandelbrot

import (
	"bytes"
)

const (
	max_iterations = 1000
	bailout = 16
)

func Calc(xx float64, yy float64) int {
	
	cr := yy - 0.5;
	ci := xx
	zi := 0.0
	zr := 0.0
	
	for i := 0; i < max_iterations; i++ {
		temp := zr * zi
		zr2 := zr * zr
		zi2 := zi * zi
		zr = zr2 - zi2 + cr
		zi = temp + temp + ci
		if (zi2 + zr2 > bailout) {
			return i;
		}
	}
	return max_iterations
}

func RenderSet(widthRadius int, heightRadius int) string {
	
	var buffer bytes.Buffer
	
	for y := -heightRadius; y < heightRadius; y++ {
	
		for x := -widthRadius; x < widthRadius; x++ {
		
			var xx float64
			var yy float64
			
			xx = float64(x) / float64(widthRadius)
			yy = float64(y) / float64(heightRadius)
			
			iterations := Calc(xx, yy)
			
			if iterations == max_iterations {
				buffer.WriteString("#")
			} else {
				buffer.WriteString(" ")
			}
			
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// Render a single row and send a string result back on the output channel
func renderRow(widthRadius int, yy float64, output chan string) {
	var buffer bytes.Buffer

	for x := -widthRadius; x < widthRadius; x++ {
		var xx float64 = float64(x) / float64(widthRadius)
		
		iterations := Calc(xx, yy)
		
		if iterations == max_iterations {
			buffer.WriteString("#")
		} else {
			buffer.WriteString(" ")
		}
	}
	output <- buffer.String()
}

func RenderSetParallel(widthRadius int, heightRadius int) string {
	var buffer bytes.Buffer

	// First make a 'slice' of channels to get result rows from
	chans := make([]chan string, heightRadius*2+1) //+1 for zero
	for i := range chans {
		chans[i] = make(chan string)
	}
	
	// Spawn a go routine to calculate each row
	i := 0
	for y := -heightRadius; y < heightRadius; y++ {
		i += 1
		yy := float64(y) / float64(heightRadius)
		go renderRow(widthRadius, yy, chans[i])
	}
	
	// Retrieve the results in each channel and generate the result
	i = 0
	for y := -heightRadius; y < heightRadius; y++ {
		i += 1
		buffer.WriteString(<- chans[i])
		buffer.WriteString("\n")
	}
	
	return buffer.String()
}

