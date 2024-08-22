package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/pkg/xerr"
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
	findPost, err := l.svcCtx.DB.Post.FindOne(l.ctx, in.GetPostData().GetId())
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	if findPost == nil {
		return nil, xerr.NotFoundErr.SetMessage("推文不存在")
	}

	if findPost.UserId != in.GetPostData().GetUserId() {
		return nil, xerr.PostErr.SetMessage("仅有作者可更新该推文")
	}

	findPost.Title = in.GetPostData().GetTitle()
	findPost.Content = in.GetPostData().GetContent()
	findPost.UpdatedAt = time.Now().Unix()

	err = l.svcCtx.DB.Post.Update(l.ctx, findPost)
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	return &post.UpdatePostResp{}, nil
}
