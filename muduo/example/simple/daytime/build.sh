set -x

g++ daytime.cc main.cc -I../../../include -L../../../lib -lmuduo_base -lmuduo_net -lpthread -o daytime
