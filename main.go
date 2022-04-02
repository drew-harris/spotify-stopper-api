package main

import (
	handler "github.com/drew-harris/spotify-stopper-api/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Test)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
