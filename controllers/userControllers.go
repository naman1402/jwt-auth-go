package controllers

import (
	"jwt-go/initializers"
	"jwt-go/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	// request body from client in JSON format
	var body struct {
		Email    string
		Password string
	}
	// this function tries to parse the request body (in JSON format) and populate in the body struct
	// this function returns err, is err is nil `if` will not be executed
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Password Hashing: GenerateFromPassword() from brcypt library is used to get the hashed version of body.version
	//  second parameter is cost, controls amount of work done for hashing
	hash, e := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})
		return
	}

	// User Creation
	// 1. new variable of user of User type (struct)
	user := models.User{Email: body.Email, Password: string(hash)}
	// 2. using gorm library, create() creates a new user record in the db
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Successful Response
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}
	// Request body(JSON) binding into body struct and checking for error
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// First() is used to search for a user in the database with email from the req body, using GORM library
	// not found when id == 0
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	// ID not found in the DB
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Comparing the hashed password and stored password of the req body and stored user
	// if it does matches then invalid email or password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// JWT Token creation: sub is the user ID, exp is the expiration time which here is 30 days
	// Signing method used here is SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Token signing using secret stored in .env file
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	// Gin library, SameSite attribute for cookies (security reasons)
	c.SetSameSite(http.SameSiteLaxMode)
	// Cookie named Authorization is set with generated token that expires in 30 days
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	// SUCCESSFUL RESPONSE
	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
