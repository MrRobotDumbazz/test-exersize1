package delivery

import (
	"blockchain/internal/repository"
	"blockchain/models"
	"html/template"
	"log"
	"net/http"
)

func blockchain_page(w http.ResponseWriter, r *http.Request) {
	blockchainid := r.URL.Query().Get("id")
	if blockchainid == "" {
		Errors(w, http.StatusNotFound, "")
		return
	}
	log.Println(blockchainid)
	blockchains := models.Blockchains
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("templates/blockchainpage.html")
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = repository.GetJson("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1", &blockchains)
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		index := make(map[string]models.Blockchain)
		for _, item := range blockchains {
			index[item.ID] = item
		}
		err = tmpl.Execute(w, index[blockchainid])
		if err != nil {
			Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}
