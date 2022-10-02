package config

//const XmobManagerAddress string = "0xe71C9Ec4CE47265d10B01C0cc0994617d43b3AeA"
//const SeaportAddress string = "0x00000000006c3852cbEf3e08E8dF289169EdE581"

type ContractAddress struct {
	XmobManagerAddress    string
	SeaportAddress        string
	SeaportConduitAddress string
}

var AddressConfig = map[Network]ContractAddress{
	Ethereum: {
		XmobManagerAddress:    "",
		SeaportAddress:        "0x00000000006c3852cbEf3e08E8dF289169EdE581",
		SeaportConduitAddress: "0x1E0049783F008A0085193E00003D00cd54003c71",
	},
	Goerli: {
		XmobManagerAddress:    "0xbEA5755210639fa29808bbF240e3507a87E3020E",
		SeaportAddress:        "0x00000000006c3852cbEf3e08E8dF289169EdE581",
		SeaportConduitAddress: "0x1E0049783F008A0085193E00003D00cd54003c71",
	},
	Rinkeby: {
		XmobManagerAddress:    "0xA9d8C7848F9b390425CF5DD9DfA5BA1Bb2e447de",
		SeaportAddress:        "0x00000000006c3852cbEf3e08E8dF289169EdE581",
		SeaportConduitAddress: "0x1E0049783F008A0085193E00003D00cd54003c71",
	},
}

func GetContractAddress(n Network) ContractAddress {
	return AddressConfig[n]
}
