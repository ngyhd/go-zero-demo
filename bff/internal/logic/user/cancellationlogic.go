package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/user/user"

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

func (l *CancellationLogic) Cancellation(req *types.CancellationReq) (resp *types.LoginResp, err error) {
	reqData := &user.CancellationReq{
		UserId: gconv.Int64(req.UserId),
	}
	_, err = l.svcCtx.UserRpc.Cancellation(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
