package app

type (
	//Context is the application context
	Context struct {
		DB *DBCollections
	}
)

// Ctx is app-context variable that contains golabal connections to various external libraries.
// It can be imported and used throughout the project.
var Ctx Context

// Initialize func creates all connections needed in the project
func Initialize() (err error) {

	// Initialize MongoDB
	Ctx.DB, err = initializeMongo()
	if err != nil {
		return
	}
	return nil
}
