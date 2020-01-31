package context

import "user-ratings/db"

type (
	//AppContext is the application context
	AppContext struct {
		DB *db.DBCollections
	}
)
