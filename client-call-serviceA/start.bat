@echo on

call stop.bat

rem set ZEUS=zeus_v2@v0.0.1
set ZEUS=zeus
if not exist %GOPATH%/pkg/mod/gitlab.ztgame.com/tech/public/go-service/%ZEUS% (
	echo "get zeus start" 
	set GOPROXY=
	git config http.extraheader "PRIVATE-TOKEN: AiyKkDd3XFzxmnQjXGgt"
	go get -v -insecure gitlab.ztgame.com/tech/public/go-service/%ZEUS%
	echo "get zeus stop"
)

set GOPROXY=https://goproxy.io
set GOBIN=%~dp0bin

go install ./src/ServerApp
go install ./src/Client

cd res\config
if not exist server.toml (copy server.toml.example server.toml)
cd ../../

start serverapp.bat
@ping -n 4 127.1>nul
start startClient.bat

