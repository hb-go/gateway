module github.com/hb-chen/gateway/example/micro

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/hb-chen/gateway v0.0.0-20200430074931-45e413665a8f
	github.com/hb-go/grpc-contrib v0.0.0-20200504062014-d5f5132289c5
	github.com/micro/go-micro/v2 v2.4.0
	google.golang.org/grpc v1.26.0
)

replace github.com/hb-chen/gateway v0.0.0-20200430074931-45e413665a8f => ../../../gateway
