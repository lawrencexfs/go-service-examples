cd bin

flag="app"
echo starting ServerApp2
./ServerApp $flag --configfile=../res/config/server2.toml --pprof-port=58080
exit
