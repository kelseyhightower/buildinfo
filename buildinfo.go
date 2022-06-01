package buildinfo

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime/debug"

	"github.com/kelseyhightower/buildinfo/grpc_buildinfo"
	"google.golang.org/grpc"
)

var ErrBuildInfoNotSupported = errors.New("build information not supported")

type Server struct {
	grpc_buildinfo.UnimplementedBuildInfoServiceServer
}

func Register(s *grpc.Server) {
	grpc_buildinfo.RegisterBuildInfoServiceServer(s, &Server{})
}

func (s *Server) GetBuildInfo(ctx context.Context, r *grpc_buildinfo.EmptyRequest) (*grpc_buildinfo.BuildInfo, error) {
	var b grpc_buildinfo.BuildInfo

	if ctx.Err() == context.Canceled {
		return nil, status.Error(codes.Canceled, "Client cancelled, abandoning.")
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, ErrBuildInfoNotSupported
	}

	b.GoVersion = info.GoVersion
	b.Path = info.Path

	// Copy main module
	b.Main = &grpc_buildinfo.Module{}
	copyModules(b.Main, &info.Main)

	// Copy module dependencies
	b.Deps = make([]*grpc_buildinfo.Module, 0)
	for _, dep := range info.Deps {
		var m grpc_buildinfo.Module
		copyModules(&m, dep)
		b.Deps = append(b.Deps, &m)
	}

	// Copy settings
	b.Settings = make([]*grpc_buildinfo.BuildSetting, 0)
	for _, setting := range info.Settings {
		s := &grpc_buildinfo.BuildSetting{
			Key:   setting.Key,
			Value: setting.Value,
		}
		b.Settings = append(b.Settings, s)
	}

	return &b, nil
}

func copyModules(dst *grpc_buildinfo.Module, src *debug.Module) {
	dst.Path = src.Path
	dst.Sum = src.Sum
	dst.Version = src.Version

	var dstModule = dst.Replace
	var srcModule = src.Replace
	for {
		if srcModule == nil {
			break
		}
		dstModule.Path = srcModule.Path
		dstModule.Sum = srcModule.Sum
		dstModule.Version = srcModule.Version

		dstModule = dstModule.Replace
		srcModule = srcModule.Replace
	}
}
