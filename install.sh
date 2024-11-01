#!/bin/bash

set -e

# Detect the operating system
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" == "x86_64" ]; then
    ARCH="amd64"
elif [[ "$ARCH" == "arm" || "$ARCH" == "aarch64" ]]; then
    ARCH="arm64"
fi

# Set the download URL
VERSION=${1:-"latest"}
if [ "$VERSION" == "latest" ]; then
    URL=$(curl -s https://api.github.com/repos/bjess9/pr-pilot/releases/latest | grep "browser_download_url" | grep "$OS" | grep "$ARCH" | cut -d '"' -f 4)
else
    URL="https://github.com/bjess9/pr-pilot/releases/download/$VERSION/prpilot_${OS}_${ARCH}.tar.gz"
fi

# Download and extract
curl -L $URL -o prpilot.tar.gz
tar -xzf prpilot.tar.gz prpilot
rm prpilot.tar.gz

# Move to /usr/local/bin
sudo mv prpilot /usr/local/bin/
echo "PR Pilot installed successfully! Run 'prpilot configure' to get started."
