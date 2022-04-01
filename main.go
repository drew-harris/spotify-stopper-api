package main

import (
	handler "github.com/drew-harris/spotify-stopper-api/api"
	"net/http"
)

func main() {
	http.HandleFunc("/test", handler.Test)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
