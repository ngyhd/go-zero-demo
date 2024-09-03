package posts

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/post/post"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatLogic {
	return &CreatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatLogic) Creat(req *types.CreatePostReq) (resp *types.CreatePostResp, err error) {
	reqData := &post.CreatePostReq{
		PostData: &post.PostData{
			Title:    req.Title,
			Content:  req.Content,
			UserId:   gconv.Int64(req.UserId),
			Views:    0,
			Likes:    0,
			Comments: 0,
			Shares:   0,
			Collects: 0,
		},
	}
	_, err = l.svcCtx.PostRpc.CreatePost(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
