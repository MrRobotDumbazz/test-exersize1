package delivery

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Errors(w http.ResponseWriter, status int, msg string) {
	log.Println(msg)
	w.WriteHeader(status)
	t, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+"Error parsing file", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	errstruct := struct {
		StatusInt        int
		StatusText       string
		StatusIntAndText string
		Message          string
	}{
		status, http.StatusText(status), strconv.Itoa(status) + " " + http.StatusText(status), msg,
	}
	if err := t.Execute(w, errstruct); err != nil {
		log.Print(err)
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+"Error executing file", http.StatusInternalServerError)
		return
	}
}
