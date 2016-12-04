Golang Mandelbrot Example
-------------------------

Mandelbrot set rendering in golang with parallel and non parallel version.

Run unit tests:

	go test

Results:

	PASS
	ok  	github.com/timestretch/mandelbrot_go/mandelbrot	0.041s

Run benchmark:

	go test -bench=".*"

Results:

	BenchmarkMandelbrotRender-4           	       1	32014223686 ns/op
	BenchmarkMandelbrotRenderParallel-4   	       1	10662301401 ns/op
	PASS
	ok  	github.com/timestretch/mandelbrot_go/mandelbrot	42.729s

Render:

	cd render
	go build
	./render 60 30

Results:

	                              #                             
	                              #                             
	                              #                             
	                             ###                            
	                            #####                           
	                         ###########                        
	                        #############                       
	                       ###############                      
	                       ###############                      
	                        #############                       
	                         ###########                        
	                            #####                           
	                       ###############                      
	                 # ####################### #                
	                #############################               
	            #####################################           
	              #################################             
	           #######################################          
	           #######################################          
	           #######################################          
	      #################################################     
	    #####################################################   
	      ###  #######################################  ###     
	            #####################################           
	            #####################################           
	             ###################################            
	               ###############################              
	             ### ############# ############# ###            
	                    # ####         #### #                   
	                                                            

MIT License
-------

Copyright (c) 2015 Erik Wrenholt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

