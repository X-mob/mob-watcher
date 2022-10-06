package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/opensea"
)

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
	mobInfo := db.GetMob(address)
	members := db.GetJoinDetailsByMobAddress(address)
	mob := Mob{
		Info:    mobInfo,
		Members: members,
	}
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
