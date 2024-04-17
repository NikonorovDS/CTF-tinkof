package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"ticket/internal/model"
	"ticket/internal/storage"
	"ticket/pkg/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func FindTicketByID(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id").(uint)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errWrongParamType.Error()})
			return
		}

		ticket, err := s.Tickets().FindById(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": errIdNotFound.Error()})
			return
		}

		if ticket.UserID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": errPermissionDenied.Error()})
			return
		}

		c.JSON(http.StatusOK, ticket.ToDTO())
	}
}

func BuyTicket(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id").(uint)

		ticket, err := s.Tickets().BuyTicket(userID)
		if err != nil {
			logger.Errorf("%+v", err)
			if err.Error() == errNotEnoughMoney.Error() {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": errDBSomethingWrong.Error()})
				return
			}

		}
		c.JSON(http.StatusOK, ticket.ToDTO())
	}
}

func EatTicket(s storage.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id").(uint)

		var ticketReqDTO model.TicketReqDTO
		decoder := json.NewDecoder(c.Request.Body)
		if err := decoder.Decode(&ticketReqDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": errWrongJsonData.Error()})
			return
		}
		defer c.Request.Body.Close()

		newLuck, ticketNumber, err := s.Tickets().EatTicket(userID, ticketReqDTO.ID)
		if err != nil {
			var responseStatusCode int
			responseErr := err
			switch err {
			case errIdNotFound, errTicketAlreadyUsed:
				responseStatusCode = http.StatusBadRequest
			case errPermissionDenied:
				responseStatusCode = http.StatusForbidden
			default:
				responseStatusCode = http.StatusInternalServerError
				responseErr = errDBSomethingWrong
			}
			c.JSON(responseStatusCode, gin.H{"error": responseErr.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"luck": newLuck, "ticket": ticketNumber})
	}
}
