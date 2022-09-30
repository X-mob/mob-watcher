package opensea

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/X-mob/mob-watcher/config"
)

func init() {
	networkSymbol = config.GlobalConfig.Network.String()
}

var networkSymbol string

func RetrieveListings(opt RetrieveListingsOption) RetrieveResponse {
	urlPath := "/orders/" + networkSymbol + "/seaport/listings"
	var queryParams []HttpUrlQueryParam
	if opt.Limit != 0 {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "limit", Value: fmt.Sprint(opt.Limit)})
	}
	if opt.AssetContractAddress != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "asset_contract_address", Value: opt.AssetContractAddress})
	}
	if len(opt.TokenIds) != 0 {
		for _, tokenId := range opt.TokenIds {
			queryParams = append(queryParams, HttpUrlQueryParam{Key: "token_ids", Value: tokenId})
		}
	}

	var sendOpt SendRequestOption
	sendOpt.Path = urlPath
	sendOpt.QueryParams = queryParams
	sendOpt.Method = Get

	res := SendRequestV2(sendOpt)

	// parse result
	var response RetrieveResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response
}

func RetrieveOffers(opt RetrieveOffersOption) RetrieveResponse {
	urlPath := "/orders/" + networkSymbol + "/seaport/offers"
	var queryParams []HttpUrlQueryParam
	if opt.Limit != 0 {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "limit", Value: fmt.Sprint(opt.Limit)})
	}
	if opt.OrderBy != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "order_by", Value: opt.OrderBy})
	} else {
		// default value
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "order_by", Value: "created_date"})
	}
	if opt.OrderDirection != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "order_direction", Value: opt.OrderDirection})
	} else {
		// default value
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "order_direction", Value: "desc"})
	}
	if opt.AssetContractAddress != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "asset_contract_address", Value: opt.AssetContractAddress})
	}
	if opt.Maker != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "maker", Value: opt.Maker})
	}
	if opt.Taker != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "taker", Value: opt.Taker})
	}
	if opt.ListedAfter != 0 {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "listed_after", Value: fmt.Sprint(opt.ListedAfter)})
	}
	if opt.ListedBefore != 0 {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "listed_before", Value: fmt.Sprint(opt.ListedBefore)})
	}
	if len(opt.TokenIds) != 0 {
		for _, tokenId := range opt.TokenIds {
			queryParams = append(queryParams, HttpUrlQueryParam{Key: "token_ids", Value: tokenId})
		}
	}

	var sendOpt SendRequestOption
	sendOpt.Path = urlPath
	sendOpt.QueryParams = queryParams
	sendOpt.Method = Get

	res := SendRequestV2(sendOpt)

	// parse result
	var response RetrieveResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response
}

func CreateListing(orderParameters OrderParameters, signature string) RetrieveResponse {
	urlPath := "/orders/" + networkSymbol + "/seaport/listings"
	s := CreateListingPayload{
		Parameters: orderParameters,
		Signature:  signature,
	}

	pl, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	payload := string(pl[:])
	fmt.Println(payload)

	opt := SendRequestOption{
		Path:    urlPath,
		Method:  Post,
		Payload: payload,
	}

	res := SendRequestV2(opt)

	// parse result
	var response RetrieveResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response
}

func CreateOffer(orderParameters OrderParameters, signature string) RetrieveResponse {
	urlPath := "/orders/" + networkSymbol + "/seaport/offers"

	op, err := json.Marshal(orderParameters)
	if err != nil {
		log.Fatalf("json.Marshal(orderParameters), err: %s", err)
	}

	params := url.Values{}
	params.Add("order_parameters", string(op))
	params.Add("signature", signature)
	payload := params.Encode()

	opt := SendRequestOption{
		Path:    urlPath,
		Method:  Post,
		Payload: payload,
	}

	res := SendRequestV2(opt)

	// parse result
	var response RetrieveResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response
}

func GetAssetContract(address string) AssetContract {
	urlPath := "/asset/" + address
	var sendOpt SendRequestOption
	sendOpt.Path = urlPath
	res := SendRequestV1(sendOpt)

	// parse result
	var response AssetContract
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response

}

func GetAssets(token string, tokenId string, includeOrders bool) GetAssetsResponse {
	urlPath := "/asset/" + token + "/" + tokenId + "/"
	var sendOpt SendRequestOption
	sendOpt.Path = urlPath

	if includeOrders == true {
		sendOpt.QueryParams = append(sendOpt.QueryParams, HttpUrlQueryParam{Key: "include_orders", Value: "true"})
	} else {
		sendOpt.QueryParams = append(sendOpt.QueryParams, HttpUrlQueryParam{Key: "include_orders", Value: "false"})
	}

	res := SendRequestV1(sendOpt)

	// parse result
	var response GetAssetsResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response

}

func GetCollectionStats(slug string) CollectionStatsResponse {
	urlPath := "/collection/" + slug + "/stats"
	var sendOpt SendRequestOption
	sendOpt.Path = urlPath
	res := SendRequestV1(sendOpt)

	// parse result
	var response CollectionStatsResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response
}

// usage
/*
	var opt lib.RetrieveOffersOption
	opt.Limit = 1
	res := lib.RetrieveOffers(opt)
	fmt.Printf("%+v\n", res)
*/
