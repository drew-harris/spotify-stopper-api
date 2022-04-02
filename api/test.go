package handler

import (
	"fmt"
	"net/http"
	"text/template"
)

type Page struct {
	Message string `json:"message"`
}

func Test(w http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("./api/hello.gohtml")
	if err != nil {
		panic(err)
	}

	p := Page{Message: "Hello World"}

	err = t.Execute(w, p)

	if err != nil {
		fmt.Println(err)
		return
	}
}
