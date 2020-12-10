package redis

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/wk331100/iFTY/app/libs/errorCode"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/system/helper"
)

const MASTER = "master"

type Redis struct {
	Connector *redis.Client
	ctx context.Context
}

func (this *Redis) Connect() (*Redis, interface{})  {
	return this.ConnectCluster(MASTER)
}

func (this *Redis) ConnectCluster(cluster string) (*Redis, interface{}) {
	redisConfig := config.RedisConfig
	if redisConfig[cluster] == nil {
		return nil,errorCode.ERR_CACHE
	}
	conf := redisConfig[cluster].(helper.Map)
	conn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d",conf["host"], conf["port"]),
		Password: conf["password"].(string),
		DB:       conf["db"].(int),
	})

	this.ctx = context.Background()

	pong, err := conn.Ping(this.ctx).Result()
	if err != nil {
		fmt.Println("redis 连接失败：", pong, err)
		return nil,errorCode.ERR_CACHE
	}

	this.Connector = conn
	return this,nil
}

func (this *Redis) Set(key string, value interface{}) (bool, interface{})  {
	err := this.Connector.Set(this.ctx, key, value, 0).Err()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return true,nil
}

func (this *Redis) Get(key string) (interface{}, interface{})  {
	val, err  := this.Connector.Get(this.ctx, key).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}