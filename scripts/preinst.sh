#!/bin/bash

if test -f "/etc/systemd/system/uri-one.service"; then
    systemctl stop uri-one
    systemctl disable uri-one

    systemctl daemon-reload
    systemctl reset-failed
fi

if ! [ -d /var/lib/urione/ ]; then
    mkdir /var/lib/urione
fi