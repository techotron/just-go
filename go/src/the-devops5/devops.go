package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type health struct {
	Health 	bool 	`json: health`
	Metrics metrics `json: metrics`
}

type metrics struct {
	RequestLatency 	float64	`json: requestLatency`
	DbLatency 		float64 `json: dbLatency`
	CacheLatency	float64 `json: cacheLatency`
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func get_json(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	foo2 := health{}
	get_json("http://localhost/health", &foo2)
	fmt.Println(foo2.Health)
}

