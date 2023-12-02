package handlers

import (
	"errors"
	"net/http"

	db "api/db"
	"api/models"
	"api/security"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var login LoginData
	if err := c.BindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid login data"})
		return
	}
	var user db.User
	if res := db.DB.First(&user, "email = ?", login.Email); errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(404, gin.H{"message": "username or password not found"})
		return
	}

	err := security.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		c.JSON(404, gin.H{"message": "username or password not found"})
		return
	}

	token, tokErr := security.CreateJwt(user.UserTypeID, user.Email, uint(user.ID))
	if tokErr != nil {
		c.JSON(400, gin.H{"message": "something went wrong"})
		return
	}

	c.IndentedJSON(200, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var register models.Register
	if err := c.BindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user data"})
		return
	}

	if !(register.ConfirmPassword == register.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "passwords do not match"})
		return
	}

	user := db.User{}
	user.Email = register.Email

	hashedPass, err := security.HashPassword([]byte(register.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	user.Password = string(hashedPass)
	user.EmailConfirmationCode = uuid.New().String()
	result := db.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}
	c.IndentedJSON(200, "success")
}
