package http

import "net/http"

func setAllowOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	return
}

func configRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		setAllowOrigin(w, r)

		switch r.Method {
		case "GET":
			get(w, r)
		case "POST":
			post(w, r)
		default:
		}

	})
}
