package opensea

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/X-mob/mob-watcher/config"
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
	Path        string // start with slash
	Method      ReqMethod
	QueryParams []HttpUrlQueryParam
	Payload     string
}

func SendRequest(opt SendRequestOption, baseUrl string) []byte {
	url := baseUrl + opt.Path
	req, _ := http.NewRequest(opt.Method.String(), url, strings.NewReader(opt.Payload))

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
	return SendRequest(opt, config.GlobalConfig.OpenseaApiV1BaseUrl)
}

func SendRequestV2(opt SendRequestOption) []byte {
	return SendRequest(opt, config.GlobalConfig.OpenseaApiV2BaseUrl)
}
