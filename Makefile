GIT_URL = https://github.com/X-mob/metamob.git
BRANCH = main

.PHONY: fmt clean

fmt:
	go fmt ./...
	
abi:
	rm -rf deps
	mkdir -p deps
	cd deps && git clone -b ${BRANCH} ${GIT_URL} metamob
	cd deps/metamob && yarn && yarn deps && yarn build
	jq .abi deps/metamob/artifacts/contracts/XmobManage.sol/XmobManage.json > artifacts/XmobManage.abi
	jq .abi deps/metamob/artifacts/contracts/XmobExchangeCore.sol/XmobExchangeCore.json > artifacts/XmobExchangeCore.abi
	jq .abi deps/metamob/artifacts/contracts/deps/seaport/Seaport.sol/Seaport.json > artifacts/Seaport.abi


abi-binding:
	abigen --abi artifacts/XmobManage.abi --pkg lib --type XmobManage --out lib/XmobManage.go
	abigen --abi artifacts/XmobExchangeCore.abi --pkg lib --type XmobExchangeCore --out lib/XmobExchangeCore.go
	abigen --abi artifacts/Seaport.abi --pkg seaport --type Seaport --out seaport/Seaport.go

clean-db:
	rm -rf kvstore
