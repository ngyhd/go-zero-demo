package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-demo/pkg/consts"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/post/model"
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
func (l *DeletePostLogic) DeletePost(in *post.DeletePostReq) (*post.DeletePostResp, error) {
	findOne, err := l.svcCtx.DB.Post.FindOneWithoutStatus(l.ctx, in.GetPostId())
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	if findOne == nil {
		return nil, xerr.NotFoundErr.SetMessage("推文不存在")
	}
	findOne.Status = consts.PostStatusDeleted
	findOne.DeletedAt = time.Now().Unix()
	err = l.svcCtx.DB.Post.Update(l.ctx, findOne)
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	l.Infof("删除推文成功: postId=%d", in.GetPostId())
	return &post.DeletePostResp{}, nil
}
