package delivery

import (
	"blockchain/internal/repository"
	"blockchain/models"
	"html/template"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errors(w, http.StatusNotFound, "")
		return
	}
	switch r.Method {
	case "GET":

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		blockchains := models.Blockchains
		err = repository.GetJson("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1", &blockchains)
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		err = tmpl.Execute(w, blockchains)
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}
