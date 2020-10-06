package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type spaHandler struct {
	staticPath string
	indexPath string
}

// implementation got from https://github.com/gorilla/mux#serving-single-page-applications
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)
	log.Print(path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		log.Print(err)
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Print("serving")

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	router := mux.NewRouter()

	path, _ := os.Getwd()

	h := spaHandler{
		staticPath: fmt.Sprintf("%s/frontend/dist", path),
		indexPath:  "index.html",
	}

	router.PathPrefix("/").Handler(h)

	server := &http.Server{
		Addr:              "127.0.0.1:8181",
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
