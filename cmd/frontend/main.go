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
	log.Printf("request came to the server, req: %v", *r)
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

	log.Printf("got access to %s", h.staticPath)
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	port := os.Getenv("APP_PORT")
	log.Println("Starting frontend server - Smogger v1.0 by Kamil Głusiński")
	router := mux.NewRouter()

	path, _ := os.Getwd()

	h := spaHandler{
		staticPath: fmt.Sprintf("%s/static", path),
		indexPath:  "index.html",
	}

	router.PathPrefix("/").Handler(h)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}

	log.Printf("Listening on %s port", port)
	log.Fatal(server.ListenAndServe())
}
