package main

import (
	"log"
	"net/http"
	"text/template"

	_ "embed"
)

//go:embed index.html
var index []byte

const (
	TARGET_MODULE_HEADER string = "X-Target-Module"
	LISTEN_PORT          string = ":80"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page, err := template.New("index").Parse(string(index))
		if err != nil {
			log.Fatal(err)
		}

		page.Execute(w, r.Header.Get(TARGET_MODULE_HEADER))
	})

	log.Fatal(http.ListenAndServe(LISTEN_PORT, nil))
}
