# X-mob NFT trading watcher client

Mob-Watcher is a piece of software you can run on your local machine to update the [xmob smart contract](https://github.com/X-mob/metamob) to keep automatically trading NFTs. When Mob-Watcher start, it also establish a simple http api server to allow end users to query data from itself.

Every one can run a watcher node, and all the nodes suppose to form a decentralized network together and gain tokens as incentives in the future. If you are interesting at Xmob, please [contact us](https://twitter.com/XMob_eth) and help building the project~

## How to run

```sh
$ cat > ./.env <<EOF

RPC_URL=<required, ethereum rpc url, eg: infura/alchemy>
PRIVATE_KEY=<required, private key in HexString without 0x, for watcher account to send tx>
OPENSEA_API_KEY=<required, opensea api key>
NETWORK=<goerli/ethereum, optional, default to ethereum>
STORE_PATH=<db store path, optional, default to /tmp/badger>
```

then

```sh
go build
./mob-watcher help
```

## Development

update smart contract abi bindings

```sh
make abi
make abi-binding
```

We use `go version go1.16.7` for master development, and will upgrade to go1.19.2 to enable advance features like generic.
