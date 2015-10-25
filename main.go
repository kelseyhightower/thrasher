package main

import (
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := 0.0001
		for i := 0; i <= 1000000; i++ {
			x += math.Sqrt(x)
		}
		w.Write([]byte("OK!"))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
