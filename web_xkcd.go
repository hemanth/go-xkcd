
/* Web Server to render randomly fetched XKCD images on a TCP network address 
 * Author: Anjali Sharma (https://github.com/anjali-sharma) 
 */

// Default listening and serving TCP address: 'http://localhost:3000'

package main

import (
	"net/http"
	. "github.com/hemanth/go-xkcd"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Go-XKCD Web Server")
	w.Header().Set("Content-Type", "application/json")
  w.Write(Comic())
}

func main() {
	http.HandleFunc("/", handler)
	//Change the listening TCP network address as required
	http.ListenAndServe(":3000", nil)
}
