# domestic_system

browserless sub web 3

#install docker
sudo service docker start
sudo docker run hello-world

#jpeg to ANSI
sudo snap install imgcat // graphic based
sudo apt install jp2a //txt based
curl https://live---metadata-5covpqijaa-uc.a.run.app/images/5268 --output - | jp2a --size=40x20 --colors --clear - //usage

#curl rpc commands running ganache-cli
curl http://127.0.0.1:8545 -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}'
curl http://127.0.0.1:8545 -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_accounts","params": [],"id":1}'

#future of the command line writeup
https://github.com/readme/featured/future-of-the-command-line

#more examples at the bottom of the podcast
https://changelog.com/podcast/481

#mnemonic generators
https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
24 word generator

#Generates mnemonic and ssh key from it
go run generate.go | melt restore - --seed - //replace first "-" with a filename to generate both pub and prv key
