package main

import (
	"fmt"
	"math"
	"net/http"

	rtreego "github.com/dhconnelly/rtreego"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RouteRequest struct {
	Start    Location `json:"start"`
	Distance float64  `json:"distance"`
}

type RouteResponse struct {
	Distance float64    `json:"distance"`
	Nodes    []Location `json:"nodes"`
}

func (res *RouteResponse) fromSearchResult(s *SearchResult) {
	res.Distance = s.Distance
	for _, node := range s.Nodes {
		location := Location{
			Latitude:  node.Lat,
			Longitude: node.Lon,
		}
		res.Nodes = append(res.Nodes, location)
	}
}

const NumRoutes = 9

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/route", createRoute)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func createRoute(c *gin.Context) {
	var request RouteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	graph, err := LoadGraph("map.osm")
	if err != nil {
		panic(err)
	}

	goalDistance := request.Distance

	startPoint := rtreego.Point{request.Start.Longitude, request.Start.Latitude}
	startID := graph.Tree.NearestNeighbor(startPoint).(RTNode).ID
	fmt.Println(startID)
	startNode := graph.Nodes[startID]
	routes := graph.Search(startNode, goalDistance)

	var bestRoute SearchResult
	bestRoute.Distance = goalDistance * 2
	for _, route := range routes {
		if math.Abs(route.Distance-goalDistance) < math.Abs(bestRoute.Distance-goalDistance) {
			bestRoute = route
		}
	}

	var r RouteResponse
	r.fromSearchResult(&bestRoute)
	c.JSON(http.StatusOK, r)
}
