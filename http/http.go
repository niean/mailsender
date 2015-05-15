package http

import (
	"encoding/json"
	"github.com/niean/mailsender/g"
	"log"
	"net/http"
)

func Start() {
	go startHttpServer()
}

// add url mapping config
func addRoutes() {
	configCommonRoutes()
	configMailSenderApiRoutes()
	configProcHttpRoutes()
}

func startHttpServer() {
	if !g.GetConfig().Http.Enable {
		return
	}

	addr := g.GetConfig().Http.Listen
	if addr == "" {
		return
	}

	addRoutes()
	s := &http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 30,
	}

	log.Println("http.startHttpServer ok, listening", addr)
	log.Fatalln(s.ListenAndServe())
}

// interfaces
type Dto struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RenderDataJson(w http.ResponseWriter, data interface{}) {
	renderJson(w, Dto{Msg: "success", Data: data})
}

func renderJson(w http.ResponseWriter, v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bs)
}
