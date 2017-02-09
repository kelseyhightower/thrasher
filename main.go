package main

import (
	"log"
	"math"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	log.Println("Starting thrasher...")
	log.Printf("Number of CPUs: %d", runtime.NumCPU())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var wg sync.WaitGroup
		start := time.Now()

		for i := 0; i <= 10; i++ {
			wg.Add(1)
			go sqrt(&wg)
		}
		wg.Wait()

		elapsed := time.Since(start)
		log.Printf("Runtime took %s", elapsed)
		w.Write([]byte("OK!"))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}

func sqrt(wg *sync.WaitGroup) {
	defer wg.Done()
	x := 0.0001
	for i := 0; i <= 4000000000; i++ {
		x += math.Sqrt(x)
	}
}
