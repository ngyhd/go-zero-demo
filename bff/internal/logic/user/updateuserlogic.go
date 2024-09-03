package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/user/user"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	reqData := &user.UpdateUserReq{
		UserInfo: &user.UserInfo{
			Avatar:   req.Avatar,
			Nickname: req.Nickname,
			Password: req.Password,
			Bio:      gconv.PtrString(req.Bio),
			Gender:   req.Gender,
			Region:   req.Region,
		},
	}
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
