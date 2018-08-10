package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	port      int
	heartbeat *time.Ticker
	router    *gin.Engine
)

// PingTheAPI ...
func PingTheAPI(c *gin.Context) {
	content := gin.H{"payload": "PONG"}
	c.JSON(http.StatusOK, content)
}

func init() {
	port = 7718

	// Set up the heartbeat ticker.
	heartbeat = time.NewTicker(60 * time.Second)

	// Setup the service router.
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()

	router.GET("/system/ping", PingTheAPI)
} // func

// main ...
func main() {

	go func() {
		for {
			select {
			case <-heartbeat.C:
				fmt.Printf("bump,Bump...\n")
			} // select
		} // for
	}() // go func

	fmt.Printf("ready on port %d\n", port)
	router.Run(":" + strconv.Itoa(port))

} // func
