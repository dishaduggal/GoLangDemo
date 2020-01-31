package initialize

import (
	"user-ratings/db"
	ctx "user-ratings/initialize/context"
)

//MockInit will have all mocks of dependencies
func MockInit() (err error) {

	ctx := ctx.AppContext{}
	// Init mock MongoDB
	ctx.DB = db.MockMdb()

	// Init mock config

	// Init hhtp handler
	return
}
