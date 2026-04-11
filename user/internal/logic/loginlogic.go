package logic

import (
	"context"
	"errors"
	"fmt"
	"go-zero-demo/pkg/auth"
	"go-zero-demo/pkg/consts"
	"go-zero-demo/pkg/utils"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
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
	l.Info("用户登录请求", "account", utils.HashAccount(in.Account))

	account, err := l.svcCtx.DB.User.FindOneByAccount(l.ctx, in.GetAccount())
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			l.Infof("账号不存在: %s", utils.HashAccount(in.Account))
			return nil, xerr.NotFoundErr.SetMessage("账号不存在")
		} else {
			l.Errorf("数据库错误: %v", err)
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
	}

	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(in.GetPassword())); err != nil {
		l.Infof("密码错误: %s", utils.HashAccount(in.Account))
		return nil, xerr.AccountErr.SetMessage("账号或密码错误")
	}

	// 如果是在注销中的用户，那么恢复该用户的注销
	if account.Status == consts.UserStatusCancelled {
		account.Status = consts.UserStatusNormal
		err = l.svcCtx.DB.User.Update(l.ctx, account)
		if err != nil {
			l.Errorf("恢复注销状态失败: %v", err)
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
		l.Infof("恢复用户注销状态: userId=%d", account.Id)
	}

	l.Infof("用户登录成功: userId=%d", account.Id)

	// 生成 JWT Token
	now := time.Now().Unix()
	token, err := auth.GetJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, l.svcCtx.Config.JwtAuth.AccessExpire, fmt.Sprintf("%d", account.Id))
	if err != nil {
		l.Errorf("生成 JWT Token 失败: %v", err)
		return nil, xerr.SystemErr.SetMessage("生成 Token 失败")
	}

	return &user.LoginResp{
		Token:  token,
		UserId: account.Id,
	}, nil
}
