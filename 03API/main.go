package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	TITLE  string  `json:"title"`
	ARTIST string  `json:"artist"`
	PRICE  float32 `json:"price"`
}

var albums = []album{
	{ID: "1", TITLE: "RED", ARTIST: "Taylor swift", PRICE: 10000.99},
	{ID: "2", TITLE: "Lover", ARTIST: "Taylor swift", PRICE: 700.99},
	{ID: "3", TITLE: "Tortured poets", ARTIST: "Taylor swift", PRICE: 12000.99},
	{ID: "4", TITLE: "Reputation", ARTIST: "Taylor swift", PRICE: 12000.99},
}

func getalbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getalbums)
	router.GET("/albums/:id",getAlbumByID)
	router.Run("localhost:8000")
}

