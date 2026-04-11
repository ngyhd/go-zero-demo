package logic

import (
	"context"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchPostLogic {
	return &BatchPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获取推文
func (l *BatchPostLogic) BatchPost(in *post.BatchPostReq) (*post.BatchPostResp, error) {
	posts, err := l.svcCtx.DB.Post.GetPostByIds(l.ctx, in.GetPostIds())
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	return &post.BatchPostResp{
		Infos: ConvertPostsToDataList(posts),
	}, nil
}
