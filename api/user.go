package api

import (
	"github.com/labstack/echo"
	"net/http"
	ctx "user-ratings/initialize/context"
	"user-ratings/modules/user"
)

const (
	// UserRatingRoute Info
	UserRatingRoute = "/user-rating"
)

type (
	UserRatingListing struct {
		UserId int `query:"userid" valid:"required"`
	}
)

//GetUserRatings fetches all ratings cast by a user
func GetUserRatings(echoCtx echo.Context) error {

	req := new(UserRatingListing)
	if err := echoCtx.Bind(req); err != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	if err := echoCtx.Validate(req); err != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}

	appCtx := echoCtx.Get("appCtx").(*ctx.AppContext)
	userRatings, err := user.GetUserRatingsResponse(req.UserId, appCtx)
	if err != nil {
		return echoCtx.JSON(http.StatusBadRequest, nil)
	}
	return echoCtx.JSON(http.StatusOK, userRatings)
}
