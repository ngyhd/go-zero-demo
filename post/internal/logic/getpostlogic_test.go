package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"
	"testing"
)

func TestGetPostLogic_GetPost(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *post.GetPostReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *post.GetPostResp
		wantErr bool
	}{
		{
			name: "查看推文",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &post.GetPostReq{
					PostId: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GetPostLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.GetPost(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
