package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

const (
	users        = "users"
	ratings      = "ratings"
	databaseName = "users"
)

// DBCollections contains all collections of the database
type DBCollections struct {
	Users   Collection
	Ratings Collection
}

// GlobalSession for mongo connection
var (
	GlobalSession     *mgo.Session
	StrCommExpiryMins string
)

// initializeMongo func to connect to mongoDB
func InitializeMongo() (tbl *DBCollections, err error) {

	// create a global connection if it does not exist
	if GlobalSession == nil {
		GlobalSession, err = mgo.Dial("localhost")
		if err != nil {
			log.Fatalf("Could not connect to Database")
			return nil, err
		}
		_ = GlobalSession

	}

	//Simple example of an anonymous function which initializes
	//all collections in the users database
	initializeCollections := func() *DBCollections {
		mgoSession := GlobalSession.Copy()
		session := MongoSession{mgoSession}
		database := session.DB(databaseName)

		tbl := &DBCollections{}
		tbl.Users = database.C(users)
		tbl.Ratings = database.C(ratings)
		return tbl
	}
	return initializeCollections(), nil
}

// **** Session Wrapper Start ****

// Session Wrapper
type Session interface {
	DB(name string) Database
	Close()
}

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// DB shadows *mgo.DB to returns a Database interface instead of *mgo.Database.
func (s MongoSession) DB(name string) Database {
	return &MongoDatabase{Database: s.Session.DB(name)}
}

// **** Session Wrapper End ****

// **** Database Wrapper Start ****

// Database Wrapper
type Database interface {
	C(name string) Collection
}

// MongoDatabase shadows mgo.Database struct
type MongoDatabase struct {
	*mgo.Database
}

// C shadows *mgo.DB to returns a Database interface instead of *mgo.Database.
func (d MongoDatabase) C(name string) Collection {
	return &MongoCollection{Collection: d.Database.C(name)}
}

// **** Database Wrapper End ****

// **** Collection Wrapper Start ****

// Collection Wrapper
type Collection interface {
	Find(query interface{}) Query
}

// MongoCollection shadows mgo.Collection struct
type MongoCollection struct {
	*mgo.Collection
}

// Find shadows *mgo.collection.Find to returns a Database interface instead of *mgo.Database.
func (mgoCollection MongoCollection) Find(query interface{}) Query {
	return &MongoQuery{Query: mgoCollection.Collection.Find(query)}
}

// **** Collection Wrapper End ****

// **** Query Wrapper Start ****

// Query Wrapper
type Query interface {
	Select(selector interface{}) Query
	Sort(fields ...string) Query
	Skip(n int) Query
	All(result interface{}) error
	One(result interface{}) (err error)
}

// MongoQuery shadows mgo.Query struct
type MongoQuery struct {
	*mgo.Query
}

// One shadows *mgo.Query.One()
func (mgoQuery MongoQuery) One(result interface{}) (err error) {
	return mgoQuery.Query.One(result)
}

// Select shadows *mgo.Query.Select()
func (mgoQuery MongoQuery) Select(selector interface{}) Query {
	return &MongoQuery{Query: mgoQuery.Query.Select(selector)}
}

// Sort shadows *mgo.Query.Sort()
func (mgoQuery MongoQuery) Sort(fields ...string) Query {
	return &MongoQuery{Query: mgoQuery.Query.Sort(fields...)}
}

// Skip shadows *mgo.Query.Skip()
func (mgoQuery MongoQuery) Skip(n int) Query {
	return &MongoQuery{Query: mgoQuery.Query.Skip(n)}
}

// Limit shadows *mgo.Query.Limit()
func (mgoQuery MongoQuery) Limit(n int) Query {
	return &MongoQuery{Query: mgoQuery.Query.Limit(n)}
}

// Count shadows *mgo.Query.Count()
func (mgoQuery MongoQuery) Count() (int, error) {
	return mgoQuery.Query.Count()
}

// **** Query Wrapper End ****
