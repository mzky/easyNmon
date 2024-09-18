package router

import (
	"easyNmon/common"
	"easyNmon/controllers"
	"easyNmon/pkg"
	"github.com/bingoohuang/golog/pkg/echologrus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Flag common.Flag

func (f Flag) InitRouter() {
	r := echo.New()
	r.HideBanner = true
	r.Debug = f.Debug

	r.Use(echologrus.Logger(nil, true), middleware.Recover(), Cors())

	r.GET("/report", controllers.Generate)
	r.Any("/generate/:name/", controllers.Generate)
	r.GET("/sysInfo", controllers.GetSystemInfo)
	r.GET("/start", controllers.Start)
	r.GET("/close", controllers.Close)
	r.GET("/stop", controllers.Stop)

	r.StaticFS(common.WebRoot, pkg.StaticFS())
	r.GET("/", func(c echo.Context) error { return c.Redirect(http.StatusMovedPermanently, common.WebRoot) })

	r.Logger.Fatal(r.Start(":" + f.Port)) // listen

}

// Cors 支持跨域访问
func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE", "PATCH", "HEAD"},
		AllowHeaders: []string{"Authorization", "Content-Length", "X-CSRF-Token", "Accept", "Origin",
			"Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive",
			"User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma",
			"Access-Control-Allow-Origin", "X-Api-Applicationid", "Content-Disposition"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers",
			"Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar",
			"Content-Disposition"},
		AllowCredentials: true,
	})
}
