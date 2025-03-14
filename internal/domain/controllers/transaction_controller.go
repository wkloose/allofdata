package controllers

import (
	"net/http"
	"trashure/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

// GetAllTransactions - Ambil semua transaksi
func GetAllTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of transactions"})
}

// CreateTransaction - Buat transaksi baru
func CreateTransaction(c *gin.Context) {
	var newTransaction entity.Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created", "transaction": newTransaction})
}

// UpdateTransaction - Perbarui transaksi
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var updatedTransaction entity.Transaction
	if err := c.ShouldBindJSON(&updatedTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated", "id": id, "transaction": updatedTransaction})
}

// DeleteTransaction - Hapus transaksi
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted", "id": id})
}
