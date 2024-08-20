package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
	"testing"
)

func TestCancellationLogic_Cancellation(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *user.CancellationReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *user.CancellationResp
		wantErr bool
	}{
		{
			name: "注销账户",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: svc.NewServiceContext(InitConfig()),
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &user.CancellationReq{
					UserId: 12,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &CancellationLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			_, err := l.Cancellation(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cancellation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
