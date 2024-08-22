package interceptor

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/pkg/xerr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		if xerr.IsBusinessError(causeErr) {
			e := xerr.FromBusinessError(causeErr)
			err = status.Error(codes.Code(e.GetCode()), e.GetMessage())
		} else {
			// 只记录系统错误
			logx.WithContext(ctx).Errorf("【RPC-ERR】 %+v", err)
		}
	}
	return resp, err
}
