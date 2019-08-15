#!/bin/bash
protoc --go_out=. --plugin=protoc-gen-go=$GOPATH/bin/protoc-gen-go ptbf.hello.proto