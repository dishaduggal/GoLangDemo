package controllers

import (
	"bms-movies/app"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogout(t *testing.T) {
	// mock complete context
	app.MockCtx()

	// test table
	tt := []struct {
		name       string
		req        string
		statusCode int
	}{
		{"Success", "/venue-listing?venueId=C1", http.StatusOK},
		{"Failure", "venue-listing?venue", http.StatusBadRequest},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = &CustomValidator{}
			req := httptest.NewRequest("", tc.req, strings.NewReader(""))
			rec := httptest.NewRecorder()
			echoCtx := e.NewContext(req, rec)

			GetShowTimes(echoCtx)

			res := rec.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.statusCode {
				t.Errorf("expected %d; got %v", tc.statusCode, res.StatusCode)
			}
			_, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Could not read response: %v", err.Error())
			}
		})
	}
}
