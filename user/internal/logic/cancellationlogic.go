package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/user/user"
	"google.golang.org/grpc/status"
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
	account, err := l.svcCtx.DB.User.FindOne(l.ctx, in.GetUserId())
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
	account.Status = 1
	account.DeletedAt = time.Now().Unix()
	err = l.svcCtx.DB.User.Update(l.ctx, account)
	if err != nil {
		return nil, status.Error(500, "系统错误")
	}
	return &user.CancellationResp{}, nil
}
