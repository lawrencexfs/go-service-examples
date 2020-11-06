@goto :main

�� common/proto/ �� svr_proto/ �� proto �ļ����� *.pb.go �ļ���

:main

@set ROOT=..\..
@set PROTO_DIR=%ROOT%\res\protoFile
@set SRC=%ROOT%\src
@set GOGOFASTER="--gogofaster_out=%SRC%\pb"

protoc -I%PROTO_DIR% %GOGOFASTER% game.proto


@echo off
set u2d=%ROOT%\tools\unix2dos\unix2dos.exe
for /R "%SRC%\pb" %%f in (*.go) do (
    %u2d% %%f > NUL 2>&1
)
for /R "%SRC%\spb" %%f in (*.go) do (
    %u2d% %%f > NUL 2>&1
)
echo on

pause