package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const AppName = "cantor backend"
const Version = "0.0.1"

type AppData struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	Color   string `json:"color"`
}

func main() {
	app_data := AppData{
		AppName: AppName,
		Version: Version,
		Color:   "#ddeeff",
	}
	bindaddr := ":9000"

	log.Printf("%s v%s\n", AppName, Version)
	log.Printf("listening at %s\n", bindaddr)

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, err := json.Marshal(app_data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		header := w.Header()
		header["Content-Type"] = []string{"application/json"}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(bindaddr, nil))
}
