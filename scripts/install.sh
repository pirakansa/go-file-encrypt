#!/bin/bash

# ```sh
# curl -sSf https://raw.githubusercontent.com/pirakansa/go-file-encrypt/master/scripts/install.sh | sudo bash -
# ```

PROJECT_NAME="fenc"
INSTALL_DIR="/usr/bin"
TMP_DIR=$(mktemp -d)
ARCH_TYPE=$(uname -m)
trap 'rm -fr ${TMP_DIR}; sync' EXIT

case ${ARCH_TYPE} in
    "x86_64"  ) ARCH_TYPE="amd64" ;;
    "aarch64" ) ARCH_TYPE="arm64" ;;
    "armv7l"  ) ARCH_TYPE="arm"   ;;
    *         ) echo "install error"
                exit 1 ;;
esac

curl -L http://127.0.0.1:1080/${PROJECT_NAME}_linux-${ARCH_TYPE}.tar.gz --output "${TMP_DIR}/${PROJECT_NAME}.tar.gz"
tar --gunzip --extract --directory=${TMP_DIR} --file="${TMP_DIR}/${PROJECT_NAME}.tar.gz"
install --mode 0755 "${TMP_DIR}/${PROJECT_NAME}" "${INSTALL_DIR}/${PROJECT_NAME}"

echo "installed"
