#!/usr/bin/env bash
#set up tmp file system
STORAGEFOLDER=/tmp/domestic_system_storage/cryptoraves
[ ! -d "${STORAGEFOLDER}" ] && mkdir -p ${STORAGEFOLDER}

#get number of decimals from ERC20
NODECIMALS=$(curl -s 'http://127.0.0.1:8545/' \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_call","params": [
    {"from": "0x185B520F3536362DF300EF5f5bb1ebb47900A7BF",
     "to": "0x8cEd2B545A881c220F023fa9DfBf5E466a29E6D5",
     "gas": "0x76c0","gasPrice": "0x0","value": "0x0",
     "data": "313ce567"}, "latest"],"id":1}' | jq '.result')  #data field is Keccak256("totalSupply()")     https://emn178.github.io/online-tools/keccak_256.html
echo "${NODECIMALS:3}" 
printf "%d\n" "${NODECIMALS}"


#gets total supply of ERC20
TOTALSUPPLYDATA=$(curl -s 'http://127.0.0.1:8545/' \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_call","params": [
    {"from": "0x185B520F3536362DF300EF5f5bb1ebb47900A7BF",
     "to": "0x8cEd2B545A881c220F023fa9DfBf5E466a29E6D5",
     "gas": "0x76c0","gasPrice": "0x0","value": "0x0",
     "data": "0x18160ddd"}, "latest"],"id":1}' | jq '.result' )  #data field is Keccak256("totalSupply()")

echo "${TOTALSUPPLYDATA:3}"
printf "%d\n" "${TOTALSUPPLYDATA}"
