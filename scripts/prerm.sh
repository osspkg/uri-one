#!/bin/bash

if test -f "/etc/systemd/system/uri-one.service"; then
    systemctl stop uri-one
    systemctl disable uri-one

    systemctl daemon-reload
    systemctl reset-failed
fi