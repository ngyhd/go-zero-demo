package logic

import (
	"context"
	"go-zero-demo/like/internal/model"
	"go-zero-demo/like/internal/svc"
	"go-zero-demo/like/like"
	"go-zero-demo/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLikersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLikersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLikersLogic {
	return &GetLikersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLikers 获取推文点赞列表
func (l *GetLikersLogic) GetLikers(in *like.GetLikersReq) (*like.GetLikersResp, error) {
	l.Infof("获取推文点赞列表: postId=%d", in.GetPostId())

	// 查询点赞列表
	likes, err := l.svcCtx.DB.Like.FindByPost(l.ctx, in.GetPostId())
	if err != nil && err != model.ErrNotFound {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	// 构建响应
	infos := make([]*like.LikersInfo, 0, len(likes))
	for _, likeRecord := range likes {
		infos = append(infos, &like.LikersInfo{
			UserId:    likeRecord.UserId,
			PostId:    likeRecord.PostId,
			CreatedAt: likeRecord.CreatedAt,
			// Nickname 和 Avatar 需要通过 BFF 层调用 user 服务获取
			Nickname: "",
			Avatar:   "",
		})
	}

	l.Infof("获取推文点赞列表成功: postId=%d, count=%d", in.GetPostId(), len(infos))
	return &like.GetLikersResp{
		Infos: infos,
	}, nil
}
