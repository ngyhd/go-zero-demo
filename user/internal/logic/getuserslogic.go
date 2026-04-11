package logic

import (
	"context"
	"go-zero-demo/pkg/consts"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/internal/svc"
	"go-zero-demo/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查用户信息
func (l *GetUsersLogic) GetUsers(in *user.GetUsersReq) (*user.GetUsersResp, error) {
	// 批量查询用户
	users, err := l.svcCtx.DB.User.FindByIds(l.ctx, in.GetUserIds())
	if err != nil {
		l.Errorf("批量查询用户失败: %v", err)
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}

	// 构建用户ID到用户信息的映射
	userMap := make(map[int64]*user.UserInfo)
	for _, u := range users {
		userInfo := &user.UserInfo{
			UserId:    u.Id,
			Avatar:    u.Avatar,
			Nickname:  u.Nickname,
			Account:   u.Account,
			Bio:       &u.Bio,
			Gender:    u.Gender,
			Region:    u.Region,
			CreatedAt: u.CreatedAt,
		}
		if u.Status == consts.UserStatusDisabled {
			userInfo.Nickname = "已注销"
		}
		userMap[u.Id] = userInfo
	}

	// 按照请求顺序返回用户信息
	respUsers := make([]*user.UserInfo, 0, len(in.GetUserIds()))
	for _, uid := range in.GetUserIds() {
		if info, ok := userMap[uid]; ok {
			respUsers = append(respUsers, info)
		}
	}

	return &user.GetUsersResp{
		UserInfo: respUsers,
	}, nil
}
