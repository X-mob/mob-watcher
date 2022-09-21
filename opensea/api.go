package opensea

import (
	"encoding/json"
	"fmt"
)

func RetrieveOffers(opt RetrieveOffersOption) RetrieveOffersResponse {
	router := "/orders/ethereum/seaport/offers"
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
	sendOpt.Router = router
	sendOpt.QueryParams = queryParams
	sendOpt.Method = Get

	res := SendRequestV2(sendOpt)

	// parse result
	var response RetrieveOffersResponse
	parseErr := json.Unmarshal(res, &response)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		panic(parseErr)
	}
	return response
}

func GetCollectionStats(slug string) CollectionStatsResponse {
	router := "/collection/" + slug + "/stats"
	var sendOpt SendRequestOption
	sendOpt.Router = router
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
