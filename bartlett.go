package bartlett

import (
	"database/sql"
	"net/http"
)

// A UserIDProvider is a function that is able to use an incoming request to produce a user ID.
type UserIDProvider func(r *http.Request) (interface{}, error)

// Bartlett holds all of the configuration necessary to generate an API from the database.
type Bartlett struct {
	DB     *sql.DB
	Driver Driver
	Tables []Table
	Users  UserIDProvider
}

// A ResultMarshaler interprets a database result set, selects appropriate Go types, and writes JSON to the output stream.
// This is necessary because all of the databases have different types and their drivers respond to ScanType in different ways.
type ResultMarshaler func(rows *sql.Rows, w http.ResponseWriter) error
