package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"magic-link-auth/internal/database"
	"magic-link-auth/internal/email"
	"magic-link-auth/internal/models"
	"magic-link-auth/internal/utils"
)

func SendMagicLink(c *gin.Context) {
	var body struct {
		Email string `json:"email"`
	}

	if err := c.BindJSON(&body); err != nil || body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email required"})
		return
	}

	//check if user exists
	var user  models.User
	result := database.DB.Where("email = ?", body.Email).First(&user)

	if result.RowsAffected == 0 {
		user =models.User{
			ID: uuid.New(),
			Email: body.Email,
		}
		database.DB.Create(&user)
	}

	//Generate secure token
	token := utils.GenerateToken(32)

	//Save token to DB
	magicToken := models.MagicToken{
		UserID: user.ID,
		Token: token,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	database.DB.Create(&magicToken)

	err := email.SendMagicLink(user.Email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Magic link sent to your email",
	})

}

func RegisterAuthRoutes(r *gin.Engine) {
    auth := r.Group("/auth")
    {
        auth.POST("/send-link", SendMagicLink)
    }
}