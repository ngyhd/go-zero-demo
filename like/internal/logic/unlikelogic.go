package logic

import (
	"context"
	"go-zero-demo/like/internal/svc"
	"go-zero-demo/like/like"
	"go-zero-demo/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnlikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnlikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnlikeLogic {
	return &UnlikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Unlike 取消点赞
func (l *UnlikeLogic) Unlike(in *like.UnlikeReq) (*like.UnlikeResp, error) {
	err := l.svcCtx.DB.Like.DeleteByUserAndPost(l.ctx, in.GetUserId(), in.GetPostId())
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	l.Infof("用户取消点赞成功: userId=%d, postId=%d", in.GetUserId(), in.GetPostId())
	return &like.UnlikeResp{}, nil
}
