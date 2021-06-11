package routes

import (
	"github.com/gin-gonic/gin"
	gps "iot.api/internal/services/gps"
)

type CreateParams struct {
	Id  int64 `json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
	Direction float64 `json:"direction"`
	Speed float64 `json:"speed"`
}

func Hi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hi",
	})
}

func Home(c *gin.Context) {	
	c.JSON(200, gin.H{
		"message": "Home",
	})
}

func TestWrite(c *gin.Context) {
	var p CreateParams

	c.BindJSON(&p)
	err := gps.AsyncWrite(p.Id, p.Lat, p.Lng, p.Direction, p.Speed)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "success",
		})
	}
}

func TestWriteRoutine(c *gin.Context) {
	var p CreateParams

	c.BindJSON(&p)
	go gps.AsyncWrite(p.Id, p.Lat, p.Lng, p.Direction, p.Speed)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func TestRandomWrite(c *gin.Context) {
	gps.AsyncRandomWrite(1)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func TestRandomWriteRoutine(c *gin.Context) {
	go gps.AsyncRandomWrite(1)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func TestWriteBatch(c *gin.Context) {
	gps.AsyncRandomWrite(1)

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func FindAll(c *gin.Context) {
	num := gps.FindAll()

	c.JSON(200, gin.H{
		"count": num,
	})
}
