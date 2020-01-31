package user

type (

	//Response defines the actual api response for userratings api
	Response struct {
		Name        string    `json:"name"`
		Gender      string    `json:"gender"`
		UserRatings []Ratings `json:"ratings"`
	}

	//User defines the user structure
	User struct {
		Name   string `bson:"name"`
		Gender string `bson:"gender"`
	}

	//Ratings defines the user's rating structure
	Ratings struct {
		MovieName string `bson:"movieName"`
		Rating    int    `bson:"rating"`
	}
)
