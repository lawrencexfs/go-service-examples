@echo on

call stop.bat
call flushRedis.bat

rem set ZEUS=zeus@v0.0.29
set ZEUS=go-service
rem if not exist %GOPATH%/pkg/mod/github.com/GA-TECH-SERVER/%ZEUS% (
rem 	echo "get zeus start" 
rem  	set GOPROXY=
rem 	git config http.extraheader "PRIVATE-TOKEN: AiyKkDd3XFzxmnQjXGgt"
rem 	go get -v -insecure github.com/giant-tech/go-service/%ZEUS%
rem 	echo "get zeus stop"
rem )

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
