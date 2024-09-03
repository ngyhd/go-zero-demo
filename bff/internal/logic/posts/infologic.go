package posts

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/post/post"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoReq) (resp *types.InfoResp, err error) {
	reqData := &post.GetPostReq{
		PostId: gconv.Int64(req.Id),
	}
	postResp, err := l.svcCtx.PostRpc.GetPost(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	d := postResp.GetInfo()
	resp = &types.InfoResp{
		Post: types.Post{
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
		},
	}
	return
}
