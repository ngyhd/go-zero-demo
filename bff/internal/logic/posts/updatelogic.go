package posts

import (
	"context"
	"go-zero-demo/post/post"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdatePostReq) (resp *types.UpdatePostResp, err error) {
	reqData := &post.UpdatePostReq{
		PostData: &post.PostData{
			Title:   req.Title,
			Content: req.Content,
		},
	}
	_, err = l.svcCtx.PostRpc.UpdatePost(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
