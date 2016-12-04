package main

import (
    "fmt"
    "os"
    "strconv"
    "errors"
	"github.com/timestretch/mandelbrot_go/mandelbrot"
)

func getDimensions(args []string) (int, int, error) {

	if x, err := strconv.Atoi(args[1]); err == nil && x > 0 {
		if y, err := strconv.Atoi(args[2]); err == nil && y > 0 {
			return x, y, nil
		}
	}
	
	return 0, 0, errors.New("Please provide two positive integers!")
}

func main() {

	args := os.Args

	if len(args) == 3 {
	
		if x, y, err := getDimensions(args); err == nil {
			fmt.Printf(mandelbrot.RenderSetParallel(x/2.0, y/2.0))
		} else {
			fmt.Println(err)
		}
	
	} else {

		// Print help
		fmt.Printf("To render a 60 wide by 30 tall character set: \n")
		fmt.Printf("./render 60 30\n")
		
	}

}



