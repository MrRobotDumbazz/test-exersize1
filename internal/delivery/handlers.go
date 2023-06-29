package delivery

import "net/http"

func Handlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home_page)
	mux.HandleFunc("/blockchain", blockchain_page)
	return mux
}
