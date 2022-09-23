GIT_URL = https://github.com/RetricSu/xmob-contract.git
BRANCH = refactor-code

.PHONY: fmt

fmt:
	go fmt ./...
	
abi:
	rm -rf deps
	mkdir -p deps
	cd deps && git clone -b ${BRANCH} ${GIT_URL} metamob
	cd deps/metamob && yarn && yarn deps && yarn build
	jq .abi deps/metamob/artifacts/contracts/XmobManage.sol/XmobManage.json > artifacts/XmobManage.abi
	jq .abi deps/metamob/artifacts/contracts/XmobExchangeCore.sol/XmobExchangeCore.json > artifacts/XmobExchangeCore.abi


abi-binding:
	abigen --abi artifacts/XmobManage.abi --pkg lib --type XmobManage --out lib/XmobManage.go
	abigen --abi artifacts/XmobExchangeCore.abi --pkg lib --type XmobExchangeCore --out lib/XmobExchangeCore.go
