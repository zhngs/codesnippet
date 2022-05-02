set -x

g++ roundtrip.cc -I../../include -L../../lib -lmuduo_base -lmuduo_net -lpthread -o roundtrip
