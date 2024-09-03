package user

import (
	"context"
	"go-zero-demo/user/user"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	reqData := &user.RegisterReq{
		Account:  req.Username,
		Password: req.Password,
	}
	_, err = l.svcCtx.UserRpc.Register(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
