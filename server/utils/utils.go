package utils

import (
	"github.com/globalsign/mgo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// DB is our exported package global reference to mgo connection
var DB *mgo.Database

// GoogleOauth is our exported package global reference to oauth config for google signin
var GoogleOauth *oauth2.Config

// ConnectToDb is our database connection function to set DB
func ConnectToDb(url, name string) (*mgo.Session, error) {

	// connect to mongodb
	session, err := mgo.Dial(url)
	if err != nil {
		return session, err
	}

	// set DB to our db instance
	DB = session.DB(name)

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
