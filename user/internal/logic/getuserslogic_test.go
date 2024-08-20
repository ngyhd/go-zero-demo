package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"testing"
)

func TestGetUsersLogic_GetUsers(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *user.GetUsersReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.GetUsersResp
		wantErr bool
	}{
		{
			name: "正常查找用户",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.GetUsersReq{
					UserId: 12,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &GetUsersLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.GetUsers(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
