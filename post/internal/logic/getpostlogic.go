package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostLogic {
	return &GetPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单条推文
func (l *GetPostLogic) GetPost(in *post.GetPostReq) (*post.GetPostResp, error) {
	findOne, err := l.svcCtx.DB.Post.FindOne(l.ctx, in.GetPostId())
	if err != nil && !errors.Is(err, sqlx.ErrNotFound) {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	if findOne == nil || findOne.Status == 1 {
		return nil, xerr.NotFoundErr.SetMessage("推文不存在")
	}
	return &post.GetPostResp{
		Info: &post.PostData{
			Id:       findOne.Id,
			UserId:   findOne.UserId,
			Title:    findOne.Title,
			Content:  findOne.Content,
			Views:    findOne.Views,
			Likes:    findOne.Likes,
			Comments: findOne.Comments,
			Shares:   findOne.Shares,
			Collects: findOne.Collects,
		},
	}, nil
}
