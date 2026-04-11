package logic

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go-zero-demo/pkg/regexp"
	"go-zero-demo/pkg/utils"
	"go-zero-demo/pkg/xerr"
	"go-zero-demo/user/model"
	"go-zero-demo/user/user"
	"time"

	"go-zero-demo/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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
	l.Info("用户注册请求", "account", utils.HashAccount(in.Account))

	// 账号格式校验
	match := regexp.Match(regexp.Account, in.GetAccount())
	if !match {
		l.Infof("账号格式不匹配: %s", utils.HashAccount(in.Account))
		return nil, xerr.AccountErr.SetMessage("账号格式不匹配，6-16位字符，且字首字符必须为字母")
	}
	// 密码格式校验
	match = regexp.Match(regexp.Pwd, in.GetPassword())
	if !match {
		l.Infof("密码格式不匹配: %s", utils.HashAccount(in.Account))
		return nil, xerr.AccountErr.SetMessage("账号格式不匹配，8-32位字符，且字必须包含大小写字母")
	}

	// 查询用户是否存在
	// 写入数据库
	// 密码哈希（使用 cost=12 提高安全性）
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), 12)
	if err != nil {
		l.Errorf("密码加密失败: %v", err)
		return nil, xerr.SystemErr.SetMessage("密码加密失败")
	}
	data := &model.User{
		Account:   in.Account,
		Password:  string(hashedPassword),
		CreatedAt: time.Now().Unix(),
	}
	insert, err := l.svcCtx.DB.User.Insert(l.ctx, data)
	if err != nil {
		var e *mysql.MySQLError
		if errors.As(err, &e) && e.Number == 1062 {
			l.Infof("账号已存在: %s", utils.HashAccount(in.Account))
			return nil, xerr.AccountErr.SetMessage("账号已经存在")
		}
		l.Errorf("数据库错误: %v", err)
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	id, err := insert.LastInsertId()
	if err != nil {
		l.Errorf("获取LastInsertId失败: %v", err)
		return nil, xerr.SystemErr.SetMessage(err.Error())
	}
	l.Infof("用户注册成功: userId=%d", id)
	return &user.RegisterResp{
		UserId: id,
	}, nil
}
