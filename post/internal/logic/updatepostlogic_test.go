package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/post/internal/svc"
	"go-zero-demo/post/post"
	"testing"
)

func TestUpdatePostLogic_UpdatePost(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *post.UpdatePostReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "更新推文",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &post.UpdatePostReq{
					PostData: &post.PostData{
						Id:      2,
						UserId:  111,
						Title:   "推文UP",
						Content: "第一条推文UP",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &UpdatePostLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.UpdatePost(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
