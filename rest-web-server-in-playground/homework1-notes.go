// simple web server with two json rest end points
// - /hello returns hello world
// - /divide receive two numbers, divides them and returns result
// (note that you can run a web server with the go playground)
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

	
func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.ListenAndServe(":8090", nil)
}


//go run http-servers.go &

//curl localhost:8090/hello


// do it in the Go playground
package main

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

/* Output:
Hello, world!
Divide: 27.300000 / 4.170000 = 6.546763 
Divide: 2908.000000 / 43.000000 = 67.627907 

Program exited.
*/





Divide 
&http.Request{
	Method:"GET",
	URL:(*url.URL)(0xc0001921b0),
	Proto:"HTTP/1.1",
	ProtoMajor:1,
	ProtoMinor:1,
	Header:http.Header{"Accept-Encoding":[]string{"gzip"},
	"User-Agent":[]string{"Go-http-client/1.1"}
	}, 
Body:http.noBody{}, 
GetBody:(func() (io.ReadCloser, error))(nil),
ContentLength:0,
TransferEncoding:[]string(nil),
Close:false,
Host:"[::]:19522",
Form:url.Values(nil),
PostForm:url.Values(nil),
MultipartForm:(*multipart.Form)(nil),
Trailer:http.Header(nil),
RemoteAddr:"127.0.0.1:62241",
RequestURI:"/divide?dividend=25?divisor=5",
TLS:(*tls.ConnectionState)(nil),
Cancel:(<-chan struct {})(nil),
Response:(*http.Response)(nil)
ctx:(*context.cancelCtx)(0xc0001a6050)
}