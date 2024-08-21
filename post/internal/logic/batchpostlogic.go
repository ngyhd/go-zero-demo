package logic

import (
	"context"
	"google.golang.org/grpc/status"

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
		return nil, status.Error(500, "系统错误")
	}
	infos := make([]*post.PostData, 0)
	for _, p := range posts {
		info := &post.PostData{
			Id:       p.Id,
			UserId:   p.UserId,
			Title:    p.Title,
			Content:  p.Content,
			Views:    p.Views,
			Likes:    p.Likes,
			Comments: p.Comments,
			Shares:   p.Shares,
			Collects: p.Collects,
		}
		infos = append(infos, info)
	}

	return &post.BatchPostResp{
		Infos: infos,
	}, nil
}
