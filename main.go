package main

import (
	"net/http"

	_ "github.com/douglas0n/go-api/docs/go-api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// mock
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// @title GO API!
// @version 1.0
// @description by Douglas Alves - TIMDA / DevOps.
// @termsOfService https://github.com/douglas0n/go-api/blob/master/LICENCE

// @contact.name TIMDA - DevOps
// @contact.url https://github.com/douglas0n
// @contact.email douglas.venancio@einstein.br

// @license.name MIT
// @license.url https://github.com/douglas0n/go-api/blob/master/LICENCE

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.GET("/", healthCheck)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("localhost:8080")
}

// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func healthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Return all albuns.
// @Description get all albuns.
// @Tags albuns
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// @Summary Return all albuns.
// @Description get all albuns.
// @Tags albuns
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// @Summary Return an albuns by given an ID.
// @Description get an albuns by given an ID.
// @Tags albuns
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [post]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
