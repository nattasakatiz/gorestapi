package scheduleUserController

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/nattasakatiz/gorestapi/models"
)

// New Schedule User
func New(c *gin.Context) {
	scheduleUser := models.ScheduleUser{}

	c.HTML(http.StatusOK, "schedule-users/form", gin.H{
		"title":        "New Schedule User",
		"scheduleUser": scheduleUser,
	})
}

// Create an scheduleUser
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	scheduleUser := models.ScheduleUser{}
	err := c.Bind(&scheduleUser)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionScheduleUser).Insert(scheduleUser)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/schedule-users")
}

// Edit an scheduleUser
func Edit(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	scheduleUser := models.ScheduleUser{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionScheduleUser).FindId(oID).One(&scheduleUser)
	if err != nil {
		c.Error(err)
	}

	c.HTML(http.StatusOK, "schedule-users/form", gin.H{
		"title":        "Edit Schedule User",
		"scheduleUser": scheduleUser,
	})
}

// List all schedule-users
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	scheduleUsers := []models.ScheduleUser{}
	err := db.C(models.CollectionScheduleUser).Find(nil).All(&scheduleUsers)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{"title": "Schedule Users", "scheduleUsers": scheduleUsers})
}

// Update an scheduleUser
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	scheduleUser := models.ScheduleUser{}
	err := c.Bind(&scheduleUser)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	doc := bson.M{
		"title": scheduleUser.ScheduleId,
		"body":  scheduleUser.UserId,
		// "updated_on": time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionScheduleUser).Update(query, doc)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/schedule-users")
}

// Delete an scheduleUser
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
	err := db.C(models.CollectionScheduleUser).Remove(query)
	if err != nil {
		c.Error(err)
	}
	c.Redirect(http.StatusMovedPermanently, "/schedule-users")
}
