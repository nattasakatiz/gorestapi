package models

import "gopkg.in/mgo.v2/bson"

const (
	// CollectionScheduleUser holds the name of the schedule_users collection
	CollectionScheduleUser = "schedule_users"
)

// Schedule User model
type ScheduleUser struct {
	Id         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ScheduleId int           `bson:"schedule_id" json:"schedule_id"`
	UserId     int           `bson:"user_id" json:"user_id"`
	Active     bool          `bson:"active" json:"active"`
}
