package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nattasakatiz/gorestapi/controllers/scheduleUserController"
	"github.com/nattasakatiz/gorestapi/db"
	"github.com/nattasakatiz/gorestapi/middlewares"
)

const (
	Port = "8080"
)

// INIT RUN BEFORE MAIN
func init() {
	db.Connect()
}

// MAIN FUNCTION
func main() {

	// LOAD ENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// SET UP GIN AS ROUTER
	router := gin.Default()

	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true

	// Middlewares
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	// Statics
	router.Static("/public", "./public")

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/schedule-users")
	})

	router.GET("/schedule-users", scheduleUserController.Index)

	// Listen and Server (in port:8080)
	router.Run(":" + Port)
}
