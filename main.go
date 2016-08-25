package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/fujimisakari/turntable-build/route"
	"github.com/labstack/echo/engine/fasthttp"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	port := ":8880"
	fmt.Println("Start", port)

	router := route.Init()
	router.Run(fasthttp.New(port))
}
