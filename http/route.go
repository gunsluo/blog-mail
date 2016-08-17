package http

import "net/http"

func configRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			get(w, r)
		case "POST":
			post(w, r)
		default:
		}

	})
}
