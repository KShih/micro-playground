export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379
export MICRO_API_HANDLER=http

.PHONY: build
build:
	go build -o ./micro main.go

.PHONY: run
run:build
	./micro --registry=etcd api --handler=api --enable_rpc