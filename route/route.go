package route

import (
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"

	"turntable-build/api"
	"turntable-build/db"
	tberr "turntable-build/error"
	"turntable-build/jsonschema"
	tbMw "turntable-build/middleware"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Debug()

	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.SetHTTPErrorHandler(tberr.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	e.Use(tbMw.JSONSchemaHandler(jsonschema.GetSchemaMapper()))
	e.Use(tbMw.TransactionHandler(db.Init()))

	// Routes
	urlGroup := e.Group("/api")
	{
		urlGroup.GET("/artiste", api.GetArtisteAll())
		urlGroup.GET("/artiste/:id", api.GetArtiste())
	}
	return e
}
