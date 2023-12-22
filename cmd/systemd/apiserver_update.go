package systemd

const ApiserverUpdateScript = `#!/bin/bash

# This script is to update the apiserver binary. We have to download the new
# binary and run it instead of the currently executing binary, which runs the
# older version.

if [ $# -ne 1 ]; then
    echo "ApiServer version must not be empty. Usage: $0 v0.1.7"
    exit 1
fi

ver=$1
url="https://github.com/NaoNaoOnline/apiserver/releases/download/${ver}/apiserver-linux-amd64"
pat="/usr/local/bin/apiserver-new"
bin="/usr/local/bin/apiserver"
uni="apiserver.daemon.service"

# Download the binary
curl --location "${url}" --output "${pat}"

# Check if download was successful
if [ $? -ne 0 ]; then
    echo "Failed to download the binary. Exiting."
    exit 1
fi

# Make the binary executable
chmod +x "${pat}"

# Stop the systemd unit
systemctl stop "${uni}"

# Replace the binary
mv "${pat}" "${bin}"

# Start the systemd unit
systemctl start "${uni}"

# Print the unit status
systemctl status "${uni}"`
