package middleware

import (
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/xeipuuv/gojsonschema"

	"turntable-build/jsonschema"
	twerr "turntable-build/error"
)

func JSONSchemaHandler(schemaMapper map[string]interface{}) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			// Define schema
			apiSchema, ok := schemaMapper[c.Path()].(jsonschema.APISchema)
			if !ok {
				logrus.Error("Dose not exsit SchemaMapper: ", c.Path())
				jsonError := twerr.GetJSONError(fasthttp.StatusNotFound)
				return c.JSON(fasthttp.StatusNotFound, jsonError)
			}

			// Request validate check
			reqSchema := apiSchema.GetRequestSchema()
			reqContext := createRequestContext(c)
			if result := validate(reqSchema, reqContext); result != true {
				logrus.Error("Request validate error")
				jsonError := twerr.GetJSONError(fasthttp.StatusForbidden)
				return c.JSON(fasthttp.StatusForbidden, jsonError)
			}

			// Execute echo handler
			if err := next(c); err != nil {
				return err
			}

			// Response validate check
			resSchema := apiSchema.GetResponseSchema()
			resContext := c.Get("context")
			if result := validate(resSchema, resContext); result != true {
				logrus.Error("Response validate error")
				jsonError := twerr.GetJSONError(fasthttp.StatusInternalServerError)
				return c.JSON(fasthttp.StatusInternalServerError, jsonError)
			}

			return c.JSON(fasthttp.StatusOK, resContext)
		})
	}
}

func validate(schema interface{}, context interface{}) bool {
	schemaLoader := gojsonschema.NewGoLoader(schema)
	contextLoader := gojsonschema.NewGoLoader(context)

	result, err := gojsonschema.Validate(schemaLoader, contextLoader)
	if err != nil {
		logrus.Error("JsonSchema Validate Error:", err)
		panic(err.Error())
	}
	if !result.Valid() {
		logrus.Error("The document is not valid. see errors")
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			logrus.Error("- ", err)
		}
	}
	return result.Valid()
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
