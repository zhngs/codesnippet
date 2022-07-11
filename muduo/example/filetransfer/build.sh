set -x

g++ download3.cc -I../../include -L../../lib -lmuduo_base -lmuduo_net -lpthread -o uploadfile
