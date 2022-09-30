# X-mob NFT trading watcher client

How to run

```sh
$ cat > ./.env <<EOF

RPC_URL=<required, ethereum rpc url, eg: infura>
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

update abi bindings

```sh
make abi
make abi-binding
```
