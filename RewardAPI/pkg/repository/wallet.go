package repository

import (
	"reward/pkg/config"
	"reward/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func GetWallet(c *gin.Context) {
	uuidParam := c.MustGet("uuid").(uuid.UUID)
	var wallet model.Wallet

	// Find wallet by UUID
	db := config.SetupDatabaseConnection()
	result := db.First(&wallet, "wallet_id = ?", uuidParam)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"error": "Wallet not found",
		})
		return
	}

	c.JSON(200, wallet)
}
