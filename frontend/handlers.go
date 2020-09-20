/*** Arnulfo Jose Suarez Gaekel - 2020 - All rights reserved - Contact Email: ajoses@gmail.com */

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
)

var templates *template.Template

func renderHTTPError(r *http.Request, w http.ResponseWriter, err error, code int) {
	errMsg := fmt.Sprintf("%+v", err)

	w.WriteHeader(code)
	templates.ExecuteTemplate(w, "error", map[string]interface{}{
		"error":       errMsg,
		"status_code": code,
		"status":      http.StatusText(code)})
}

func (fe *frontendServer) peopleHandler(w http.ResponseWriter, r *http.Request) {

	people, err := fe.backendSvc.GetPeople()

	if err != nil {
		renderHTTPError(r, w, errors.Wrap(err, "could not retrieve people"), http.StatusInternalServerError)
	}

	if err = templates.ExecuteTemplate(w, "home", people); err != nil {
		fmt.Println(err)
	}
}

func (fe *frontendServer) frequencyHandler(w http.ResponseWriter, r *http.Request) {

	people, err := fe.backendSvc.GetFrequency()

	if err != nil {
		renderHTTPError(r, w, errors.Wrap(err, "could not retrieve frequency"), http.StatusInternalServerError)
	}

	if err = templates.ExecuteTemplate(w, "frequency", people); err != nil {
		fmt.Println(err)
	}
}

func (fe *frontendServer) duplicatesHandler(w http.ResponseWriter, r *http.Request) {

	people, err := fe.backendSvc.GetDuplicate()

	if err != nil {
		renderHTTPError(r, w, errors.Wrap(err, "could not retrieve frequency"), http.StatusInternalServerError)
	}

	if err = templates.ExecuteTemplate(w, "duplicates", people); err != nil {
		fmt.Println(err)
	}
}
