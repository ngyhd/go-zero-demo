package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/pkg/consts"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/user"
	"time"

	"go-zero-demo/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancellationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancellationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancellationLogic {
	return &CancellationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销
func (l *CancellationLogic) Cancellation(in *user.CancellationReq) (*user.CancellationResp, error) {
	l.Infof("用户注销请求: userId=%d", in.GetUserId())

	account, err := l.svcCtx.DB.User.FindOne(l.ctx, in.GetUserId())
	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			l.Infof("账号不存在: userId=%d", in.GetUserId())
			return nil, xerr.NotFoundErr.SetMessage("账号不存在")
		} else {
			l.Errorf("数据库错误: %v", err)
			return nil, xerr.SystemErr.SetMessage(err.Error())
		}
	}
	if account.Status == consts.UserStatusCancelled {
		l.Infof("账号已注销: userId=%d", in.GetUserId())
		return nil, xerr.UserDisabledErr.SetMessage("账号已注销")
	}
	account.Status = consts.UserStatusCancelled
	account.DeletedAt = time.Now().Unix()
	err = l.svcCtx.DB.User.Update(l.ctx, account)
	if err != nil {
		l.Errorf("更新用户状态失败: %v", err)
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	l.Infof("用户注销成功: userId=%d", in.GetUserId())
	return &user.CancellationResp{}, nil
}
