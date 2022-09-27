package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/X-mob/mob-watcher/config"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/opensea"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello From XMob Watcher!")
}

func OpenSeaOrder(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	tokenAddress := r.URL.Query().Get("asset_contract")
	tokenId := r.URL.Query().Get("token_id")

	if tokenAddress == "" {
		http.Error(w, "token address is null", http.StatusBadRequest)
		return
	}

	var opt opensea.RetrieveListingsOption
	opt.AssetContractAddress = tokenAddress
	opt.Limit = 1
	if tokenId != "" {
		opt.TokenIds = []string{tokenId}
	}

	orders := opensea.RetrieveListings(opt)

	json.NewEncoder(w).Encode(orders)
}

func MobItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "query address is null", http.StatusBadRequest)
		return
	}
	mob := db.GetMob(address)
	json.NewEncoder(w).Encode(mob)
}

func Mobs(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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
