package http

import (
	"log"
	"net/http"

	"github.com/pquerna/ffjson/ffjson"
)

type Dto struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RenderJson(w http.ResponseWriter, v interface{}) {
	bs, err := ffjson.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}

func RenderSuccessDataJson(w http.ResponseWriter) {
	RenderJson(w, Dto{Msg: "Success", Data: "ok"})
}

func RenderFailedDataJson(w http.ResponseWriter, data interface{}) {

	log.Printf("send mail failed: %v\n", data)
	RenderJson(w, Dto{Msg: "Failed", Data: data})
}
