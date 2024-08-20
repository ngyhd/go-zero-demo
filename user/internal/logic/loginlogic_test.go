package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"testing"
)

func TestLoginLogic_Login(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *user.LoginReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.LoginResp
		wantErr bool
	}{
		{
			name: "正常账号登录",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.LoginReq{
					Account:  "abc3t3t",
					Password: "Abc12345",
				},
			},
			wantErr: false,
		},
		{
			name: "错误账号登录",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.LoginReq{
					Account:  "abc3t3txxx",
					Password: "Abc12345",
				},
			},
			wantErr: true,
		},
		{
			name: "错误密码登录",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.LoginReq{
					Account:  "abc3t3t",
					Password: "Abc12345xxx",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LoginLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.Login(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
