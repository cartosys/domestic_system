# domestic_system

browserless sub web 3

#install docker
sudo service docker start
sudo docker run hello-world

#curl rpc commands running ganache-cli
curl http://127.0.0.1:8545 -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}'
curl http://127.0.0.1:8545 -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_accounts","params": [],"id":1}'

#future of the command line writeup
https://github.com/readme/featured/future-of-the-command-line
