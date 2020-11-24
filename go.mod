module micro-playground

go 1.13

require (
	github.com/bcicen/grmon v0.0.0-20190725134940-6c3770b6af49
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.4.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcd v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v2 v2.9.3
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
