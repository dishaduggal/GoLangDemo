package models

import (
	"bms-movies/app"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"
)

const (
	active   = "LIVE"
	inactive = "INACTIVE"
)

type (
	RequestMovieListing struct {
		VenueId string `query:"venueId" valid:"required"`
	}
	ResponseMovieListing struct {
		VenueName    string  `json:"venueName"`
		AllowParking bool    `json:"allowParking"`
		Shows        []Shows `json:"shows"`
	}
	CtxShowDetails struct {
		Wg       sync.WaitGroup
		Request  *RequestMovieListing
		Response ResponseMovieListing
		Error    chan error
	}
	Venues struct {
		Id           string `bson:"id"`
		Name         string `bson:"name"`
		AllowParking bool   `bson:"allowParking"`
	}
	ShowsData struct {
		ID    string `bson:"venueId"`
		Shows Shows  `bson:"shows"`
	}
	Shows struct {
		ID         string    `bson:"id"`
		ShowDate   time.Time `bson:"showDate"`
		FilmName   string    `bson:"filmName"`
		Categories []string  `bson:"showDate"`
	}

	Categories struct {
		ID         int    `bson:"id"`
		Name       string `bson:"name" json:"name"`
		Price      string `bson:"price" json:"price"`
		TotalSeats int    `bson:"totalSeats" json:"totalSeats"`
	}
)

//GetShowDetails fetches details of the offer
func GetShowDetails(req *RequestMovieListing) (ResponseMovieListing, error) {
	ctx := new(CtxShowDetails)
	ctx.Request = req
	ctx.GetCinemaDetails()
	return ctx.Response, nil
}

func (ctx *CtxShowDetails) GetCinemaDetails() error {

	venue := Venues{}
	queryVenues := bson.M{"id": ctx.Request.VenueId}
	err := app.Ctx.DB.Venues.Find(queryVenues).One(&venue)
	if err != nil {
		return err
	}
	ctx.Response.VenueName = venue.Name
	ctx.Response.AllowParking = venue.AllowParking
	return nil
}

func (ctx *CtxShowDetails) GetShowTimes(categories chan int) error {
	shows := Shows{}
	queryVenues := bson.M{"venueId": ctx.Request.VenueId}
	err := app.Ctx.DB.Venues.Find(queryVenues).One(&shows)
	if err != nil {
		return err
	}

	return nil
}

func (ctx *CtxShowDetails) GetShowCategories() {
}
