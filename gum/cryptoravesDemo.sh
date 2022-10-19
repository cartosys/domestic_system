#!/usr/bin/env bash
#set up tmp file system
STORAGEFOLDER=/tmp/domestic_system_storage/cryptoraves
[ ! -d "${STORAGEFOLDER}" ] && mkdir -p ${STORAGEFOLDER}

#get number of decimals from ERC20
NODECIMALS=$(curl -s 'http://127.0.0.1:8545/' \
    -X POST \
    -H "Content-Type: application/json" \
    -d '{"jsonrpc":"2.0","method":"eth_call","params": [
    {"from": "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266",
     "to": "0x6a8413f056f1e53bc58a8a04fa746be4c6131336",
     "gas": "0x76c0","gasPrice": "0x0","value": "0x0",
     "data": "0x313ce567"}, "latest"],"id":1}' | jq '.result' | tr -d '"' )  #data field is Keccak256("totalSupply()")     https://emn178.github.io/online-tools/keccak_256.html

NODECIMALS=$(echo "${NODECIMALS:2}" | sed 's/^0*//' | tr '[:lower:]' '[:upper:]')
NODECIMALS=$(echo "obase=10; ibase=16; ${NODECIMALS}" | bc)
echo ${NODECIMALS}
echo "Number of decimals in ERC20 contract: "$NODECIMALS


#decimals() 0x313ce567
#totalSupply() 0x18160ddd




#get eth block Number
#curl -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'

#gets total supply of ERC20
#TODO hex value is too large. how to  handle bigints in bash?
#TOTALSUPPLYDATA=$(curl -s 'http://127.0.0.1:8545/' \
#    -X POST \
#    -H "Content-Type: application/json" \
#    -d '{"jsonrpc":"2.0","method":"eth_call","params": [
#    {"from": "0x185B520F3536362DF300EF5f5bb1ebb47900A7BF",
#     "to": "0x8cEd2B545A881c220F023fa9DfBf5E466a29E6D5",
#     "gas": "0x76c0","gasPrice": "0x0","value": "0x0",
#     "data": "0x18160ddd"}, "latest"],"id":1}' | jq '.result' | tr -d '"')  #data field is Keccak256("totalSupply()")

#TOTALSUPPLYDATA=$(echo "${TOTALSUPPLYDATA:2}" | sed 's/^0*//')

#echo "Total Supply in ERC20 contract: "$TOTALSUPPLYDATA
#TOTALSUPPLYDATA=$(echo $((${TOTALSUPPLYDATA})) )
#echo "Total Supply in ERC20 contract: "$TOTALSUPPLYDATA
