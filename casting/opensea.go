package casting

import (
	"math/big"

	"github.com/X-mob/mob-watcher/lib"
	"github.com/X-mob/mob-watcher/opensea"
	"github.com/ethereum/go-ethereum/common"
)

func OpenSeaToSeaportOfferItem(offer opensea.OfferItem) lib.OfferItem {
	return lib.OfferItem{
		ItemType:             offer.ItemType,
		Token:                common.HexToAddress(offer.Token),
		IdentifierOrCriteria: StringToBigInt(offer.IdentifierOrCriteria, 10),
		StartAmount:          StringToBigInt(offer.StartAmount, 10),
		EndAmount:            StringToBigInt(offer.EndAmount, 10),
	}
}

func OpenSeaToSeaportConsiderationItem(consideration opensea.ConsiderationItem) lib.ConsiderationItem {
	return lib.ConsiderationItem{
		ItemType:             consideration.ItemType,
		Token:                common.HexToAddress(consideration.Token),
		IdentifierOrCriteria: StringToBigInt(consideration.IdentifierOrCriteria, 10),
		StartAmount:          StringToBigInt(consideration.StartAmount, 10),
		EndAmount:            StringToBigInt(consideration.EndAmount, 10),
		Recipient:            common.HexToAddress(consideration.Recipient),
	}
}

func OpenSeaToSeaportOffers(offers []opensea.OfferItem) []lib.OfferItem {
	var offer []lib.OfferItem
	for _, o := range offers {
		offer = append(offer, OpenSeaToSeaportOfferItem(o))
	}
	return offer
}

func OpenSeaToSeaportConsiderations(considerations []opensea.ConsiderationItem) []lib.ConsiderationItem {
	var consideration []lib.ConsiderationItem
	for _, c := range considerations {
		consideration = append(consideration, OpenSeaToSeaportConsiderationItem(c))
	}
	return consideration
}

func OpenSeaToSeaportOrder(order opensea.ProtocolData) lib.Order {
	offers := order.Parameters.Offer
	considerations := order.Parameters.Consideration

	var offer []lib.OfferItem = OpenSeaToSeaportOffers(offers)
	var consideration []lib.ConsiderationItem = OpenSeaToSeaportConsiderations(considerations)

	orderParameters := lib.OrderParameters{
		Offerer:       order.Parameters.Offerer,
		Zone:          order.Parameters.Zone,
		OrderType:     order.Parameters.OrderType,
		Offer:         offer,
		Consideration: consideration,
		StartTime:     StringToBigInt(order.Parameters.StartTime, 10),
		EndTime:       StringToBigInt(order.Parameters.EndTime, 10),
		ZoneHash:      StringToByte32(order.Parameters.ZoneHash),
		Salt:          StringToBigInt(order.Parameters.Salt, 10),
		ConduitKey:    StringToByte32(order.Parameters.ConduitKey),
	}

	contractOrder := lib.Order{
		Parameters: orderParameters,
		Signature:  []byte(order.Signature),
	}
	return contractOrder
}

func OpenSeaToSeaportBasicOrderParameter(order opensea.Order) lib.BasicOrderParameters {
	offer := order.ProtocolData.Parameters.Offer[0]
	consideration := order.ProtocolData.Parameters.Consideration[0]
	additionalReceipts := ConsiderationsToAdditionalReceipts(order.ProtocolData.Parameters.Consideration)
	return lib.BasicOrderParameters{
		ConsiderationToken:                common.HexToAddress(consideration.Token),
		ConsiderationIdentifier:           StringToBigInt(consideration.IdentifierOrCriteria, 10),
		ConsiderationAmount:               StringToBigInt(consideration.StartAmount, 10),
		Offerer:                           order.ProtocolData.Parameters.Offerer,
		Zone:                              order.ProtocolData.Parameters.Zone,
		OfferToken:                        common.HexToAddress(offer.Token),
		OfferIdentifier:                   StringToBigInt(offer.IdentifierOrCriteria, 10),
		OfferAmount:                       StringToBigInt(offer.StartAmount, 10),
		BasicOrderType:                    0, // todo: fix this
		StartTime:                         StringToBigInt(order.ProtocolData.Parameters.StartTime, 10),
		EndTime:                           StringToBigInt(order.ProtocolData.Parameters.EndTime, 10),
		ZoneHash:                          StringToByte32(order.ProtocolData.Parameters.ZoneHash),
		Salt:                              StringToBigInt(order.ProtocolData.Parameters.Salt, 10),
		OffererConduitKey:                 StringToByte32(order.ProtocolData.Parameters.ConduitKey),
		FulfillerConduitKey:               [32]byte{0},
		TotalOriginalAdditionalRecipients: big.NewInt(int64(len(additionalReceipts))),
		AdditionalRecipients:              additionalReceipts,
		Signature:                         []byte(order.ProtocolData.Signature),
	}
}

func ConsiderationsToAdditionalReceipts(_considerations []opensea.ConsiderationItem) []lib.AdditionalRecipient {
	considerations := OpenSeaToSeaportConsiderations(_considerations)
	var additionalReceipt []lib.AdditionalRecipient
	for _, c := range considerations {
		var receipt lib.AdditionalRecipient
		receipt.Recipient = c.Recipient
		receipt.Amount = c.StartAmount
		additionalReceipt = append(additionalReceipt, receipt)
	}
	return additionalReceipt
}
