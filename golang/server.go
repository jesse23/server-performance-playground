package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// time.Sleep(2 * time.Second)

	done := make(chan int)

	// for i := 0; i < runtime.NumCPU(); i++ {
	for i := 0; i < 1; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					// do nothing
				}
			}
		}()
	}

	time.Sleep(time.Second * 1)
	close(done)

	io.WriteString(w, "Hello world!")
}

func main() {
	runtime.GOMAXPROCS(1)
	http.HandleFunc("/", hello)
	fmt.Print("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
