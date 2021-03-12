package auth

import (
	"local/gonotes/common"
	"local/gonotes/db"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func postRegisterView(c *gin.Context) {
	var userCredentials UserCredentials
	if err := c.BindJSON(&userCredentials); err != nil {
		c.JSON(http.StatusBadRequest, common.StoreError(err))
		return
	}

	if err := userCredentials.hashAndSalt(); err != nil {
		c.JSON(http.StatusInternalServerError, common.StoreError(err))
		return
	}

	db := db.GetDbConnection()
	defer db.Close()

	if err := db.Insert(&userCredentials); err != nil {
		c.JSON(http.StatusInternalServerError, common.StoreError(err))
		return
	}

	c.JSON(http.StatusOK, userCredentials.ID)
}

func postLoginView(c *gin.Context) {
	var loggingUser UserCredentials
	if err := c.BindJSON(&loggingUser); err != nil {
		c.JSON(http.StatusBadRequest, common.StoreError(err))
		return
	}

	db := db.GetDbConnection()
	defer db.Close()

	storedUser := &UserCredentials{Username: loggingUser.Username}
	err := db.Model(storedUser).Where("username = ?", loggingUser.Username).First()
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.StoreError(err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loggingUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.StoreError(err))
		return
	}

	atClaims := jwt.MapClaims{"authorized": true, "user_id": storedUser.ID, "exp": time.Now().Add(time.Minute * 15).Unix()}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.StoreError(err))
		return
	}

	c.JSON(http.StatusOK, LoginData{token})

	jwt.
}
