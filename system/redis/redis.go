package redis

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/wk331100/iFTY/app/libs/errorCode"
	"github.com/wk331100/iFTY/config"
	"github.com/wk331100/iFTY/system/helper"
	"time"
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

func (this *Redis) SetEX(key string, value interface{}, ttl time.Duration) (bool, interface{})  {
	err := this.Connector.SetEX(this.ctx, key, value, ttl).Err()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return true,nil
}

func (this *Redis) Incr(key string) (bool, interface{})  {
	err := this.Connector.Incr(this.ctx, key).Err()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return true,nil
}

func (this *Redis) Decr(key string) (bool, interface{})  {
	err := this.Connector.Decr(this.ctx, key).Err()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return true,nil
}

func (this *Redis) MSet(keys map[string]interface{}) (bool, interface{})  {
	err := this.Connector.MSet(this.ctx, keys).Err()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return true,nil
}

func (this *Redis) MGet(keys []string) (interface{}, interface{})  {
	val, err  := this.Connector.MGet(this.ctx, keys...).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) HSet(key string, values map[string]interface{}) (bool, interface{})  {
	err := this.Connector.HSet(this.ctx, key, values).Err()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return true,nil
}

func (this *Redis) HGet(key, field string) (interface{}, interface{})  {
	val, err  := this.Connector.HGet(this.ctx, key, field).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) SetNX(key string, value interface{}, ttl time.Duration) (interface{}, interface{})  {
	val, err  := this.Connector.SetNX(this.ctx, key, value, ttl).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}


func (this *Redis) MSetNX(values map[string]interface{}) (interface{}, interface{})  {
	val, err  := this.Connector.MSetNX(this.ctx, values).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) Del(key string) (interface{}, interface{})  {
	val, err  := this.Connector.Del(this.ctx, key).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) HDel(key string, fields []string) (interface{}, interface{})  {
	val, err  := this.Connector.HDel(this.ctx, key, fields...).Result()
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) Exists(key []string) (interface{}, interface{})  {
	val, err  := this.Connector.Exists(this.ctx, key...).Result()
	fmt.Println("EXIST: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) HExists(key string, field string) (interface{}, interface{})  {
	val, err  := this.Connector.HExists(this.ctx, key, field).Result()
	fmt.Println("HEXIST: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) HKeys(key string) (interface{}, interface{})  {
	val, err  := this.Connector.HKeys(this.ctx, key).Result()
	fmt.Println("HKEYS: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) HLen(key string) (interface{}, interface{})  {
	val, err  := this.Connector.HLen(this.ctx, key).Result()
	fmt.Println("HLEN: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) LPush(key string, values []interface{}) (interface{}, interface{})  {
	val, err  := this.Connector.LPush(this.ctx, key, values).Result()
	fmt.Println("LPush: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) LPop(key string) (interface{}, interface{})  {
	val, err  := this.Connector.LPop(this.ctx, key).Result()
	fmt.Println("LPop: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) RPush(key string, values []interface{}) (interface{}, interface{})  {
	val, err  := this.Connector.RPush(this.ctx, key, values).Result()
	fmt.Println("RPush: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) RPop(key string) (interface{}, interface{})  {
	val, err  := this.Connector.RPop(this.ctx, key).Result()
	fmt.Println("RPop: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}

func (this *Redis) Exec(command,key,value interface{}) (interface{}, interface{})  {
	val, err  := this.Connector.Do(this.ctx, command, key, value).Result()
	fmt.Println("Exec: ", val)
	if err != nil {
		return false,errorCode.ERR_CACHE
	}
	return val,nil
}