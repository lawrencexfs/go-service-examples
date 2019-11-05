#!/bin/bash

flag="app"
bash stop.sh

# set ZEUS=zeus_v2@v0.0.1
export ZEUS=zeus
if [ ! -d "$GOPATH/pkg/mod/gitlab.ztgame.com/tech/public/go-service/$ZEUS" ]; then
	echo "get zeus start" 
	set GOPROXY=
	git config http.extraheader "PRIVATE-TOKEN: AiyKkDd3XFzxmnQjXGgt"
	go get -v -insecure gitlab.ztgame.com/tech/public/go-service/$ZEUS
	echo "get zeus stop"
fi

export GOPROXY=https://goproxy.io
export GOBIN=$PWD/bin

go install ./src/ServerApp

cd res/config
if [ ! -f  "server.toml" ]; then
	cp server.toml.example server.toml
fi

cd ../../

nohup ./serverapp.sh $flag > ServerApp_`date +%Y-%m-%d_%H-%M-%S`.log &
ps ux | grep $flag |grep -v grep
