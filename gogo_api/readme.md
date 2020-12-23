# gen

protoc-gen-gofast v1.25.0

protoc -I. -I%GOPATH%/src  --gogo_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:%GOPATH%/src ./*.proto
protoc -I. -I$GOPATH/src  --gogo_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types:$GOPATH/src ./*.proto
