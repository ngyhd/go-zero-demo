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

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新推文
func (l *UpdatePostLogic) UpdatePost(in *post.UpdatePostReq) (*post.UpdatePostResp, error) {
	findedPost, err := l.svcCtx.DB.Post.FindOne(l.ctx, in.GetPostData().GetId())
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, status.Error(500, "系统错误")
	}
	if findedPost == nil {
		return nil, status.Error(400, "推文不存在")
	}

	if findedPost.UserId != in.GetPostData().GetUserId() {
		return nil, status.Error(400, "仅有作者可更新该推文")
	}

	findedPost.Title = in.GetPostData().GetTitle()
	findedPost.Content = in.GetPostData().GetContent()
	findedPost.UpdatedAt = time.Now().Unix()

	err = l.svcCtx.DB.Post.Update(l.ctx, findedPost)
	if err != nil {
		return nil, status.Error(500, "系统错误")
	}
	return &post.UpdatePostResp{}, nil
}
