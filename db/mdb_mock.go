package db

import (
	"gopkg.in/mgo.v2/bson"
)

func MockMdb() *DBCollections {
	mockSession := NewMockSession()
	mockdb := mockSession.DB("")

	mockTbl := &DBCollections{}
	mockTbl.Users = mockdb.C(users)
	mockTbl.Ratings = mockdb.C(ratings)

	return mockTbl
}

type (
	// MockSession satisfies Session and act as a mock of *mgo.session.
	MockSession struct{}

	// MockDatabase satisfies DataLayer and act as a mock.
	MockDatabase struct{}

	// MockCollection satisfies Collection and act as a mock.
	MockCollection struct {
		collection string
	}

	// MockQuery satisfies Query and act as a mock.
	MockQuery struct {
		collection string
		query      map[string]interface{}
	}

	// MockPipe satisfies Pipe and act as a mock.
	MockPipe struct {
		collection string
		pipe       []bson.M
	}
)

// NewMockSession mock NewSession.
func NewMockSession() Session {
	return MockSession{}
}

// ********* Mocking Session functions**********

// Close mocks mgo.Session.Close().
func (fs MockSession) Close() {}

// DB mocks mgo.Session.DB().
func (fs MockSession) DB(name string) Database {
	mockDatabase := MockDatabase{}
	return mockDatabase
}

// ********* Mocking Database functions**********

// C mocks mgo.Database(name).Collection(name).
func (db MockDatabase) C(name string) Collection {
	return MockCollection{
		collection: name,
	}
}

// ********* Mocking Collection functions**********

// Find mock.
func (fc MockCollection) Find(query interface{}) Query {
	return &MockQuery{
		collection: fc.collection,
		query:      query.(bson.M),
	}
}

// ********* Mocking Query functions**********

// Select mock
func (fq MockQuery) Select(selector interface{}) Query {
	return fq
}

// Sort mock
func (fq MockQuery) Sort(fields ...string) Query {
	return fq
}

// Skip mock
func (fq MockQuery) Skip(n int) Query {
	return fq
}

func (fp MockPipe) All(result interface{}) error {
	switch fp.collection {
	// case "offerSales":
	// 	matchData := fp.pipe[0]["$match"]
	// 	offerCode := matchData.(bson.M)["offerCode"]
	// 	if offerCode == "STAR12" || offerCode == "STAR17" {
	// 		return fmt.Errorf("NotFound")
	// 	}
	// 	strOfferSales := fmt.Sprintf("[%v]", samples.OfferAggregation[offerCode.(string)])
	// 	b := []byte(strOfferSales)

	// 	err := json.Unmarshal(b, &result)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	break
	// case "venueSales":
	// 	venueSalePipe := fp.pipe[0]["$match"]
	// 	fpVenueCodes := venueSalePipe.(bson.M)["venueCode"]
	// 	if venueCodes, ok := fpVenueCodes.(bson.M)["$in"]; ok {
	// 		venueCode := venueCodes.([]string)
	// 		if venueCode[0] == "UDAY" {
	// 			strVenueSales := fmt.Sprintf("[%v]", samples.VenueSales[venueCode[0]])
	// 			b := []byte(strVenueSales)
	// 			err := json.Unmarshal(b, &result)
	// 			if err != nil {
	// 				return err
	// 			}
	// 		}

	// 		if len(venueCode) > 2 && venueCode[2] == "INKP" {
	// 			return fmt.Errorf("NotFound")
	// 		}

	// 	}
	// 	break
	}
	return nil
}

// All mock
func (fq MockQuery) All(result interface{}) error {
	switch fq.collection {
	// case "boxOfficeReports":
	// 	venue := fq.query["venueCode"]
	// 	if venue == "PPPP" {
	// 		return fmt.Errorf("ErrNotFound")
	// 	}
	// 	data := fmt.Sprintf("[%v]", samples.BoxOfficeReports[venue.(string)])
	// 	byt := []byte(data)
	// 	err := json.Unmarshal(byt, &result)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	break
	}
	return nil
}

// One mock
func (fq MockQuery) One(result interface{}) error {
	switch fq.collection {

	// case "communicationCount":
	// 	key := fq.query["communicationKey"].(string)
	// 	if key == "OTP-0" || key == "RC-99301050256" || key == "RC-99301050257" {
	// 		return mgo.ErrNotFound
	// 	}
	// 	if key == "OTP-Error" {
	// 		return fmt.Errorf("Error")
	// 	}
	// 	b := []byte(samples.CommunicationCount[key])
	// 	err := json.Unmarshal(b, result)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return nil
	// }
	}

	return nil
}
