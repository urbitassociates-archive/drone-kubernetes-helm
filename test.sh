#!/bin/sh

set -e

source .env

sed -e '
  s/{{KUBE_CONFIG}}/'$KUBE_CONFIG'/g;
  s/{{KUBE_CA}}/'$KUBE_CA'/g;
  s/{{KUBE_CLIENT_CERT}}/'$KUBE_CLIENT_CERT'/g;
  s/{{KUBE_CLIENT_KEY}}/'$KUBE_CLIENT_KEY'/g;
' ./test-drone-config.json.tpl > ./test-drone-config.json
docker build --no-cache . -t test
docker run -i test < test-drone-config.json
