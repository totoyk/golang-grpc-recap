#!/bin/sh

protoc proto/helloworld.proto \
  --js_out=import_style=commonjs:client/src/helloworld \
  --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:client/src/helloworld \
  --go_out=api/helloworld \
  --go-grpc_out=api/helloworld \
  --descriptor_set_out=./proxy/pb/helloworld.pb
