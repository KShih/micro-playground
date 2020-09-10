.PHONY: proto
proto:
	protoc -I . --micro_out=. --go_out=. ./proto/helloworld/greeter.proto

.PHONY: client
client:
	go run client/client.go client/plugin.go

.PHONY: serv
serv:
	go run main.go

.PHONY: api-serv
api-serv:
	go run api/api.go

.PHONY: api-micro
api-micro:
	micro api --handler=api

.PHONY: client-etcd
client-reg:
	go run client/client.go client/plugin.go --registry=etcd

.PHONY: serv-etcd
serv-reg:
	go run main.go --registry=etcd
