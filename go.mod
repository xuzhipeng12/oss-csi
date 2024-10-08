module oss-csi

go 1.21

//require (
//	github.com/container-storage-interface/spec v1.10.0
//	google.golang.org/grpc v1.65.0
//	k8s.io/klog v1.0.0
//)
//
//require (
//	golang.org/x/net v0.25.0 // indirect
//	golang.org/x/sys v0.20.0 // indirect
//	golang.org/x/text v0.15.0 // indirect
//	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
//	google.golang.org/protobuf v1.34.1 // indirect
//)

require (
	github.com/aliyun/aliyun-oss-go-sdk v3.0.2+incompatible
	github.com/container-storage-interface/spec v1.10.0
	google.golang.org/grpc v1.65.0
	k8s.io/klog v1.0.0
)

require (
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
