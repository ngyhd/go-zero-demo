// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"go-zero-demo/user/internal/logic"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 注册
func (s *UserServer) Register(ctx context.Context, in *user.RegisterReq) (*user.RegisterResp, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

// 登录
func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 注销
func (s *UserServer) Cancellation(ctx context.Context, in *user.CancellationReq) (*user.CancellationResp, error) {
	l := logic.NewCancellationLogic(ctx, s.svcCtx)
	return l.Cancellation(in)
}

// 查用户信息
func (s *UserServer) GetUsers(ctx context.Context, in *user.GetUsersReq) (*user.GetUsersResp, error) {
	l := logic.NewGetUsersLogic(ctx, s.svcCtx)
	return l.GetUsers(in)
}

// 更新用户信息
func (s *UserServer) UpdateUser(ctx context.Context, in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}