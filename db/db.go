package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// Connect connects to mongodb
func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("MONGODB_HOST")
	port := os.Getenv("MONGODB_PORT")
	uri := host + ":" + port
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	dbname := os.Getenv("MONGODB_DATABASE")

	// Database connection
	mongo := &mgo.DialInfo{
		Addrs:    []string{uri},
		Database: dbname,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(mongo)
	if err != nil {
		panic(err)
	}

	session.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = session
	Mongo = mongo

}
