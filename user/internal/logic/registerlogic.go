package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go-zero-demo/pkg/regexp"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/model"
	"go-zero-demo/user/user"
	"time"

	"go-zero-demo/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 账号格式校验
	match := regexp.Match(regexp.Account, in.GetAccount())
	if !match {
		return nil, xerr.AccountErr.SetMessage("账号格式不匹配，6-16位字符，且字首字符必须为字母")
	}
	// 密码格式校验
	match = regexp.Match(regexp.Pwd, in.GetPassword())
	if !match {
		return nil, xerr.AccountErr.SetMessage("账号格式不匹配，8-32位字符，且字必须包含大小写字母")
	}

	// 查询用户是否存在
	// 写入数据库
	data := &model.User{
		Account:   in.Account,
		Password:  in.Password,
		CreatedAt: time.Now().Unix(),
	}
	insert, err := l.svcCtx.DB.User.Insert(l.ctx, data)
	if err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e) && e.Number == 1062 {
			return nil, xerr.AccountErr.SetMessage("账号已经存在")
		}
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	return &user.RegisterResp{
		UserId: id,
	}, nil
}
