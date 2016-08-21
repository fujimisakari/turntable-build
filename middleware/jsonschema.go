package middleware

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/xeipuuv/gojsonschema"

	"turntable-build/jsonschema"
)

func JsonschemaHandler(schemaMapper map[string]interface{}) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			// Define schema
			apiSchema, ok := schemaMapper[c.Path()].(jsonschema.APISchema)
			if !ok {
				logrus.Debug("aaerror--------", c.Path())
			}

			// Request validate check
			reqSchema := apiSchema.GetRequestSchema()
			reqContext := createRequestContext(c)
			if err := validate(reqSchema, reqContext); err != nil {
				logrus.Debug("Request validate error", err)
				return err
			}

			// Execute echo handler
			if err := next(c); err != nil {
				return err
			}

			// Response validate check
			resSchema := apiSchema.GetResponseSchema()
			resContext := c.Get("context")
			if err := validate(resSchema, resContext); err != nil {
				logrus.Debug("Response validate error", err)
				return err
			}

			return c.JSON(fasthttp.StatusOK, resContext)
		})
	}
}

func validate(schema interface{}, context interface{}) error {
	schemaLoader := gojsonschema.NewGoLoader(schema)
	contextLoader := gojsonschema.NewGoLoader(context)

	result, err := gojsonschema.Validate(schemaLoader, contextLoader)
	if err != nil {
		logrus.Debug("JsonSchema Validate Error:", err)
		panic(err.Error())
	}
	if !result.Valid() {
		logrus.Debug("The document is not valid. see errors")
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			logrus.Debug("- ", err)
		}
		return nil
	}
	return nil
}

func createRequestContext(c echo.Context) map[string]interface{} {
	mapList := make(map[string]interface{})

	switch c.Request().Method() {
	case echo.GET:
		for _, n := range c.ParamNames() {
			mapList[n] = c.Param(n)
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
