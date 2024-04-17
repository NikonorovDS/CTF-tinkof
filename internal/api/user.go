package api

import (
	"fmt"
	"net/http"
	"ticket/internal/model"
	"ticket/internal/storage"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserSelfInfo(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id").(uint)

		user, err := s.Users().FindById(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%w: %w", errDBSomethingWrong, err)})
			return
		}

		c.JSON(http.StatusOK, user.ToDTO())
	}
}

func GetUserSelfTickets(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id").(uint)

		user, err := s.Users().FindById(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("%w: %w", errDBSomethingWrong, err)})
			return
		}

		tickets, err := s.Tickets().AllByUserId(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errDBSomethingWrong.Error()})
			return
		}

		c.JSON(http.StatusOK, model.ToTicketDTOs(tickets))
	}
}
