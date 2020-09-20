/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package main

import (
"fmt"
	"github.com/ajoses/voxie-engineering-test/project/backend/pkg"
	"github.com/ajoses/voxie-engineering-test/project/backend/pkg/db"
	log "github.com/sirupsen/logrus"
"net/http"
"os"
)

const defaultBind = "0.0.0.0:8000"

func main() {

	s := pkg.NewBackend(dbClient)

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
