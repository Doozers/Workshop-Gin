package controllers

import (
	"github.com/gin-gonic/gin"
	"workshop/database"
	s "workshop/structure"
)

func ItemTemplate(c *gin.Context) {
	c.JSON(200, s.Item{
		Name:        "name",
		Description: "description",
		Price:       0,
	})
}

func ItemDisplay(c *gin.Context) {
	i := new(s.Item)
	if err := c.ShouldBindJSON(i); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, i)
}

func GetItem(c *gin.Context) {
	getDatabase, err := database.GetDatabase()
	if err != nil {
		return
	}

	c.JSON(200, getDatabase)
}

func CreateItem(c *gin.Context) {
	i := new(s.Item)
	if err := c.ShouldBindJSON(i); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := database.AddItemInDB(*i)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, i)
}

func DeleteItem(c *gin.Context) {
	i := new(s.Item)
	if err := c.ShouldBindJSON(i); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := database.DeleteItemInDB(*i)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, i)
}

func UpdateItem(c *gin.Context) {
	i := new(s.Item)
	if err := c.ShouldBindJSON(i); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdateItemInDB(*i)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, i)
}

func CreateItems(c *gin.Context) {
	i := new([]s.Item)
	if err := c.ShouldBindJSON(i); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := database.AddItemsInDB(*i)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, i)
}
