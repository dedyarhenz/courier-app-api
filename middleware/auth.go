package middleware

import (
	"final-project-backend/entity"
	"final-project-backend/pkg/helper"
	"final-project-backend/pkg/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		authCheck := c.Request.Header["Authorization"]
		if len(authCheck) < 1 {
			helper.ErrorResponse(c.Writer, "no authorization", http.StatusBadRequest)
			c.Abort()
			return
		}

		authString := authCheck[0]
		tokenCheck := strings.Split(authString, " ")
		if len(tokenCheck) < 1 {
			helper.ErrorResponse(c.Writer, "no authorization", http.StatusBadRequest)
			c.Abort()
		}

		token := tokenCheck[1]

		userId, role, err := utils.CheckToken(token)
		if err != nil {
			helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
			c.Abort()
		}

		c.Set("user_id", userId)
		c.Set("role", role)

		c.Next()

		log.Println(time.Since(now))
	}
}

func UserAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		authCheck := c.Request.Header["Authorization"]
		if len(authCheck) < 1 {
			helper.ErrorResponse(c.Writer, "no authorization", http.StatusBadRequest)
			c.Abort()
			return
		}

		authString := authCheck[0]
		tokenCheck := strings.Split(authString, " ")
		if len(tokenCheck) < 1 {
			helper.ErrorResponse(c.Writer, "no authorization", http.StatusBadRequest)
			c.Abort()
			return
		}

		token := tokenCheck[1]

		userId, role, err := utils.CheckToken(token)
		if err != nil {
			helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
			c.Abort()
			return
		}

		if role != entity.UserRole {
			helper.ErrorResponse(c.Writer, "access not allowed", http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Set("role", role)

		c.Next()

		log.Println(time.Since(now))
	}
}

func AdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		authCheck := c.Request.Header["Authorization"]
		if len(authCheck) < 1 {
			helper.ErrorResponse(c.Writer, "no authorization", http.StatusBadRequest)
			c.Abort()
			return
		}

		authString := authCheck[0]
		tokenCheck := strings.Split(authString, " ")
		if len(tokenCheck) < 1 {
			helper.ErrorResponse(c.Writer, "no authorization", http.StatusBadRequest)
			c.Abort()
			return
		}

		token := tokenCheck[1]

		userId, role, err := utils.CheckToken(token)
		if err != nil {
			helper.ErrorResponse(c.Writer, err.Error(), http.StatusBadRequest)
			c.Abort()
			return
		}

		if role != entity.AdminRole {
			helper.ErrorResponse(c.Writer, "access not allowed", http.StatusForbidden)
			c.Abort()
			return
		}

		c.Set("user_id", userId)
		c.Set("role", role)

		c.Next()

		log.Println(time.Since(now))
	}
}
