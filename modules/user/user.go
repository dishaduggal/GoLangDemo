package user

import (
	"gopkg.in/mgo.v2/bson"
	ctx "user-ratings/initialize/context"
)

type (
	UserCtx struct {
		ctx *ctx.AppContext
	}
)

//GetUserRatingsResponse fetches details of the offer
func GetUserRatingsResponse(userID int, appCtx *ctx.AppContext) (Response, error) {

	//pointer receiver
	userCtx := new(UserCtx)
	userCtx.ctx = appCtx

	// make
	userChannel := make(chan User)
	go userCtx.GetUserDetails(userID, userChannel)

	ratingsChannel := make(chan []Ratings)
	go userCtx.GetUserRatings(userID, ratingsChannel)

	user := <-userChannel
	ratings := <-ratingsChannel

	response := Response{
		Name:        user.Name,
		Gender:      user.Gender,
		UserRatings: ratings,
	}

	return response, nil
}

// GetUserDetails fetches details of a user from the database
// and publishes the result into a channel of type User
func (userCtx *UserCtx) GetUserDetails(userID int, userChannel chan User) error {
	user := User{}
	queryUsers := bson.M{"id": userID}
	err := userCtx.ctx.DB.Users.Find(queryUsers).One(&user)
	if err != nil {
		return err
	}
	userChannel <- user
	return nil
}

// GetUserRatings fetches ratings cast by the user from the database
// and publishes the result into a channel of type slice of Ratings
func (userCtx *UserCtx) GetUserRatings(userID int, ratingsChannel chan []Ratings) error {
	ratings := make(map[string][]Ratings)
	queryRatings := bson.M{"id": userID}
	selectRatings := bson.M{"ratings": 1}
	err := userCtx.ctx.DB.Ratings.Find(queryRatings).Select(selectRatings).One(&ratings)
	if err != nil {
		return err
	}
	ratingsChannel <- ratings["ratings"]
	return nil
}
