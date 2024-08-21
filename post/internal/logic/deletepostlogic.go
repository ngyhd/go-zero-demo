package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"time"

	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除推文
func (l *DeletePostLogic) DeletePost(in *post.DeletePostReq) (*post.DeletePostReq, error) {
	findOne, err := l.svcCtx.DB.Post.FindOne(l.ctx, in.GetPostId())
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, status.Error(500, "系统错误")
	}
	if findOne == nil {
		return nil, status.Error(404, "推文不存在")
	}
	findOne.Status = 1
	findOne.DeletedAt = time.Now().Unix()
	err = l.svcCtx.DB.Post.Update(l.ctx, findOne)
	if err != nil {
		return nil, status.Error(500, "系统错误")
	}
	return &post.DeletePostReq{}, nil
}