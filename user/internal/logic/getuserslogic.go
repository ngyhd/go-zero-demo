package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"

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
	findUser, err := l.svcCtx.DB.User.FindOne(l.ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xerr.NotFoundErr.SetMessage("账号不存在")
		} else {
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
	}

	userInfo := user.UserInfo{
		UserId:    findUser.Id,
		Avatar:    findUser.Avatar,
		Nickname:  findUser.Nickname,
		Account:   findUser.Account,
		Bio:       &findUser.Bio,
		Gender:    findUser.Gender,
		Region:    findUser.Region,
		CreatedAt: findUser.CreatedAt,
	}
	if findUser.Status == 2 {
		userInfo.Nickname = "已注销"
	}

	return &user.GetUsersResp{
		UserInfo: &userInfo,
	}, nil
}
