package errors

import "fmt"

const (
	EMPTY  = 0

	INSERT_ERROR = 10
	UPDATE_ERROR = 11
	QUERY_ERROR  = 12
	DELETE_ERROR = 13

	EMPTY_DATA   = 20
	REDIS_ERROR  = 21
	EMPTY_CONDITION	 = 22

	VALIDATOR_REQUIRED = 50
	VALIDATOR_TYPE_ERR = 51
	VALIDATOR_NOT_BETWEEN = 52


)

type Code int

var Message = map[int]string{
	INSERT_ERROR 		: "Insert Error : %v",
	UPDATE_ERROR 		: "Update Error : %v",
	QUERY_ERROR 		: "Query Error : %v",
	DELETE_ERROR 		: "Delete Error : %v",

	EMPTY_DATA			: "Empty Data",
	REDIS_ERROR			: "Redis Error: %v",
	EMPTY_CONDITION		: "Condition Error",

	VALIDATOR_REQUIRED 	: "%s Is Required",
	VALIDATOR_TYPE_ERR  : "%s Type Is Not Match",
	VALIDATOR_NOT_BETWEEN : "%s Not Between The Region",
}

func GetMessage(code int, error string) string {
	if Message[code] != "" {
		return fmt.Sprintf(Message[code], error)
	}
	return ""
}