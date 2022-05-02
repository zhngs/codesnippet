set -x

g++ server.cc -I../../../include -L../../../lib -lmuduo_net -lmuduo_base -lpthread -o discard_server
g++ client.cc -I../../../include -L../../../lib -lmuduo_net -lmuduo_base -lpthread -o discard_client
