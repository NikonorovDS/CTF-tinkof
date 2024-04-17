package api

import (
	"fmt"
	"net/http"
	"ticket/internal/helpers"
	"ticket/internal/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CheckLuckForInterview(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id").(uint)

		user, err := s.Users().FindById(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%w: %w", errDBSomethingWrong, err)})
			return
		}

		if user.Luck != 100 {
			c.JSON(http.StatusBadRequest, gin.H{"luck": user.Luck, "error": errInterviewFailed.Error(), "message": helpers.GetReviewFailMsg()})
			return
		}

		user.Luck = 0
		_, err = s.Users().Update(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errDBSomethingWrong.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"flag": viper.GetString("flag")})
	}
}
