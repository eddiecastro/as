/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package main

import (
	"fmt"
	"github.com/ajoses/salesloft-test/frontend/internal"

	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	listenAddr string
	listenPort string
)

type frontendServer struct {
	backendHost string
	backendSvc  *internal.BackendService
}

func main() {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	templates = template.Must(template.ParseGlob(filepath.Join(exPath, "templates", "*.html")))

	mustMapEnv(&listenPort, "PORT", "8080")
	mustMapEnv(&listenAddr, "SERVICE_ADDR", "0.0.0.0")

	svc := &frontendServer{}

	mustMapEnv(&svc.backendHost, "BACKEND_HOST", "http://127.0.0.1:8000")
	svc.backendSvc = internal.NewBackendService(svc.backendHost)

	r := mux.NewRouter()
	r.HandleFunc("/", svc.peopleHandler).Methods(http.MethodGet, http.MethodHead)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(exPath, "static")))))
	r.HandleFunc("/people", svc.peopleHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/frequency", svc.frequencyHandler).Methods(http.MethodGet, http.MethodHead)
	r.HandleFunc("/duplicates", svc.duplicatesHandler).Methods(http.MethodGet, http.MethodHead)

	var handler http.Handler = r

	fmt.Println("starting server on " + listenAddr + ":" + listenPort + " (Ctrl-C to quit)")
	http.ListenAndServe(listenAddr+":"+listenPort, handler)
}

func mustMapEnv(target *string, envKey string, defaultvalue string) {
	v := os.Getenv(envKey)
	if v == "" {
		*target = defaultvalue
	} else {
		*target = v
	}
}
