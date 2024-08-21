package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"
	"testing"
)

func TestDeletePostLogic_DeletePost(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *post.DeletePostReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *post.DeletePostReq
		wantErr bool
	}{
		{
			name: "删除推文",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &post.DeletePostReq{
					PostId: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DeletePostLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.DeletePost(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
