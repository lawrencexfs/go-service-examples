@echo on

call stop.bat

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

