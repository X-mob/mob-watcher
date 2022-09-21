package task

import (
	"fmt"

	"github.com/X-mob/mob-watcher/convert"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/opensea"
)

func BuyNow(mobAddress string) {
	mobDb := db.GetMob(mobAddress)

	var opt opensea.RetrieveOffersOption
	// target the nft
	opt.AssetContractAddress = mobDb.Token.Hex()
	opt.TokenIds = []string{mobDb.TokenId.String()}
	// get cheapest one
	opt.Limit = 1
	opt.OrderBy = "eth_price"
	opt.OrderDirection = "asc"

	var order lib.BasicOrderParameters
	res := opensea.RetrieveOffers(opt)
	if len(res.Orders) == 0 {
		fmt.Printf("desired offer not found, token: %s,  tokenId: %s\n", opt.AssetContractAddress, opt.TokenIds[0])
		order = lib.BasicOrderParameters{}
	} else {
		order = convert.OpenseaToSeaportBasicOrderParameter(res.Orders[0])
	}

	lib.BuyNow(mobAddress, order)
}

func Buy() {

}
