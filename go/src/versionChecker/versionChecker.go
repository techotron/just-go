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
	envUrl := os.Getenv("URL")
	actualVersion := get_version(envUrl)
	log.Println("Actual version is " + actualVersion)
}

