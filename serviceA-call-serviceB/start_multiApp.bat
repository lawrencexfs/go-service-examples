@echo on

call stop.bat

rem set ZEUS=zeus_v2@v0.0.1
set ZEUS=zeus
if not exist %GOPATH%/pkg/mod/github.com/tech/public/go-service/%ZEUS% (
	echo "get zeus start" 
	set GOPROXY=
	git config http.extraheader "PRIVATE-TOKEN: AiyKkDd3XFzxmnQjXGgt"
	go get -v -insecure github.com/tech/public/go-service/%ZEUS%
	echo "get zeus stop"
)

set GOPROXY=https://goproxy.io
set GOBIN=%~dp0bin

go install ./src/ServerApp

cd res\config
if not exist server1.toml (copy server1.toml.example server1.toml)
if not exist server2.toml (copy server2.toml.example server2.toml)
cd ../../



start serverapp1.bat
@ping -n 4 127.1>nul
start serverapp2.bat

