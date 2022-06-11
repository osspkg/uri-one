#!/bin/bash

#!/bin/bash

if [ -f "/etc/systemd/system/uri-one.service" ]; then
    systemctl start uri-one
    systemctl enable uri-one
fi