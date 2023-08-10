package response

import "github.com/otumian-empire/go-url-shortener/util"

type T map[string]any

type response struct {
	status int
	others T
}

func new(statusCode int, others T) *response {
	return &response{
		status: statusCode,
		others: others,
	}
}

func (r *response) send() (int, T) {
	return r.status, r.others
}

func SuccessResponse(message string, data any) (int, T) {
	return new(util.SUCCESS_CODE, T{
		"success": true,
		"message": message,
		"data":    data,
	}).send()
}

func SuccessMessageResponse(message string) (int, T) {
	return new(util.SUCCESS_CODE, T{
		"success": true,
		"message": message,
		"data":    nil,
	}).send()
}

func FailureMessageResponse(message string) (int, T) {
	return new(util.SUCCESS_CODE, T{
		"success": false,
		"message": message,
		"data":    nil,
	}).send()
}

func ErrorResponse(message string) (int, T) {
	return new(util.INTERNAL_ERROR_CODE, T{
		"success": false,
		"message": message,
		"data":    nil,
	}).send()
}
