package controllers

import (
	"bms-movies/models"
	"github.com/labstack/echo"
	"net/http"
)

const (
	//MovieVenueListingRoute Info
	MovieVenueListingRoute = "/venue-listing"
)

//GetMovieListing fetches all movies playing at a venue
func GetShowTimes(context echo.Context) error {

	var response interface{}

	req := new(models.RequestMovieListing)
	if err := context.Bind(req); err != nil {
		return context.JSON(http.StatusBadRequest, response)
	}

	if err := context.Validate(req); err != nil {
		return context.JSON(http.StatusBadRequest, response)
	}

	showDetails, err := models.GetShowDetails(req)
	if err != nil {
		return context.JSON(http.StatusBadRequest, response)
	}
	return context.JSON(http.StatusOK, showDetails)
}
