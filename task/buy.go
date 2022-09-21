package task

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/X-mob/mob-watcher/casting"
	"github.com/X-mob/mob-watcher/db"
	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/opensea"
	"github.com/ethereum/go-ethereum/common"
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
		fmt.Printf("desired offer not found, token: %s,  tokenId: %s, mob: %s\n", opt.AssetContractAddress, opt.TokenIds[0], mobAddress)

		// add a mock order fot testing
		sig, _ := hex.DecodeString("95295124005c7ae9a687e481f43bd9768fa05da7ffe64331df2cc6cdb2df41e0195a9d777b2bbd047b978ee19cd096c13e6f6220ea47000a7457f54752dd11421c")

		s := "66fb18afe8bb49809ce454d7ce1decba3e6fe44c5eef2610b943fd5f37ec93e1"
		salt := new(big.Int)
		salt.SetString(s, 16)

		e := "5172014448931175958106549077934080"
		endTime := new(big.Int)
		endTime.SetString(e, 10)
		fmt.Printf(endTime.String())

		order = lib.BasicOrderParameters{
			Offerer:                           common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
			Zone:                              common.HexToAddress("0x0000000000000000000000000000000000000000"),
			BasicOrderType:                    0,
			OfferToken:                        common.HexToAddress("0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"),
			OfferIdentifier:                   big.NewInt(756),
			OfferAmount:                       big.NewInt(1),
			ConsiderationToken:                common.HexToAddress("0x0000000000000000000000000000000000000000"),
			ConsiderationIdentifier:           big.NewInt(0),
			ConsiderationAmount:               big.NewInt(2000000000000000000),
			ZoneHash:                          [32]byte{0},
			Salt:                              salt,
			StartTime:                         big.NewInt(0),
			EndTime:                           endTime,
			Signature:                         sig,
			OffererConduitKey:                 [32]byte{0},
			FulfillerConduitKey:               [32]byte{0},
			TotalOriginalAdditionalRecipients: big.NewInt(0),
			AdditionalRecipients:              []lib.AdditionalRecipient{},
		}
	} else {
		order = casting.OpenSeaToSeaportBasicOrderParameter(res.Orders[0])
	}

	lib.BuyNow(mobAddress, order)
}

func Buy() {

}
