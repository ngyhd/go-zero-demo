package posts

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"go-zero-demo/post/post"

	"go-zero-demo/bff/internal/svc"
	"go-zero-demo/bff/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.DeleteReq) (resp *types.DeleteResp, err error) {
	reqData := &post.DeletePostReq{
		PostId: gconv.Int64(req.Id),
	}
	_, err = l.svcCtx.PostRpc.DeletePost(l.ctx, reqData)
	if err != nil {
		return nil, err
	}
	return
}
