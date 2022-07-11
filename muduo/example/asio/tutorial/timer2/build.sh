set -x

g++ timer.cc -I../../../../include -L../../../../lib -lmuduo_net -lmuduo_base -lpthread -o timer
