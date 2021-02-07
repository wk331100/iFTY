package services

import (
	"fmt"
	"github.com/wk331100/iFTY/app/libs/errorCode"
	"github.com/wk331100/iFTY/app/models"
	"github.com/wk331100/iFTY/system/helper"
	"github.com/wk331100/iFTY/system/redis"
)

type TestService struct {}

func (this *TestService) Create(data helper.Map) (int, interface{}) {
	return new(models.TestModel).Instance().Insert(data), nil
}

func (this *TestService) Update(data, where helper.Map) (int, interface{}) {
	return new(models.TestModel).Instance().Update(data, where), nil
}

func (this *TestService) List(where helper.Map) ([]helper.Map, interface{}) {
	return new(models.TestModel).Instance().List(where), nil
}

func (this *TestService) Info(where helper.Map) (helper.Map , interface{}) {
	redis,err := new(redis.Redis).Connect()
	if err != nil {
		return nil, err
	}

	redis.Set("a", 444444)
	redis.HSet("b", map[string]interface{}{"key1" : "123", "key2" : "222", "key3" : "333"})
	redis.HGet("b", "key1")
	redis.MSet(map[string]interface{}{"key1" : "111", "key2" : "222"})
	redis.MGet([]string{"key1", "key2"})
	redis.Incr("d")
	redis.Decr("a")
	redis.SetNX("e", 1 , -1)
	redis.MSetNX(map[string]interface{}{"f":1, "g":1})
	redis.Del("d")
	redis.HDel("b",[]string{"key1","key3"})
	redis.Exists([]string{"a","b"})
	redis.HExists("b", "key2")
	redis.HKeys("b")
	redis.HLen("b")
	redis.LPush("queue1", []interface{}{"aaaaaaaaa","eeeeeee"})
	redis.RPush("queue2", []interface{}{"bbbbbbbbbb","dddddddddd"})
	redis.RPop("queue1")
	redis.LPop("queue2")

	a, err := redis.Get("a")
	if err != nil {
		return nil, err
	}

	fmt.Println(a)





	return new(models.TestModel).Instance().Info(where), nil
}


func (this *TestService) Delete(filter helper.Map) (bool, interface{}) {
	testModel := new(models.TestModel)
	//查询是否存在
	if !testModel.Exist(filter) {
		return false, errorCode.ERR_NOT_EXIST
	}
	//执行删除
	return testModel.Instance().Delete(filter), nil
}
