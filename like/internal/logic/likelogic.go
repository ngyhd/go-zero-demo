package logic

import (
	"context"
	"errors"
	"go-zero-demo/like/internal/model"
	"go-zero-demo/like/internal/svc"
	"go-zero-demo/like/like"
	"go-zero-demo/pkg/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Like 点赞
func (l *LikeLogic) Like(in *like.LikeReq) (*like.LikeResp, error) {
	// 检查是否已经点赞
	existing, err := l.svcCtx.DB.Like.FindByUserAndPost(l.ctx, in.GetUserId(), in.GetPostId())
	if err == nil && existing != nil {
		return nil, xerr.New(400001, "已经点赞过了")
	}
	if !errors.Is(err, model.ErrNotFound) && err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	// 创建点赞记录
	likeRecord := &model.Like{
		UserId:    in.GetUserId(),
		PostId:    in.GetPostId(),
		CreatedAt: time.Now().Unix(),
	}
	_, err = l.svcCtx.DB.Like.Insert(l.ctx, likeRecord)
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	l.Infof("用户点赞成功: userId=%d, postId=%d", in.GetUserId(), in.GetPostId())
	return &like.LikeResp{}, nil
}
