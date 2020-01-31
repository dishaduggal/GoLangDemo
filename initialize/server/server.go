package server

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"user-ratings/api"
	ctx "user-ratings/initialize/context"

	"user-ratings/modules/utils"
)

const (
	//UserGroup Info
	UserGroup = "/user"
	port      = "8080"
)

// InitializeServer method initializes the
// echo http server with defined routes
func InitializeServer(ctx *ctx.AppContext) {
	appServer := echo.New()
	appServer.Use(bindApp(ctx))
	movieGroup := appServer.Group(UserGroup)
	appServer.Validator = &utils.CustomValidator{}
	movieGroup.GET(api.UserRatingRoute, api.GetUserRatings)

	httpServer := port
	err := appServer.Start(fmt.Sprintf(":%s", httpServer))
	if err != nil {
		log.Fatalf("Server could not start %s", err)
	}
	log.Print("Initialized HTTP server")
}

func bindApp(ctx *ctx.AppContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("appCtx", ctx)
			return next(c)
		}
	}
}
