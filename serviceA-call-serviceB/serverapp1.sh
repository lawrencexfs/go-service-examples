cd bin

flag="app"
echo starting ServerApp1
./ServerApp $flag --configfile=../res/config/server1.toml --pprof-port=58080
ps ux | grep $flag |grep -v grep
exit
