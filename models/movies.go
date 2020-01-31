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
		VenueName    string          `json:"venueName"`
		AllowParking bool            `json:"allowParking"`
		Shows        []ResponseShows `json:"shows"`
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

	ResponseShows struct {
		ShowDate   time.Time    `json:"showDate"`
		FilmName   string       `json:"movieName"`
		Categories []Categories `json:"categories"`
	}

	Shows struct {
		ID         string    `bson:"id"`
		ShowDate   time.Time `bson:"showDate"`
		FilmName   string    `bson:"filmName"`
		Categories []int     `bson:"showDate"`
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
		ctx.Error <- err
		return err
	}
	ctx.Response.VenueName = venue.Name
	ctx.Response.AllowParking = venue.AllowParking
	return nil
}

func (ctx *CtxShowDetails) GetShowTimes(categoryIDS chan int) error {
	shows := []Shows{}
	queryShows := bson.M{"venueId": ctx.Request.VenueId}
	err := app.Ctx.DB.Venues.Find(queryShows).One(&shows)
	if err != nil {
		ctx.Error <- err
		return err
	}
	distinctCategoryIDs := make(map[int]struct{})
	for _, show := range shows {
		for _, id := range show.Categories {
			_, ok := distinctCategoryIDs[id]
			if !ok {
				distinctCategoryIDs[id] = struct{}{}
				categoryIDS <- id
			}
		}
	}
	close(categoryIDS)
	// push category id in channel
	return nil
}
//GetCategories fetches category details
// sent in the channel categoryIDs
func (ctx *CtxShowDetails) GetCategories(categoryIDS chan int) error {
	categories := make(map[int]Categories)
	for {
		id, ok := <-categoryIDS
		if !ok {
			break
		}
		category := Categories{}
		queryCategory := bson.M{"id": id}
		err := app.Ctx.DB.Venues.Find(queryCategory).One(&category)
		if err != nil {
			ctx.Error <- err
			return err
		}
		categories[category.ID] = category
	}
	return nil
}
