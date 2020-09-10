.PHONY: proto
proto:
	protoc -I . --micro_out=. --go_out=. ./proto/helloworld/greeter.proto

.PHONY: client
client:
	go run client/client.go client/plugin.go

.PHONY: serv
serv:
	go run main.go

.PHONY: client-reg
client-reg:
	go run client/client.go client/plugin.go --registry=etcd

.PHONY: serv-reg
serv-reg:
	go run main.go --registry=etcd
