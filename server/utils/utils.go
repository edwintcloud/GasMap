package utils

import (
	"github.com/globalsign/mgo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// DB is our exported package global reference to mgo connection
	DB *mgo.Database

	// GoogleOauth is our exported package global reference to oauth config for google signin
	GoogleOauth *oauth2.Config

	// JwtSecret is our exported package global jwt secret string
	JwtSecret string
)

// ConnectToDb is our database connection function to set DB
func ConnectToDb(url, name string) (*mgo.Session, error) {

	// connect to mongodb
	session, err := mgo.Dial(url)
	if err != nil {
		return session, err
	}

	// set DB to our db instance
	DB = session.DB(name)

	// ensure our indexes exist and duplicates don't exist for indexed fields
	DB.C("users").EnsureIndex(mgo.Index{
		Key:      []string{"email"},
		Unique:   true,
		DropDups: true, // delete duplicate documents in case they somehow get put in
	})

	// return session and nil
	return session, nil

}

// ConfigureOauth sets our GoogleOauth var
func ConfigureOauth(r, id, secret string) {
	GoogleOauth = &oauth2.Config{
		RedirectURL:  r,
		ClientID:     id,
		ClientSecret: secret,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
