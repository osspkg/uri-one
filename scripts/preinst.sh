#!/bin/bash


if ! [ -d /var/lib/uri-one/ ]; then
    mkdir /var/lib/uri-one
fi

if [ -f "/etc/systemd/system/uri-one.service" ]; then
    systemctl stop uri-one
    systemctl disable uri-one
    systemctl daemon-reload
fi
