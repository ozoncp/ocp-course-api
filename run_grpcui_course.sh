#!/bin/bash
grpcui -import-path ./api/ocp-course-api \
    -import-path ./vendor.protogen \
    -import-path $(go env GOPATH)/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.1 \
    -proto ./api/ocp-course-api/course_service.proto -plaintext 127.0.0.1:17002
