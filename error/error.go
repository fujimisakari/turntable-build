package error

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func JSONHTTPErrorHandler(err error, c echo.Context) {
	code := fasthttp.StatusInternalServerError
	msg := "Internal Server Error"
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}
	if !c.Response().Committed() {
		c.JSON(code, map[string]interface{}{
			"statusCode": code,
			"message":    msg,
		})
	}
}

func GetJSONError(statusCode int) map[string]string {
	jsonError := map[string]string{
		"statusCode": strconv.Itoa(statusCode),
		"message":    fasthttp.StatusMessage(statusCode),
	}
	return jsonError
}
