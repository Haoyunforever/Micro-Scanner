package user

import (
	"context"
	"github.com/Haoyunforever/Micro-Scanner/internal/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewDBClient(ctx)}
}

func (user *UserDao) CreateUser(in *model.User) error {
	return user.Model(&model.User{}).Create(&in).Error

}

func (user *UserDao) FindUserByUserName(userName string) (r *model.User, err error) {
	err = user.Model(&model.User{}).Where("user_name=?", userName).First(&r).Error
	if err != nil {
		return
	}
	return
}
