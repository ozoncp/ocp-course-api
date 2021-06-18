#!/bin/bash
grpcui -import-path ./api/ocp-lesson-api \
    -import-path ./vendor.protogen \
    -import-path $(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
    -proto ./api/ocp-lesson-api/lesson_service.proto -plaintext 127.0.0.1:27002
