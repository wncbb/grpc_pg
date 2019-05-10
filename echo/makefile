# The directory name is as same as the cmds, so add those cmds to .PHONY
.PHONY: server client generate

generate:
	protoc -I/usr/local/include -I. \
			-I$(GOPATH)/src \
			-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc:. \
			--grpc-gateway_out=logtostderr=true:. \
			--swagger_out=logtostderr=true:. \
			./pb/*.proto


rpc_srv:
	go run ./cmd/rpc_server/rpc_server.go

rpc_cli:
	go run ./cmd/rpc_client/rpc_client.go

gw_srv:
	go run ./cmd/gw_server/gw_server.go

