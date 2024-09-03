package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/user/user"
	"strings"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *types.GetUsersReq) (resp *types.GetUsersResp, err error) {
	reqData := &user.GetUsersReq{
		UserIds: gconv.SliceInt64(strings.Split(req.Ids, ",")),
	}
	_, err = l.svcCtx.UserRpc.GetUsers(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
