package main

import (
	"log"
	"net/http"
	"time"
)

func get_time() string {
	t := time.Now().UTC().Format(time.ANSIC)
	return t
}

func main() {
	log.Println("starting server...")
	time := get_time()
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
