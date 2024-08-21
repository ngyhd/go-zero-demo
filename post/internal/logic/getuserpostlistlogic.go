package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostListLogic {
	return &GetUserPostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户用户推文列表
func (l *GetUserPostListLogic) GetUserPostList(in *post.GetUserPostListReq) (*post.GetUserPostListResp, error) {
	posts, err := l.svcCtx.DB.Post.GetPostByUser(l.ctx, in.GetUserId())
	if err != nil {
		return nil, status.Error(500, "系统错误")
	}
	infos := make([]*post.PostData, 0)
	for _, p := range posts {
		info := &post.PostData{
			Id:       p.Id,
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

	return &post.GetUserPostListResp{
		Infos: infos,
	}, nil
}
