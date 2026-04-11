package svc

import (
	"errors"

	"go-zero-demo/pkg/xerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TranslateRpcError 将 RPC 错误转换为业务错误
func TranslateRpcError(err error) error {
	if err == nil {
		return nil
	}

	// 尝试解析为业务错误
	var bizErr xerr.BusinessError
	if errors.As(err, &bizErr) {
		return bizErr
	}

	// 从 gRPC status 中提取错误信息
	s, ok := status.FromError(err)
	if !ok {
		// 非 gRPC 错误，返回系统错误
		return xerr.SystemErr.SetMessage(err.Error())
	}

	// 根据 gRPC 状态码判断
	switch s.Code() {
	case codes.NotFound:
		return xerr.NotFoundErr
	case codes.Unauthenticated:
		return xerr.UnauthorizedErr
	case codes.PermissionDenied:
		return xerr.NoPermissionErr
	default:
		return xerr.SystemErr.SetMessage(s.Message())
	}
}
