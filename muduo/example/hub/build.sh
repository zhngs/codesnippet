set -x

g++ hub.cc codec.cc -I../../include -L../../lib -lmuduo_inspect -lmuduo_net -lmuduo_base -lpthread -o hub
g++ sub.cc pubsub.cc codec.cc -I../../include -L../../lib -lmuduo_net -lmuduo_base -lpthread -o sub
g++ pub.cc pubsub.cc codec.cc -I../../include -L../../lib -lmuduo_net -lmuduo_base -lpthread -o pub
