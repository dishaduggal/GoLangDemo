package initialize

import (
	"user-ratings/db"
	ctx "user-ratings/initialize/context"
	"user-ratings/initialize/server"
)

// Init func creates all connections needed in the project
func Init() (err error) {
	appCtx := ctx.AppContext{}

	// Initialize MongoDB
	appCtx.DB, err = db.InitializeMongo()
	if err != nil {
		return
	}

	//Any other external dependencies such as Config, Database or
	//external dependencies should be initialized here

	server.InitializeServer(&appCtx)
	return nil
}
