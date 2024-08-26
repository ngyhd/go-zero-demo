package user

import (
	"context"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancellationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancellationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancellationLogic {
	return &CancellationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancellationLogic) Cancellation(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
