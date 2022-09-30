package config

//const XmobManagerAddress string = "0xe71C9Ec4CE47265d10B01C0cc0994617d43b3AeA"
//const SeaportAddress string = "0x00000000006c3852cbEf3e08E8dF289169EdE581"

type ContractAddress struct {
	XmobManagerAddress string
	SeaportAddress     string
}

var AddressConfig = map[Network]ContractAddress{
	Ethereum: {
		XmobManagerAddress: "",
		SeaportAddress:     "0x00000000006c3852cbEf3e08E8dF289169EdE581",
	},
	Goerli: {
		XmobManagerAddress: "0x218561E0885b836C7FB51C6CDF89a3C9Ab5cD62e",
		SeaportAddress:     "0x00000000006c3852cbEf3e08E8dF289169EdE581",
	},
	Rinkeby: {
		XmobManagerAddress: "0xA9d8C7848F9b390425CF5DD9DfA5BA1Bb2e447de",
		SeaportAddress:     "0x00000000006c3852cbEf3e08E8dF289169EdE581",
	},
}

func GetContractAddress(n Network) ContractAddress {
	return AddressConfig[n]
}
