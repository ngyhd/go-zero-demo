package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/user"
	"time"

	"go-zero-demo/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	findOne, err := l.svcCtx.DB.User.FindOne(l.ctx, in.GetUserInfo().GetUserId())
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xerr.NotFoundErr.SetMessage("账号不存在")
		} else {
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
	}
	if in.GetUserInfo().GetAvatar() != "" {
		findOne.Avatar = in.GetUserInfo().GetAvatar()
	}
	if in.GetUserInfo().GetNickname() != "" {
		findOne.Avatar = in.GetUserInfo().GetNickname()
	}

	if in.GetUserInfo().Bio != nil {
		findOne.Bio = in.GetUserInfo().GetBio()
	}

	if in.GetUserInfo().GetGender() != 0 {
		findOne.Gender = in.GetUserInfo().GetGender()
	}
	if in.GetUserInfo().GetRegion() != "" {
		findOne.Region = in.GetUserInfo().GetRegion()
	}
	findOne.UpdatedAt = time.Now().Unix()

	err = l.svcCtx.DB.User.Update(l.ctx, findOne)
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	return &user.UpdateUserResp{}, nil
}
