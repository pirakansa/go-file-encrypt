#!/bin/bash

# ```sh
# curl -sSf https://raw.githubusercontent.com/pirakansa/go-file-encrypt/master/scripts/uninstall.sh | sudo bash -
# ```

PROJECT_NAME="fenc"
INSTALL_DIR="/usr/bin"

rm -f  "${INSTALL_DIR}/${PROJECT_NAME}"
rm -fr "${CONFIGURE_DIR}"
rm -fr "${CACHE_DIR}"

echo "uninstalled"
