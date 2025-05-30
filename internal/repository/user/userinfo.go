package user

import (
	"context"
	"encoding/json"
	"github.com/Haoyunforever/Micro-Scanner/internal/repository/cache"
	"github.com/redis/go-redis/v9"
	"time"
)

var userKey string = "user_key"

type UserInfo struct {
	Id uint `json:"id"`
}

func InitUserInfo(ctx context.Context, user *UserInfo) error {
	cachedUser, err := GetUserInfo(ctx)
	if err != nil {
		return err
	}

	// 如果缓存中存在用户信息，直接返回
	if cachedUser != nil {
		return nil
	}

	// 如果缓存中不存在用户信息，将用户信息存储到缓存中
	if err := SetUserInCache(ctx, user, cache.RedisClient); err != nil {
		return err
	}

	return nil
}

// 从Redis缓存中获取用户信息
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	data, err := cache.RedisClient.Get(ctx, userKey).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var user UserInfo
	//反序列化为UserInfo结构体
	if err := json.Unmarshal(data, &user); err != nil {
		// 解码数据时出错
		return nil, err
	}
	return &user, nil
}

// 将用户信息序列化后保存到Redis缓存中，设置过期时间为30s
func SetUserInCache(ctx context.Context, user *UserInfo, redisClient *redis.Client) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := redisClient.Set(ctx, userKey, data, time.Second*30).Err(); err != nil {
		return err
	}
	return nil
}
