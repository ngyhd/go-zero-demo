package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"testing"
)

func TestUpdateUserLogic_UpdateUser(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *user.UpdateUserReq
	}
	bio := "test bio"
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.UpdateUserResp
		wantErr bool
	}{
		{
			name: "更新用户信息",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.UpdateUserReq{
					UserInfo: &user.UserInfo{
						UserId:   12,
						Bio:      &bio,
						Nickname: "abc",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &UpdateUserLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.UpdateUser(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
