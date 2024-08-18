package logic

import (
	"context"
	"go-zero-demo/user/user"

	"go-zero-demo/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查用户信息
func (l *GetUsersLogic) GetUsers(in *user.GetUsersReq) (*user.GetUsersResp, error) {
	// todo: add your logic here and delete this line

	return &user.GetUsersResp{}, nil
}
