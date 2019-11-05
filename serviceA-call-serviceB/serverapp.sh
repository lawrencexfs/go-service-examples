cd bin

flag="app"
echo starting ServerApp
./ServerApp $flag --configfile=../res/config/server.toml --pprof-port=58080
ps ux | grep $flag |grep -v grep
exit
