set -x

g++ echo.cc main.cc -I../../include -L../../lib -lmuduo_base -lmuduo_net -lpthread -o echo_idleconn
