#! /bin/bash
apt-get install unzip
unzip -d ./protoc-21.12-linux-x86_64 protoc-21.12-linux-x86_64.zip
cd protoc-21.12-linux-x86_64/bin
mv protoc $GOPATH/bin
protoc --version
cd -
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
tar -zxvf etcd-v3.5.7-linux-amd64.tar.gz
cd etcd-v3.5.7-linux-amd64
cp etcd etcdctl /usr/local/bin
etcd --version
cd -