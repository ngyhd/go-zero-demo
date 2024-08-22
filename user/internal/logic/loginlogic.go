package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	account, err := l.svcCtx.DB.User.FindOneByAccount(l.ctx, in.GetAccount())
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return nil, xerr.NotFoundErr.SetMessage("账号不存在")
		} else {
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
	}
	if account.Status == 2 {
		return nil, xerr.NotFoundErr.SetMessage("账号不存在")
	}

	if account.Password != in.GetPassword() {
		return nil, xerr.AccountErr.SetMessage("密码错误")
	}
	// 如果是在注销中的用户，那么恢复该用户的注销
	if account.Status == 1 {
		account.Status = 0
		err = l.svcCtx.DB.User.Update(l.ctx, account)
		if err != nil {
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
	}
	return &user.LoginResp{}, nil
}
