#!/usr/bin/bash

rm -rf examples/.terraform*

# Define standard parameters
HOSTNAME=example.com
NAMESPACE=learn
NAME=demo
BINARY=terraform-provider-${NAME}
# Major.minor.patch
VERSION=0.1.0
# OS_ARCH=darwin_amd64 ==> For Mac
OS_ARCH=linux_amd64
PLUGIN_PATH=".terraform.d/plugins"
OS=`uname`
if [[ "$OS" == *"_NT-"* ]]; then
  OS_ARCH=windows_amd64
  PLUGIN_PATH="AppData/Roaming/terraform.d/plugins"
  BINARY=terraform-provider-${NAME}.exe
fi

# Build go binary 
go build -o ${BINARY}

# Terraform look for providers in a specific directory structure
mkdir -p ~/${PLUGIN_PATH}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
mv ${BINARY} ~/${PLUGIN_PATH}/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

# cd examples
# terraform.exe init

# TEST?=$$(go list ./... | grep -v 'vendor')
# go test -i $(TEST) || exit 1
# echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4