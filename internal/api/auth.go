package api

import (
	"errors"
	"net/http"
	"ticket/internal/helpers"
	"ticket/internal/model"
	"ticket/internal/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	msgLoginSuccess  = "вход выполнен успешно"
	msgLogoutSuccess = "выход выполнен успешно"
)

func RegisterUser(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": errEmptyUsernameOrPassword.Error()})
			return
		}

		if len(username) < 9 || len(password) < 9 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errTooShortUsernameOrPassword.Error()})
			return
		}

		_, err := s.Users().FindByUsername(username)
		if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errDBSomethingWrong.Error()})
			return
		}

		if err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errUsernameAlreadyExists.Error()})
			return
		}

		user := model.NewUser(username, password)
		createdUser, err := s.Users().Create(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errDBSomethingWrong.Error()})
			return
		}

		c.JSON(http.StatusOK, createdUser.ToDTO())
	}
}

func LoginUser(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": errEmptyUsernameOrPassword.Error()})
			return
		}

		user, err := s.Users().FindByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidLoginCreds.Error()})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": errDBSomethingWrong.Error()})
				return
			}
		}

		if !helpers.IsPasswordCorrect(password, user.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": errInvalidLoginCreds.Error()})
			return
		}

		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"message": msgLoginSuccess})
	}
}

func LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()

		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errFailedClearSession.Error()})
		}

		if cookie, err := c.Request.Cookie("session"); err == nil {
			cookie.Value = ""
			cookie.Path = "/"
			cookie.MaxAge = -1
			http.SetCookie(c.Writer, cookie)
		}

		c.JSON(http.StatusOK, gin.H{"message": msgLogoutSuccess})
	}
}
