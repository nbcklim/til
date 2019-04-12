// simple web server
// run on commandline: `$ go run simple_server.go`

package main

import (
	"io"
	"net/http"
	"math/rand"
	"strconv"
)
func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/rand", randHandler)
	http.ListenAndServe(":8080", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<html>Hello Kevin! try <a href=\"/rand\">/rand</a>, would ya?</html>")
}

func randHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, strconv.Itoa(rand.Int()))
}
