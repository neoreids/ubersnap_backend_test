package utilities

import "github.com/gofiber/fiber/v2"

type Error struct {
	StatusCode int
	Err        error
}

func (r *Error) Error() string {
	return r.Err.Error()
}

func ErrorRequest(err error, code int) error {
	return &Error{
		StatusCode: code,
		Err:        err,
	}
}

type ResponseRequest struct {
	Message *string
	Data    interface{}
	Meta    interface{}
	Code    *int
	Error   error
}

func Response(ctx *fiber.Ctx, request *ResponseRequest) error {
	resp := make(map[string]interface{})
	if request.Error != nil {
		errInfo, ok := request.Error.(*Error)
		if ok {
			resp["Code"] = errInfo.StatusCode
			resp["message"] = errInfo.Error()
			return ctx.Status(errInfo.StatusCode).JSON(resp)
		}
	}
	if request.Message != nil {
		resp["message"] = *request.Message
	} else {
		resp["message"] = "ok"
	}

	if request.Code == nil {
		statusCode := new(int)
		*statusCode = 200
		request.Code = statusCode
	}

	if request.Data != nil {
		resp["data"] = request.Data
	} else {
		resp["data"] = []fiber.Map{}
	}

	if request.Meta != nil {
		resp["meta"] = request.Meta
	}

	return ctx.Status(*request.Code).JSON(resp)
}
