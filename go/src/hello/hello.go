package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func get_version(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return string(contents)
}


func main() {
	log.Println("starting server...")
	actualVersion := get_version("https://trust1.my.intapp.com/time/mobile")
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(actualVersion))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

