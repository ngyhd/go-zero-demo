package logic

import (
	"context"
	"go-zero-demo/user/user"

	"go-zero-demo/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancellationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancellationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancellationLogic {
	return &CancellationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销
func (l *CancellationLogic) Cancellation(in *user.CancellationReq) (*user.CancellationResp, error) {
	// todo: add your logic here and delete this line

	return &user.CancellationResp{}, nil
}
