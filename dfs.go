package dfs

//go:generate protoc --proto_path=./proto-dfs --go_out=plugins=grpc,paths=source_relative:./network dfs-network.proto
//go:generate protoc --proto_path=./proto-dfs --go_out=plugins=grpc,paths=source_relative:./dag dfs-dag.proto
