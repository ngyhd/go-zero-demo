// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userclient

import (
	"context"

	"go-zero-demo/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CancellationReq  = user.CancellationReq
	CancellationResp = user.CancellationResp
	GetUsersReq      = user.GetUsersReq
	GetUsersResp     = user.GetUsersResp
	LoginReq         = user.LoginReq
	LoginResp        = user.LoginResp
	RegisterReq      = user.RegisterReq
	RegisterResp     = user.RegisterResp
	UpdateUserReq    = user.UpdateUserReq
	UpdateUserResp   = user.UpdateUserResp
	UserInfo         = user.UserInfo

	User interface {
		// 注册
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		// 登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 注销
		Cancellation(ctx context.Context, in *CancellationReq, opts ...grpc.CallOption) (*CancellationResp, error)
		// 查用户信息
		GetUsers(ctx context.Context, in *GetUsersReq, opts ...grpc.CallOption) (*GetUsersResp, error)
		// 更新用户信息
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

// 注册
func (m *defaultUser) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 登录
func (m *defaultUser) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 注销
func (m *defaultUser) Cancellation(ctx context.Context, in *CancellationReq, opts ...grpc.CallOption) (*CancellationResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Cancellation(ctx, in, opts...)
}

// 查用户信息
func (m *defaultUser) GetUsers(ctx context.Context, in *GetUsersReq, opts ...grpc.CallOption) (*GetUsersResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUsers(ctx, in, opts...)
}

// 更新用户信息
func (m *defaultUser) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}
