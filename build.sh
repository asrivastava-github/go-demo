#!/usr/bin/bash

# Define standard parameters
HOSTNAME=registry.terraform.io
NAMESPACE=learn
NAME=demo
BINARY=terraform-provider-${NAME}
# Major.minor.patch
VERSION=0.1.0
OS_ARCH=windows_amd64

# Build go binary 
go build -o ${BINARY}

# Terraform look for providers in a specific directory structure
mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

# TEST?=$$(go list ./... | grep -v 'vendor')
# go test -i $(TEST) || exit 1
# echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4