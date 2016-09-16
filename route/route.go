package route

import (
	"github.com/fujimisakari/turntable-build/api"
	"github.com/fujimisakari/turntable-build/db"
	tberr "github.com/fujimisakari/turntable-build/error"
	"github.com/fujimisakari/turntable-build/jsonschema"
	tbMw "github.com/fujimisakari/turntable-build/middleware"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
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

	// Routes
	apiGroup := e.Group("/api")
	// Set Custom MiddleWare
	apiGroup.Use(tbMw.JSONSchemaHandler(jsonschema.GetSchemaMapper()))
	apiGroup.Use(tbMw.TransactionHandler(db.Init()))
	{
		apiGroup.GET("/artiste", api.GetArtisteAll())
		apiGroup.GET("/artiste/:id", api.GetArtiste())
	}

	// API Document Routes
	apiDocGroup := e.Group("/api-doc")
	{
		apiDocGroup.GET("/url", api.GetAllURL())
		apiDocGroup.GET("/schema", api.GetSchema())
	}

	return e
}
