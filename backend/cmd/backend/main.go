/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package main

import (
	"fmt"
	"github.com/ajoses/salesloft-test/backend/pkg"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

const defaultBind = "0.0.0.0:8000"

func main() {
	var apeSalesloftURLBase, apiSalesloftURLKey string

	mustMapEnv(&apeSalesloftURLBase, "SALESLOFT_URL_BASE", "https://api.salesloft.com")
	mustMapEnv(&apiSalesloftURLKey, "SALESLOFT_API_KEY", "unknown")

	log.Info(apiSalesloftURLKey)
	log.Info(apeSalesloftURLBase)

	s := pkg.NewBackend(apeSalesloftURLBase, apiSalesloftURLKey)

	fmt.Println("Listening in ", defaultBind)
	fmt.Println("Ctrl-C to exit...")

	log.Fatal(http.ListenAndServe(defaultBind, s.Router()))
}

func mustMapEnv(target *string, envKey string, defaultvalue string) {
	v := os.Getenv(envKey)
	if v == "" {
		*target = defaultvalue
	} else {
		*target = v
	}
}
