package controllers

import (
	"net/http"
	"trashure/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of transactions"})
}

func CreateTransaction(c *gin.Context) {
	var newTransaction models.Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created", "transaction": newTransaction})
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var updatedTransaction models.Transaction
	if err := c.ShouldBindJSON(&updatedTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated", "id": id, "transaction": updatedTransaction})
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted", "id": id})
}
