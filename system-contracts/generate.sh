#!/bin/bash

docker run -v $PWD:$PWD -w $PWD ethereum/solc:0.7.0 @openzeppelin/=$(pwd)/node_modules/@openzeppelin/ --overwrite --abi --bin -o ./build \
contracts/*.sol

./bin/abigen --abi ./build/CacheBridge.abi --bin ./build/CacheBridge.bin --type CacheBridge --pkg generated --out ./generated/cache_bridge.go
./bin/abigen --abi ./build/OdinDataset.abi --bin ./build/OdinDataset.bin --type OdinDataset --pkg generated --out ./generated/odin_dataset.go
