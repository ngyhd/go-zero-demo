package user

import (
	"context"
	"go-zero-demo/user/user"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	reqData := &user.LoginReq{
		Account:  req.Username,
		Password: req.Password,
	}
	_, err = l.svcCtx.UserRpc.Login(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
