package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var router *gin.Engine

func main() {
	// Set the router as the default one provided by gin
	router = gin.Default()

	/* Process the templates at the start so that they don't have to be loaded
	from the disk again. This makes serving HTML pages very fast.
	*/
	router.LoadHTMLGlob("templates/*")
	// Define the route for the index page and display the index.html template
	// To start with, we'll use an inline route handler. Later on, we'll create
	// standalone functions that will be used as route handlers.
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (Ok)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)
	})

	// Start serving the application
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
	router.Run()
}
