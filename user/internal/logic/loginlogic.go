package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
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
			return nil, status.Error(400, "账号不存在")
		} else {
			return nil, status.Error(500, "系统错误")
		}
	}
	if account.Status == 2 {
		return nil, status.Error(400, "账号已注销")
	}

	if account.Password != in.GetPassword() {
		return nil, status.Error(400, "密码错误")
	}
	// 如果是在注销中的用户，那么恢复该用户的注销
	if account.Status == 1 {
		account.Status = 0
		err = l.svcCtx.DB.User.Update(l.ctx, account)
		if err != nil {
			return nil, status.Error(500, "系统错误")
		}
	}
	// TODO 登录态处理, 返回用户标识，并存储
	return &user.LoginResp{}, nil
}
