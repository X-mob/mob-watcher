package lib

func GenMobCallData(methodName string, args ...interface{}) []byte {
	mobAbi, _ := XmobExchangeCoreMetaData.GetAbi()
	data, err := mobAbi.Pack(methodName, args...)
	if err != nil {
		panic(err)
	}
	return data
}

func GenMobManagerCallData(methodName string, args ...interface{}) []byte {
	mobAbi, _ := XmobManageMetaData.GetAbi()
	data, err := mobAbi.Pack(methodName, args...)
	if err != nil {
		panic(err)
	}
	return data
}
