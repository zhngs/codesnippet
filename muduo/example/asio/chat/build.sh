set -x

g++ server.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chatserver
g++ server_threaded.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chatserver_threaded
g++ server_threaded_efficient.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chatserver_threaded_efficient
g++ server_threaded_highperformance.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chatserver_threaded_highperformance
g++ client.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chatclient
