package logic

import (
	"context"
	"go-zero-demo/pkg/consts"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/post/model"
	"time"

	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	MaxTitleLength   = 255   // 标题最大长度
	MaxContentLength = 10000 // 内容最大长度
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
	l.Infof("发表推文请求: userId=%d", in.GetPostData().GetUserId())

	if in.GetPostData().GetUserId() == 0 {
		l.Infof("推文必须有作者")
		return nil, xerr.PostErr.SetMessage("推文必须有作者")
	}

	title := in.GetPostData().GetTitle()
	content := in.GetPostData().GetContent()

	if len(title) == 0 {
		l.Infof("标题不能为空")
		return nil, xerr.PostTitleErr
	}
	if len(title) > MaxTitleLength {
		l.Infof("标题超过最大长度限制: %d", len(title))
		return nil, xerr.PostTitleErr.SetMessage("标题不能超过255字符")
	}
	if len(content) == 0 {
		l.Infof("内容不能为空")
		return nil, xerr.PostContentErr
	}
	if len(content) > MaxContentLength {
		l.Infof("内容超过最大长度限制: %d", len(content))
		return nil, xerr.PostContentErr.SetMessage("内容不能超过10000字符")
	}

	data := model.Post{
		Title:     title,
		Content:   content,
		UserId:    in.GetPostData().GetUserId(),
		Status:    consts.PostStatusNormal,
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
		l.Errorf("数据库错误: %v", err)
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	l.Infof("推文发表成功: postId=%d, userId=%d", data.Id, data.UserId)
	return &post.CreatePostResp{}, nil
}
