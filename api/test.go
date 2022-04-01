package handler

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Whats up</h1")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
