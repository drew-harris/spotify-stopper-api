package handler

import (
	"net/http"
)

func Test(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/index.html")
}
