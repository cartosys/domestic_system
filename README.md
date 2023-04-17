# domestic_system

browserless sub web 3

#install docker
sudo service docker start
sudo docker run hello-world

#jpeg to ANSI
sudo snap install imgcat # graphic based
sudo apt install jp2a #txt based
curl https://live---metadata-5covpqijaa-uc.a.run.app/images/5268 --output - | jp2a --size=40x20 --colors --clear - #usage

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
go run generate.go | melt restore - --seed - #replace first "-" with a filename to generate both pub and prv key

docker run \
 -p 2222:22 \
 -v \$PWD/.wishlist:/.wishlist \
 docker.io/charmcli/wishlist:latest

# https://hub.docker.com/r/linuxserver/openssh-server

#ssh server in docker:

docker run -d --name=openssh-server --hostname=openssh-server -e SUDO_ACCESS=true -e PUBLIC_KEY_FILE="`cat ~/.ssh/domestic_system.pub`" -e USER_NAME=charmer -e PUID=1000 -e PGID=1000 -e TZ=Etc/UTC -p 2222:2222 -v /path/to/appdata/config:/config --restart unless-stopped lscr.io/linuxserver/openssh-server:latest

#remote in
docker exec -it openssh-server /bin/bash
ssh -i \$PWD/.ssh/domestic_system -p 2222 charmer@172.17.0.2

#logs
docker logs -f openssh-server

#stop server
docker stop openssh-server

#start server
docker start openssh-server

#remove
docker rm openssh-server

#find docker server ip
docker container ls
docker inspect <serverID from above> | grep '"IPAddress"' | head -n 1

#other cool projects:

#discord cli
https://www.reddit.com/r/commandline/comments/12lm6au/discordo_a_tui_interface_for_discord_that_doesnt/

#spotify cli
https://www.reddit.com/r/unixporn/comments/12mavz9/oc_spotify_for_terminal_with_cover_art_rendering/

#Soothing pastel theme for the high-spirited!
https://github.com/catppuccin

#generate music from linux entropy
https://www.reddit.com/r/commandline/comments/12p9zp4/generate_music_from_the_entropy_of_linux_linuxwave/
