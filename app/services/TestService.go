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
	return new(models.TestModel).Insert(data), nil
}

func (this *TestService) Update(data, where helper.Map) (int, interface{}) {
	return new(models.TestModel).Update(data, where), nil
}

func (this *TestService) List(where helper.Map) ([]helper.Map, interface{}) {
	return new(models.TestModel).List(where), nil
}

func (this *TestService) Info(where helper.Map) (helper.Map , interface{}) {
	redis,err := new(redis.Redis).Connect()
	if err != nil {
		return nil, err
	}

	redis.Set("a", 123456)

	a, err := redis.Get("a")
	if err != nil {
		return nil, err
	}

	fmt.Println(a)





	return new(models.TestModel).Info(where), nil
}


func (this *TestService) Delete(filter helper.Map) (bool, interface{}) {
	testModel := new(models.TestModel)
	//查询是否存在
	if !testModel.Exist(filter) {
		return false, errorCode.ERR_NOT_EXIST
	}
	//执行删除
	return testModel.Delete(filter), nil
}
