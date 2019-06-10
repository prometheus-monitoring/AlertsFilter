package main

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/alertmanager/types"
)

type status string
type errorType string

const (
	statusCodeSuccess status = "success"
	statusCodeError   status = "error"
)

type response struct {
	Status    status      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	ErrorType errorType   `json:"errorType,omitempty"`
	Error     string      `json:"error,omitempty"`
}

func main() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.POST("/v1/alerts", postAlerts)
	}

	router.Run(":3909")
}

func postAlerts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var alerts []*types.Alert
	c.BindJSON(&alerts)
	responseSucces(c, alerts)
	fmt.Println(alerts)
}

func responseSucces(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application")
	c.JSON(200, response{
		Status: statusCodeSuccess,
		Data:   data,
	})
}

//
// func ()  {
//
// }
