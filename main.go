package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	_ "embed"
)

//go:embed index.html
var index []byte

const (
	LISTEN_PORT string = ":80"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, err := template.New("index").Parse(string(index))
		if err != nil {
			log.Fatal(err)
		}

		page.Execute(w, strings.TrimLeft(r.URL.Path, "/"))
	})

	log.Fatal(http.ListenAndServe(LISTEN_PORT, nil))
}
