package middleware

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/Sirupsen/logrus"
	twerr "github.com/fujimisakari/turntable-build/error"
	"github.com/fujimisakari/turntable-build/jsonschema"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/xeipuuv/gojsonschema"
)

func JSONSchemaHandler(schemaMapper map[string]interface{}) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			// Define schema
			urlMethod := strings.Replace(c.Path(), "/api/", "", 1)
			apiSchema, ok := schemaMapper[urlMethod].(jsonschema.APISchema)
			if !ok {
				logrus.Error("Dose not exsit SchemaMapper: ", c.Path())
				jsonError := twerr.GetJSONError(fasthttp.StatusNotFound)
				return c.JSON(fasthttp.StatusNotFound, jsonError)
			}

			// Request validate check
			reqSchema := apiSchema.GetRequestSchema()
			reqContext := createRequestContext(c)
			if result, msg := validate(reqSchema, reqContext); result != true {
				logrus.Error("Request validate error")
				fmt.Printf("%+v\n", msg)
				jsonError := twerr.GetJSONError(fasthttp.StatusForbidden)
				return c.JSON(fasthttp.StatusForbidden, jsonError)
			}

			// Execute echo handler
			if err := next(c); err != nil {
				jsonError := twerr.GetJSONError(fasthttp.StatusInternalServerError)
				jsonError["message"] = err.Error()
				fmt.Printf("%+v\n", err)
				return c.JSON(fasthttp.StatusInternalServerError, jsonError)
			}

			// Response validate check
			resSchema := apiSchema.GetResponseSchema()
			resContext := c.Get("context")
			if result, msg := validate(resSchema, resContext); result != true {
				logrus.Error("Response validate error")
				fmt.Printf("%+v\n", msg)
				jsonError := twerr.GetJSONError(fasthttp.StatusInternalServerError)
				return c.JSON(fasthttp.StatusInternalServerError, jsonError)
			}

			return c.JSON(fasthttp.StatusOK, resContext)
		})
	}
}

func validate(schema interface{}, context interface{}) (bool, string) {
	schemaLoader := gojsonschema.NewGoLoader(schema)
	contextLoader := gojsonschema.NewGoLoader(context)

	result, err := gojsonschema.Validate(schemaLoader, contextLoader)
	if err != nil {
		logrus.Error("JsonSchema Validate Error:", err)
	}

	var msg string
	if !result.Valid() {
		msg += "The Json is not valid. see errors"
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			msg += fmt.Sprintf("\n- %s", err)
		}
	}
	return result.Valid(), msg
}

func createRequestContext(c echo.Context) map[string]interface{} {
	mapList := make(map[string]interface{})

	switch c.Request().Method() {
	case echo.GET:
		for _, n := range c.ParamNames() {
			val, err := strconv.ParseInt(c.Param(n), 0, 64)
			if err != nil {
				mapList[n] = c.Param(n)
			} else {
				mapList[n] = val
			}
		}
	case echo.POST:
	case echo.PUT:
	case echo.DELETE:
	case echo.PATCH:
	case echo.OPTIONS:
	case echo.HEAD:
	case echo.CONNECT:
	case echo.TRACE:
	}
	return mapList
}
