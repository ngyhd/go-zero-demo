package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/post/internal/config"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"
	"sync"
	"testing"
)

func InitConfig() config.Config {
	var once sync.Once
	var c config.Config
	once.Do(func() {
		var configFile = "post/etc/post.yaml"
		conf.MustLoad(configFile, &c)
	})

	return c
}

func TestCreatePostLogic_CreatePost(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *post.CreatePostReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "发表推文",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &post.CreatePostReq{
					PostData: &post.PostData{
						UserId:  111,
						Title:   "推文",
						Content: "第一条推文",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &CreatePostLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.CreatePost(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
