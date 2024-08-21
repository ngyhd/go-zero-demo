package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"
	"testing"
)

func TestBatchPostLogic_BatchPost(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *post.BatchPostReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *post.BatchPostResp
		wantErr bool
	}{
		{
			name: "批量获取推文",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &post.BatchPostReq{
					PostIds: []int64{2, 3},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &BatchPostLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.BatchPost(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("BatchPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
