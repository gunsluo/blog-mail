package http

import (
	"log"
	"net/http"

	"github.com/gunsluo/blog-mail/common"
)

func Start() {

	configRoutes()

	addr := common.Config().Http.Listen
	if addr == "" {
		return
	}
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}
	log.Println("http listening", addr)
	log.Fatalln(s.ListenAndServe())

}
