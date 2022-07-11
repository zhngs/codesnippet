#!/bin/bash
# set -x

if [[ "$1" = "install" ]]; then
    cp pullfile /bin/
    cp pushfileTo /bin/
fi

if [[ "$1" = "uninstall" ]]; then
    rm /bin/pullfile
    rm /bin/pushfileTo
fi