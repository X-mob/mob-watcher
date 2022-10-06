package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/X-mob/mob-watcher/config"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Run() {
	// api
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/mobs", Mobs)
	http.HandleFunc("/mob", MobItem)
	http.HandleFunc("/opensea/order", OpenSeaOrder)

	// get port
	const defaultListenPort = "8080"
	var port string
	if config.GlobalConfig.ListenPort == "" {
		port = defaultListenPort
	} else {
		port = config.GlobalConfig.ListenPort
	}

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
