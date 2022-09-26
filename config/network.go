package config

import "log"

// blockchain network

type Network int64

const (
	Ethereum Network = iota
	Goerli
	Kovan
	Rinkeby
)

func (s Network) String() string {
	switch s {
	case Ethereum:
		return "ethereum"
	case Goerli:
		return "goerli"
	case Kovan:
		return "kovan"
	case Rinkeby:
		return "rinkeby"
	}
	return "unknown"
}

func GetNetwork(s string) Network {
	switch s {
	case "ethereum":
		return Ethereum
	case "goerli":
		return Goerli
	case "kovan":
		return Kovan
	case "rinkeby":
		return Rinkeby
	}

	log.Fatalf("unknown %s, use default network ethereum", s)
	return Ethereum
}

func GetOpenSeaApiBaseUrlByNetwork(n Network) (OpenseaApiV1BaseUrl string, OpenseaApiV2BaseUrl string) {
	if n == Ethereum { // mainnet
		const v1 = "https://api.opensea.io/api/v1"
		const v2 = "https://api.opensea.io/v2"

		return v1, v2
	}

	const v1 = "https://testnets-api.opensea.io/api/v1"
	const v2 = "https://testnets-api.opensea.io/v2"
	return v1, v2
}

func GetOpenSeaV1BaseUrl(n Network) string {
	v1, _ := GetOpenSeaApiBaseUrlByNetwork(n)
	return v1
}

func GetOpenSeaV2BaseUrl(n Network) string {
	_, v2 := GetOpenSeaApiBaseUrlByNetwork(n)
	return v2
}
