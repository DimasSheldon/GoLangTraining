package controller

import (
	"encoding/json"
	"finalproject/database"
	"finalproject/helper"
	"finalproject/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	userRequest := model.CreateUserRequest{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	user := model.User{
		Age:             userRequest.Age,
		Email:           userRequest.Email,
		Password:        helper.HashPass(userRequest.Password),
		Username:        userRequest.Username,
		ProfileImageUrl: "https://profile.com",
	}

	err := db.Debug().Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	userString, _ := json.Marshal(user)
	userResponse := model.CreateUserResponse{}
	json.Unmarshal(userString, &userResponse)

	c.JSON(http.StatusCreated, userResponse)
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helper.GetContentType(c)

	userRequest := model.LoginUserRequest{}

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	password := userRequest.Password
	user := model.User{}

	err := db.Debug().Where("email = ?", userRequest.Email).Take(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	comparePass := helper.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helper.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UpdateUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helper.GetContentType(c)

	userRequest := model.UpdateUserRequest{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	} else {
		if err := c.ShouldBind(&userRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
	}

	user := model.User{}
	user.ID = userID

	updateString, _ := json.Marshal(userRequest)
	updateData := model.User{}
	json.Unmarshal(updateString, &updateData)

	err := db.Model(&user).Updates(updateData).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	_ = db.First(&user, user.ID).Error

	userString, _ := json.Marshal(user)
	userResponse := model.CreateUserResponse{}
	json.Unmarshal(userString, &userResponse)

	c.JSON(http.StatusCreated, userResponse)
}

func DeleteUser(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)

	userID := uint(userData["id"].(float64))

	user := model.User{}
	user.ID = userID

	err := db.Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
