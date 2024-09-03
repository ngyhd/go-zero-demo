package posts

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/post/post"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPostsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPostsLogic {
	return &GetUserPostsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPostsLogic) GetUserPosts(req *types.GetUserPostsReq) (resp *types.GetUserPostsResp, err error) {
	reqData := &post.GetUserPostListReq{
		UserId: gconv.Int64(req.UserId),
	}
	list, err := l.svcCtx.PostRpc.GetUserPostList(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	respPosts := []types.Post{}
	for _, d := range list.GetInfos() {
		postInfo := types.Post{
			Id:        gconv.String(d.GetId()),
			Title:     d.GetTitle(),
			Content:   d.GetContent(),
			UserId:    gconv.String(d.GetUserId()),
			Views:     d.GetViews(),
			Likes:     d.GetLikes(),
			Comments:  d.GetComments(),
			Shares:    d.GetShares(),
			Collects:  d.GetCollects(),
			CreatedAt: gconv.String(d.GetCreatedAt()),
			UpdatedAt: gconv.String(d.GetUpdatedAt()),
		}
		respPosts = append(respPosts, postInfo)
	}
	resp = &types.GetUserPostsResp{Posts: respPosts}
	return
}
