package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user/internal/config"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"sync"
	"testing"
)

func InitConfig() config.Config {
	var once sync.Once
	var c config.Config
	once.Do(func() {
		var configFile = "user/etc/user.yaml"
		conf.MustLoad(configFile, &c)
	})

	return c
}

func TestRegisterLogic_Register(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *user.RegisterReq
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "错误账号注册",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.RegisterReq{
					Account:  "1abc3t3t",
					Password: "cccAAA@@@@",
				},
			},
			wantErr: true,
		},
		{
			name: "错误密码注册",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.RegisterReq{
					Account:  "abc3t3t",
					Password: "2222451251",
				},
			},
			wantErr: true,
		},
		{
			name: "正常注册",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.RegisterReq{
					Account:  "abc3t3t",
					Password: "Abc12345",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &RegisterLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
