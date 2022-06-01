# Go buildinfo package Example

## Library Usage

```
package main

import (
	"net"
	"log"

	"google.golang.org/grpc"
	"github.com/kelseyhightower/buildinfo"
)

func main() {
	tpcListener, err := net.Listen("tcp", "127.0.0.1:10080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	buildinfo.Register(grpcServer)

	log.Fatal(grpcServer.Serve(tpcListener))
}
```

Call it with grpcurl:

```
cat "{}" | \
  grpcurl -d @ \
  --plaintext 127.0.0.1:10080 \
  buildinfo.BuildInfoService.GetBuildInfo
```

## Usage

Build the binary:

```
cd cmd
```

```
$ go build -o buildinfo .
```

Get the build info for the resulting binary:

```
$ ./buildinfo -binary buildinfo
```

Output

```
Getting build info for the buildinfo binary...
Go version: go1.18beta1
Build Settings: 
  -compiler: gc
  CGO_ENABLED: 1
  CGO_CFLAGS: 
  CGO_CPPFLAGS: 
  CGO_CXXFLAGS: 
  CGO_LDFLAGS: 
  GOARCH: arm64
  GOOS: darwin
Dependencies: 
  github.com/beorn7/perks: v1.0.1 h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=
  github.com/cespare/xxhash/v2: v2.1.1 h1:6MnRN8NT7+YBpUIWxHtefFZOKTAPgGjpQSxqLNn0+qY=
  github.com/golang/protobuf: v1.4.3 h1:JjCZWpVbqXDqFVmTfYWEVTMIYrL/NPdPSCHPJ0T/raM=
  github.com/matttproud/golang_protobuf_extensions: v1.0.1 h1:4hp9jkHxhMHkqkrB3Ix0jegS5sx/RkqARlsWZ6pIwiU=
  github.com/prometheus/client_golang: v1.11.0 h1:HNkLOAEQMIDv/K+04rukrLx6ch7msSRwf3/SASFAGtQ=
  github.com/prometheus/client_model: v0.2.0 h1:uq5h0d+GuxiXLJLNABMgp2qUWDPiLvgCzz2dUR+/W/M=
  github.com/prometheus/common: v0.26.0 h1:iMAkS2TDoNWnKM+Kopnx/8tnEStIfpYA0ur0xQzzhMQ=
  github.com/prometheus/procfs: v0.6.0 h1:mxy4L2jP6qMonqmq+aTtOx1ifVWUgG/TAmntgbh3xv4=
  golang.org/x/sys: v0.0.0-20210603081109-ebe580a85c40 h1:JWgyZ1qgdTaF3N3oxC+MdTV7qvEEgHo3otj+HB5CM7Q=
  google.golang.org/protobuf: v1.26.0-rc.1 h1:7QnIQpGRHE5RnLKnESfDoxm2dTapTZua5a0kS0A+VXQ=
```