package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"
	"testing"
)

func TestGetUserPostListLogic_GetUserPostList(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *post.GetUserPostListReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *post.GetUserPostListResp
		wantErr bool
	}{
		{
			name: "查看用户推文",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &post.GetUserPostListReq{
					UserId: 111,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GetUserPostListLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.GetUserPostList(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserPostList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
