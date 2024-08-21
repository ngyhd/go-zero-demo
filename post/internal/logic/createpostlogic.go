package logic

import (
	"context"
	"go-zero-demo/post/model"
	"google.golang.org/grpc/status"
	"time"

	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发表推文
func (l *CreatePostLogic) CreatePost(in *post.CreatePostReq) (*post.CreatePostResp, error) {
	if in.GetPostData().GetUserId() == 0 {
		return nil, status.Error(400, "推文必须有作者")
	}
	data := model.Post{
		Title:     in.GetPostData().GetTitle(),
		Content:   in.GetPostData().GetContent(),
		UserId:    in.GetPostData().GetUserId(),
		Status:    0,
		Views:     0,
		Likes:     0,
		Comments:  0,
		Shares:    0,
		Collects:  0,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
		DeletedAt: 0,
	}

	_, err := l.svcCtx.DB.Post.Insert(l.ctx, &data)
	if err != nil {
		return nil, status.Error(500, "系统错误")
	}

	return &post.CreatePostResp{}, nil
}
