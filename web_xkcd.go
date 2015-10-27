/* Web Server to render randomly fetched XKCD images on a TCP network address 
 * Author: Anjali Sharma (https://github.com/anjali-sharma) 
 */

// Default listening and serving TCP address: 'http://localhost:3000'

package main

import (
 	"net/http"
 	"github.com/rs/cors"
 	. "github.com/hemanth/go-xkcd"
 	)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
		})

	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", "Go-XKCD Web Server")
		w.Header().Set("Content-Type", "application/json")
		w.Write(Comic())
		})

	handler := c.Handler(mux)
	http.ListenAndServe(":3000", handler)
}
