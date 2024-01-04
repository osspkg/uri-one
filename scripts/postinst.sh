#!/bin/bash


if [ -f "/etc/systemd/system/urione.service" ]; then
    systemctl start urione
    systemctl enable urione
    systemctl daemon-reload
fi
