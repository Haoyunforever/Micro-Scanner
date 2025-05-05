package service

import (
	"context"
	"github.com/Haoyunforever/Micro-Scanner/idl/pb"
	"github.com/Haoyunforever/Micro-Scanner/internal/model"
	user2 "github.com/Haoyunforever/Micro-Scanner/internal/repository/user"
	"github.com/Haoyunforever/Micro-Scanner/pkg/code"
	"gorm.io/gorm"
	"sync"
)

type UserSrv struct {
}

var UserSrvIns *UserSrv
var UserSrvOnce sync.Once

// 单例模式,确保只有一次实例执行
func NewUserSrv() *UserSrv {
	UserSrvOnce.Do(func() {
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

func (u *UserSrv) UserLogin(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	resp.Code = code.SUCCESS
	user, error := user2.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if error != nil {
		resp.Code = code.ERROR_NOT_EXIST_USER
		return
	}
	if !user.CheckPassword(req.Password) {
		resp.Code = code.ERROR_NOT_COMPARE
		return
	}
	resp.UserDetail = BuildUser(user)
	return
}

func (u *UserSrv) UserRegister(ctx context.Context, req *pb.UserRequest, resp *pb.UserDetailResponse) (err error) {
	if req.Password != req.PasswordConfirm {
		resp.Code = code.ERROR_NOT_COMPARE
		return
	}
	resp.Code = code.SUCCESS
	_, err := user2.NewUserDao(ctx).FindUserByUserName(req.UserName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//
		} else {
			resp.Code = code.ERROR
			return
		}
	}
	user := &model.User{
		UserName: req.UserName,
	}
	//加密密码
	if err = user.SetPassword(req.Password); err != nil {
		resp.Code = code.ERROR
		return
	}
	if err = user2.NewUserDao(ctx).CreateUser(user); err != nil {
		resp.Code = code.ERROR
		return
	}
	resp.UserDetail = BuildUser(user)
	return
}

func BuildUser(item *model.User) *pb.UserModel {
	userModel := pb.UserModel{
		Id:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}
