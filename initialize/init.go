package initialize

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"user-ratings/api"
	"user-ratings/db"
	appCtx "user-ratings/initialize/context"

	"user-ratings/modules/utils"
)

const (
	//UserGroup Info
	UserGroup = "/user"
	port      = "8080"
)

// Init func creates all connections needed in the project
func Init() (err error) {
	appCtx := appCtx.AppContext{}

	// Initialize MongoDB
	appCtx.DB, err = db.InitializeMongo()
	if err != nil {
		return
	}

	//Any other external dependencies such as Config, Database or
	//external dependencies should be initialized here

	//Start server as dependencies are resolved
	StartServer(&appCtx)
	return nil
}

// StartServer method initializes the
// echo http server with defined routes
func StartServer(ctx *appCtx.AppContext) {
	appServer := echo.New()
	appServer.Use(BindApp(ctx))
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

//BindApp utilizes echo's middleware framework to bind our application's
//context - appCtx and use it subsequently for Dependency Injection during
//server initialization
func BindApp(ctx *appCtx.AppContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("appCtx", ctx)
			return next(c)
		}
	}
}
