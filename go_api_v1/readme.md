# gen

protoc-gen-go v1.25.0

protoc -I. -I%GOPATH%/src --go_opt=paths=source_relative --go_out=. ./*.proto
protoc -I. -I$GOPATH/src --go_opt=paths=source_relative --go_out=. ./*.proto
