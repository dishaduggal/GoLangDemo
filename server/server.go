package server

import (
	"bms-movies/controllers"
	"fmt"
	"github.com/labstack/echo"
	"log"
)

const (
	//MovieGroup Info
	MovieGroup = "/movies"
)

// InitializeServer method initializes the
// echo http server with defined routes
func InitializeServer() {
	appServer := echo.New()
	appServer.HideBanner = true
	movieGroup := appServer.Group(MovieGroup)
	appServer.Validator = &controllers.CustomValidator{}
	movieGroup.GET(controllers.MovieVenueListingRoute, controllers.GetShowTimes)

	httpServer := "8080"
	err := appServer.Start(fmt.Sprintf(":%s", httpServer))
	if err != nil {
		log.Fatalf("Server could not start %s", err)
	}
	log.Print("Initialized HTTP server")
}
