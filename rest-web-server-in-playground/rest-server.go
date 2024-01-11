package main

/*
Homework:
Make a simple web server in Go with two JSON "REST" endpoints:
  + /hello -> returns hello world
  + /divide -> receives 2 numbers, divides them and returns the result
  + (note that you can run a web server with the Go playground)

Output:
Hello, world!
Divide: 27.300000 / 4.170000 = 6.546763
Divide: 2908.000000 / 43.000000 = 67.627907

Program exited.

References:
https://pkg.go.dev/net@go1.21.5
https://pkg.go.dev/net/http@go1.21.5#Serve
https://pkg.go.dev/net/url@go1.21.5#ParseQuery
https://gobyexample.com/url-parsing
*/

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
)

func main() {
	// Establish handler functions
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	divideHandler := func(w http.ResponseWriter, req *http.Request) {
		u, err := url.Parse(req.RequestURI)
		if err != nil {
			fmt.Fprintf(w, "Divide: %v \n", err)
			return
		}
		m, err := url.ParseQuery(u.RawQuery)
		if err != nil {
			fmt.Fprintf(w, "Divide: %v \n", err)
			return
		}
		dividend, err := strconv.ParseFloat(m["dividend"][0], 64)
		if err != nil {
			fmt.Fprintf(w, "Divide: %v \n", err)
			return
		}
		divisor, err := strconv.ParseFloat(m["divisor"][0], 64)
		if err != nil {
			fmt.Fprintf(w, "Divide: %v \n", err)
			return
		}
		if divisor == 0.0 {
			fmt.Fprintf(w, "Divide: %f / %f: divide by zero, not allowed \n", dividend, divisor)
			return
		}
		quotient := dividend / divisor
		fmt.Fprintf(w, "Divide: %f / %f = %f \n", dividend, divisor, quotient)
	}

	// setup and start the http server in a go routine
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/divide", divideHandler)
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("Listen", err)
	}
	go func() {
		log.Fatal("Serve", http.Serve(ln, nil))
	}()
	runtime.Gosched() // get that go-routine server running

	// issue multiple requests to the server at In.Addr()
	urls := []string{
		"http://%s/hello",
		"http://%s/divide?dividend=27.3&divisor=4.17",
		"http://%s/divide?dividend=2908&divisor=43",
	}
	var resp *http.Response
	for _, url := range urls {
		resp, err = http.Get(fmt.Sprintf(url, ln.Addr()))
		if err != nil {
			log.Fatal("Get", err)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal("io.Copy", err)
		}
	}
	resp.Body.Close()
}
