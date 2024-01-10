package middleware

import (
	"fmt"
	"net/http"
	"os"
	"reward/pkg/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		// Парсим и проверяем токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Возвращаем секретный ключ или отказ в парсинге токена
			return []byte(os.Getenv("JWT_SECRET")), nil //
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(string) // Access the "user_id" field from claims
		fmt.Println(userID)
		// Use the userID variable here
		id := uuid.Must(uuid.FromString(userID))
		c.Set("uuid", id)
		// Check if an instance of Wallet model exists with the given uuid
		var wallet model.Wallet
		result := db.Where("wallet_id = ?", id).First(&wallet)
		if result != nil {
			//c.JSON(http.StatusOK, gin.H{"message": "Wallet exists"})
			//return
			//c.Set("uuid", id)
			fmt.Println("Wallet Exists")
		} else {
			wallet = model.Wallet{
				WalletID: id,
				Amount:   0, // Set initial amount to 0 or any other value as desired
			}
			db.Create(&wallet)
			//c.Set("uuid", id)
		}

	}
}
