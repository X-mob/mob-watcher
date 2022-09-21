package lib

import (
	//"bytes"
	//"encoding/gob"
	//"encoding/json"
	"encoding/json"
	"fmt"

	//"hash"
	"io/ioutil"
	"net/http"

	"github.com/X-mob/mob-watcher/config"
	"github.com/ethereum/go-ethereum/common"
)

type ReqMethod int64

const (
	Get ReqMethod = iota
	Post
	Option
)

func (s ReqMethod) String() string {
	switch s {
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Option:
		return "OPTION"
	}
	return "unknown"
}

type HttpUrlQueryParam struct {
	Key   string
	Value string
}
type SendRequestOption struct {
	Router      string // start with slash
	Method      ReqMethod
	QueryParams []HttpUrlQueryParam
}

func SendRequest(opt SendRequestOption, baseUrl string) []byte {
	url := baseUrl + opt.Router
	req, _ := http.NewRequest(opt.Method.String(), url, nil)

	q := req.URL.Query()
	for _, p := range opt.QueryParams {
		q.Add(p.Key, p.Value)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", config.GlobalConfig.OpenSeaApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("req failed, statusCode %d, url %s", res.StatusCode, res.Request.URL))
	}

	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func SendRequestV1(opt SendRequestOption) []byte {
	return SendRequest(opt, config.OpenseaApiV1BaseUrl)
}

func SendRequestV2(opt SendRequestOption) []byte {
	return SendRequest(opt, config.OpenseaApiV2BaseUrl)
}

type ProtocolDataParameterOfferItem struct {
	ItemType             uint8  `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
}

type ProtocolDataParameterConsiderationItem struct {
	ItemType             uint8  `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
	Recipient            string `json:"recipient"`
}

type ProtocolDataParameters struct {
	Offerer                         common.Address                           `json:"offerer"`
	Zone                            common.Address                           `json:"zone"`
	Offer                           []ProtocolDataParameterOfferItem         `json:"offer"`
	Consideration                   []ProtocolDataParameterConsiderationItem `json:"consideration"`
	OrderType                       uint8                                    `json:"orderType"`
	StartTime                       string                                   `json:"startTime"`
	EndTime                         string                                   `json:"endTime"`
	ZoneHash                        string                                   `json:"zoneHash"`
	Salt                            string                                   `json:"salt"`
	ConduitKey                      string                                   `json:"conduitKey"`
	TotalOriginalConsiderationItems int                                      `json:"totalOriginalConsiderationItems"`
	Counter                         int                                      `json:"counter"`
}

type ProtocolData struct {
	Parameters ProtocolDataParameters `json:"parameters"`
	Signature  string                 `json:"signature"`
}

type AssetCollection struct {
	Slug string `json:"slug"`
}
type Asset struct {
	Id                   string          `json:"id"`
	ImageUrl             string          `json:"image_url"`
	ImagePreviewUrl      string          `json:"image_preview_url"`
	ImageThumbnailUrl    string          `json:"image_thumbnail_url"`
	ImageOriginalUrl     string          `json:"image_original_url"`
	AnimationUrl         string          `json:"animation_url"`
	AnimationOriginalUrl string          `json:"animation_original_url"`
	Name                 string          `json:"name"`
	Description          string          `json:"description"`
	Collection           AssetCollection `json:"collection"`
}

type AssetBundle struct {
	assets []Asset
}
type Offer struct {
	OrderHash        string       `json:"order_hash"`
	ProtocolData     ProtocolData `json:"protocol_data"`
	ProtocolAddress  string       `json:"protocol_address"`
	Side             string       `json:"side"`
	OrderType        string       `json:"order_type"`
	Cancelled        bool         `json:"cancelled"`
	Finalized        bool         `json:"finalized"`
	MarkedInvalid    bool         `json:"marked_invalid"`
	MakerAssetBundle AssetBundle  `json:"maker_asset_bundle"`
}

type RetrieveOffersResponse struct {
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Orders   []Offer `json:"orders"`
}

type RetrieveOffersOption struct {
	Limit                int
	OrderBy              string
	OrderDirection       string
	AssetContractAddress string
	TokenIds             []string
	Maker                string
	Taker                string
	ListedAfter          uint64
	ListedBefore         uint64
}

type CollectionStats struct {
	FloorPrice   float64 `json:"floor_price"`
	AveragePrice float64 `json:"average_price"`
}

type CollectionStatsResponse struct {
	Stats CollectionStats `json:"stats"`
}

func RetrieveOffers(opt RetrieveOffersOption) RetrieveOffersResponse {
	router := "/orders/ethereum/seaport/offers"
	var queryParams []HttpUrlQueryParam
	if opt.Limit != 0 {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "limit", Value: fmt.Sprint(opt.Limit)})
	}
	if opt.OrderBy != "" {
		queryParams = append(queryParams, HttpUrlQueryParam{Key: "orderBy", Value: opt.OrderBy})
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
