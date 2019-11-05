@echo on

call stop.bat

rem set ZEUS=zeus_v2@v0.0.1
set ZEUS=zeus
rem if not exist %GOPATH%/pkg/mod/github.com/GA-TECH-SERVER/%ZEUS% (
rem 	echo "get zeus start" 
rem 	set GOPROXY=
rem 	git config http.extraheader "PRIVATE-TOKEN: AiyKkDd3XFzxmnQjXGgt"
rem 	go get -v -insecure github.com/GA-TECH-SERVER/%ZEUS%
rem 	echo "get zeus stop"
rem )

set GOPROXY=https://goproxy.io
set GOBIN=%~dp0bin

go install ./src/ServerApp

cd res\config
if not exist server.toml (copy server.toml.example server.toml)
cd ../../


start serverapp.bat
