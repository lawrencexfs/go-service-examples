@echo on

call stop.bat
call flushRedis.bat


set GOPROXY=https://goproxy.io
set GOBIN=%~dp0bin

go install ./src/ServerApp

cd res\config
if not exist server.toml (copy server.toml.example server.toml)
cd ../../

start serverapp.bat
