set -x

g++ chargen.cc main.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chargen
g++ chargenclient.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o chargenclient
