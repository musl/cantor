package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

const AppName = "cantor backend"
const Version = "0.0.1"

type AppData struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
	Color   string `json:"color"`
}

func finish(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	io.WriteString(w, message)
}

func main() {
	app_data := AppData{
		AppName: AppName,
		Version: Version,
		Color:   "#ddeeff",
	}
	bindaddr := ":9000"
	docroot := "./build/"
	index := "index.html"

	log.Printf("%s v%s\n", AppName, Version)
	log.Printf("listening at %s\n", bindaddr)

	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			finish(w, http.StatusMethodNotAllowed, "")
			return
		}

		body, err := json.Marshal(app_data)
		if err != nil {
			finish(w, http.StatusInternalServerError, err.Error())
			return
		}

		header := w.Header()
		header["Content-Type"] = []string{"application/json"}
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req_path string
		pattern := regexp.MustCompile(`/$`)

		if r.URL.Path == "/" || pattern.MatchString(r.URL.Path) {
			req_path = path.Join(docroot, index)
		} else {
			req_path = filepath.Clean(r.URL.Path)
			req_path = path.Join(docroot, req_path)
		}

		_, err := os.Stat(req_path)
		if err != nil {
			finish(w, 404, "File Not Found")
			return
		}

		f, err := os.Open(req_path)
		if err != nil {
			finish(w, 500, fmt.Sprintf("Unable to open file: %s", req_path))
			return
		}

		i, err := os.Stat(req_path)
		if err != nil {
			finish(w, 500, fmt.Sprintf("Unable to stat file: %s", req_path))
			return
		}

		http.ServeContent(w, r, req_path, i.ModTime(), f)
		return
	})

	log.Fatal(http.ListenAndServe(bindaddr, nil))
}
