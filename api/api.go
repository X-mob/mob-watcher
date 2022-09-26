package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/X-mob/mob-watcher/config"
	"github.com/X-mob/mob-watcher/db"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello From XMob Watcher!")
}

func MobItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	address := r.URL.Query().Get("address")
	m := db.GetMob(address)
	mob := DBToAPIMob(m)

	json.NewEncoder(w).Encode(mob)
}

func Mobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var mobs []ShortMob

	address := db.GetAllMobAddress()
	for _, a := range address {
		m := db.GetMob(a)
		mobs = append(mobs, DBToAPIShortMob(m))
	}

	json.NewEncoder(w).Encode(mobs)
}

func Run() {
	// api
	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/mobs", Mobs)
	http.HandleFunc("/mob", MobItem)

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
