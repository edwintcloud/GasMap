package utils

import (
	"github.com/globalsign/mgo"
)

// DB is our exported package global reference to mgo connection
var DB *mgo.Database

// Connect is our database connection function to set DB
func ConnectToDb(url string, name string) (*mgo.Session, error) {

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
