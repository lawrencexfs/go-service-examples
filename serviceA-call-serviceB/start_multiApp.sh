#!/bin/bash

flag="app"
bash stop.sh

# set ZEUS=zeus_v2@v0.0.1
export ZEUS=zeus
if [ ! -d "$GOPATH/pkg/mod/github.com/tech/public/go-service/$ZEUS" ]; then
	echo "get zeus start" 
	set GOPROXY=
	git config http.extraheader "PRIVATE-TOKEN: AiyKkDd3XFzxmnQjXGgt"
	go get -v -insecure github.com/tech/public/go-service/$ZEUS
	echo "get zeus stop"
fi

export GOPROXY=https://goproxy.io
export GOBIN=$PWD/bin

go install ./src/ServerApp

cd res/config
if [ ! -f  "server1.toml" ]; then
	cp server1.toml.example server1.toml
fi
if [ ! -f  "server2.toml" ]; then
	cp server2.toml.example server2.toml
fi

cd ../../


#cd bin
nohup ./serverapp1.sh $flag > ServerApp_`date +%Y-%m-%d_%H-%M-%S`.log &
#nohup ./ServerApp $flag --configfile=../res/config/server1.toml --pprof-port=58080 > ServerApp_`date +%Y-%m-%d_%H-%M-%S`.log &
sleep 5

nohup ./serverapp2.sh $flag > ServerApp_`date +%Y-%m-%d_%H-%M-%S`.log &
#nohup ./ServerApp $flag --configfile=../res/config/server2.toml --pprof-port=58081 > ServerApp_`date +%Y-%m-%d_%H-%M-%S`.log &
sleep 1


ps ux | grep $flag |grep -v grep
