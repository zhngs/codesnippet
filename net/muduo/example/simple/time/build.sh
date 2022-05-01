set -x

g++ time.cc main.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o timeserver
g++ timeclient.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o timeclient
