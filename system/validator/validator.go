package validator

import (
	"fmt"
	errors "github.com/wk331100/iFTY/system/error"
	"github.com/wk331100/iFTY/system/helper"
	"strconv"
	"strings"
)

const(
	REQUIRED 	= "required"
	FILLED		= "filled"

	BOOL		= "bool"
	INT			= "int"
	DIGITAL		= "digital"
	STRING	    = "string"
	ARRAY		= "array"
	BETWEEN		= "between"
	LARGE	    = "large"
	LESS		= "less"

	EMAIL		= "email"
	PHONE		= "phone"

)


type validator struct {
	Failure []string
}

func (v *validator) Fails() bool {
	fmt.Println("fails: ", v.Failure)
	if len(v.Failure) > 0 {
		return true
	}
	return false
}

func Make(params,condition helper.Map) (*validator,errors.Code) {
	fmt.Println(condition)

	if len(condition) <= 0 {
		return nil,errors.EMPTY_CONDITION
	}
	validator := new(validator)
	for key,item := range condition{
		cond := item.(string)

		condSlice := helper.Explode("|", cond)
		for _,condItem := range condSlice {
			if condItem == FILLED && params[key] == nil {
				break
			}

			if res,msg := match(params, key, condItem);  res == false {
				validator.Failure = append(validator.Failure, msg)
			}
		}
	}

	return validator, errors.EMPTY
	
}

func match(params helper.Map, key,cond string) (bool,string) {
	condKey := cond
	if strings.Index(cond, ":") != -1{
		condKey,_ = helper.List(":", cond)
	}
	switch condKey {
	case REQUIRED:
		if params[key] == nil {
			return false, errors.GetMessage(errors.VALIDATOR_REQUIRED, key)
		}
	case BETWEEN:
		_,region := helper.List(":", cond)
		minStr,maxStr := helper.List(",", region)
		min,_ := strconv.Atoi(minStr)
		max,_ := strconv.Atoi(maxStr)
		keyLen := len([]rune(params[key].(string)))
		if keyLen > max || keyLen < min {
			return false, errors.GetMessage(errors.VALIDATOR_NOT_BETWEEN, key)
		}
	}

	return true,""
}