#!/bin/bash

set -e

[ -z "${API_KEY}" ] && echo "Can't build image: Missing required API key. Specify the ${API_KEY} variable." && exit 1
[ -z "${BASIC_AUTH_USERNAME}" ] && echo "Can't build image: Missing required basic auth username. Specify the ${BASIC_AUTH_USERNAME} variable." && exit 1
[ -z "${BASIC_AUTH_PASSWORD}" ] && echo "Can't build image: Missing required basic auth password. Speicfy the ${BASIC_AUTH_PASSWORD} variable." && exit 1

BUILD_MODE=${BUILD_MODE:-appc}
BUILD_DIR=build

IMG_VERSION=${IMG_VERSION:-1.0.0}
IMG_ARCH=${IMG_ARCH:-amd64}
IMG_OS=${IMG_OS:-linux}
IMG_NAME=blued-${IMG_VERSION}-${IMG_OS}-${IMG_ARCH}.aci

PACKAGE_SERVER=${PACKAGE_SERVER:-github.com/ihcsim/bluelens/cmd/blued}
PACKAGE_CLI=${PACKAGE_CLI:-github.com/ihcsim/bluelens/tool/blue}

# build statically linked binary
go build -v -o blued -ldflags '-linkmode external -extldflags "-static"' ${PACKAGE_SERVER}

function acbuild_end {
  rm -f blued blue
  export EXIT=$?
  acbuild --debug end && exit $EXIT
}

# begin building image
acbuild --debug begin --build-mode=${BUILD_MODE}
trap acbuild_end EXIT

acbuild --debug annotation add authors "Ivan Sim <ihcsim@gmail.com>"

home=/opt/blued
www=/var/www/blued
acbuild --debug copy-to-dir blued ${home}
acbuild --debug copy-to-dir etc/*.json  ${home}/data/
acbuild --debug copy-to-dir cmd/blued/swagger/* ${www}/
acbuild --debug copy tls /etc/ssl

acbuild --debug set-exec -- ${home}/blued \
  -followees ${home}/data/followees.json -history ${home}/data/history.json -music ${home}/data/music.json \
  -apikey=${API_KEY} -user=${BASIC_AUTH_USERNAME} -password=${BASIC_AUTH_PASSWORD} \
  -private /etc/ssl/localhost.key -cert /etc/ssl/localhost.crt
acbuild --debug port add www tcp 443

if [ "${BUILD_MODE}" = "appc" ];
then
  acbuild --debug set-name isim/blued
  acbuild --debug label add version ${IMG_VERSION}
  acbuild --debug label add arch ${IMG_ARCH}
  acbuild --debug label add os ${IMG_OS}
fi

mkdir -p ${BUILD_DIR}
acbuild --debug write --overwrite ${BUILD_DIR}/${IMG_NAME}
