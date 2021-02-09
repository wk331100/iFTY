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



)

type Code int

var Message = map[int]string{
	INSERT_ERROR 		: "Insert Error : %v",
	UPDATE_ERROR 		: "Update Error : %v",
	QUERY_ERROR 		: "Query Error : %v",
	DELETE_ERROR 		: "Delete Error : %v",

	EMPTY_DATA			: "Empty Data",
	REDIS_ERROR			: "Redis Error: %v",
}

func GetMessage(code int, error string) string {
	if Message[code] != "" {
		return fmt.Sprintf(Message[code], error)
	}
	return ""
}