package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"websocketMicroservice/ws"
)

var Hub *ws.Hub


func main() {
	router := mux.NewRouter()

	Hub = ws.NewHub()
	go Hub.Run()

	router.HandleFunc("/ws",wsEndpoint).Methods("GET","OPTIONS")
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}

func wsEndpoint (w http.ResponseWriter, r *http.Request){
	fmt.Println("PASE A EJECUTAR EL ENDPOINT")
	r.Header.Add("Sec-Websocket-Version","13")
	ws.ServeWs(Hub, w, r)
}