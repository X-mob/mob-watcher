package opensea

import "github.com/ethereum/go-ethereum/common"

type OfferItem struct {
	ItemType             uint8  `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
}

type ConsiderationItem struct {
	ItemType             uint8  `json:"itemType"`
	Token                string `json:"token"`
	IdentifierOrCriteria string `json:"identifierOrCriteria"`
	StartAmount          string `json:"startAmount"`
	EndAmount            string `json:"endAmount"`
	Recipient            string `json:"recipient"`
}

type OrderParameters struct {
	Offerer                         common.Address      `json:"offerer"`
	Zone                            common.Address      `json:"zone"`
	ZoneHash                        string              `json:"zoneHash"`
	StartTime                       string              `json:"startTime"`
	EndTime                         string              `json:"endTime"`
	OrderType                       uint8               `json:"orderType"`
	Salt                            string              `json:"salt"`
	ConduitKey                      string              `json:"conduitKey"`
	Offer                           []OfferItem         `json:"offer"`
	Consideration                   []ConsiderationItem `json:"consideration"`
	TotalOriginalConsiderationItems int                 `json:"totalOriginalConsiderationItems"`
	Counter                         int                 `json:"counter"`
}

type ProtocolData struct {
	Parameters OrderParameters `json:"parameters"`
	Signature  string          `json:"signature"`
}

type Asset struct {
	Id                   string     `json:"id"`
	ImageUrl             string     `json:"image_url"`
	ImagePreviewUrl      string     `json:"image_preview_url"`
	ImageThumbnailUrl    string     `json:"image_thumbnail_url"`
	ImageOriginalUrl     string     `json:"image_original_url"`
	AnimationUrl         string     `json:"animation_url"`
	AnimationOriginalUrl string     `json:"animation_original_url"`
	Name                 string     `json:"name"`
	Description          string     `json:"description"`
	Collection           Collection `json:"collection"`
}

type AssetBundle struct {
	assets []Asset
}
type Order struct {
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

type RetrieveResponse struct {
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Orders   []Order `json:"orders"`
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

type RetrieveListingsOption struct {
	Limit                int
	AssetContractAddress string
	TokenIds             []string
}

type CollectionStats struct {
	FloorPrice   float64 `json:"floor_price"`
	AveragePrice float64 `json:"average_price"`
}

type CollectionStatsResponse struct {
	Stats CollectionStats `json:"stats"`
}

type Collection struct {
	Name           string          `json:"name"`
	Slug           string          `json:"slug"`
	ExternalLink   string          `json:"external_link"`
	ImageUrl       string          `json:"image_url"`
	BannerImageUrl string          `json:"banner_image_url"`
	Description    string          `json:"description"`
	Stats          CollectionStats `json:"stats"`
}

type AssetContract struct {
	Name              string     `json:"name"`
	CreatedDate       string     `json:"created_date"`
	Address           string     `json:"address"`
	Collection        Collection `json:"collection"`
	ImageUrl          string     `json:"image_url"`
	SchemaName        string     `json:"schema_name"`
	Symbol            string     `json:"symbol"`
	TotalSupply       string     `json:"total_supply"`
	ExternalLink      string     `json:"external_link"`
	Description       string     `json:"description"`
	AssetContractType string     `json:"asset_contract_type"`
}
type GetAssetsResponse struct {
	Id                   int           `json:"id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	BackgroundColor      string        `json:"background_color"`
	ImageUrl             string        `json:"image_url"`
	ImagePreviewUrl      string        `json:"image_preview_url"`
	ImageThumbnailUrl    string        `json:"image_thumbnail_url"`
	AnimationUrl         string        `json:"animation_url"`
	AnimationOriginalUrl string        `json:"animation_original_url"`
	AssetContract        AssetContract `json:"asset_contract"`
	Collection           Collection    `json:"collection"`
}

type CreateListingPayload struct {
	Parameters OrderParameters `json:"parameters"`
	Signature  string          `json:"signature"`
}
