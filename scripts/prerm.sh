#!/bin/bash


if [ -f "/etc/systemd/system/urione.service" ]; then
    systemctl stop urione
    systemctl disable urione
    systemctl daemon-reload
fi
