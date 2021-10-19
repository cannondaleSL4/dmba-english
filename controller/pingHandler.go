package controller

import (
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ping")
	//t, _ := template.ParseFiles(INDEX)
	//t.Execute(w, nil)
}
