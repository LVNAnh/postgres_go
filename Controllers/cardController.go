package Controllers

import (
	"net/http"
	"postgres_go/Models"
	"postgres_go/config"

	"github.com/gin-gonic/gin"
)

func GetCards(c *gin.Context) {
	var cards []Models.CreditCard
	result := config.DB.Find(&cards)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, cards)
}

func GetCard(c *gin.Context) {
	var card Models.CreditCard
	id := c.Param("id")
	result := config.DB.First(&card, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	c.JSON(http.StatusOK, card)
}

func AddCard(c *gin.Context) {
	var card Models.CreditCard

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user Models.User
	if err := config.DB.First(&user, card.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	result := config.DB.Create(&card)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Credit card created successfully!",
		"card":    card,
	})
}

func UpdateCard(c *gin.Context) {
	var card Models.CreditCard
	id := c.Param("id")

	if err := config.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&card)
	c.JSON(http.StatusOK, gin.H{"message": "Card updated successfully!", "card": card})
}

func DeleteCard(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&Models.CreditCard{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Card deleted successfully!"})
}
