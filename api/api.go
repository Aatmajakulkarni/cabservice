package api

import (
	"fmt"
	"strings"
	db "cabservice/db"
	utils "cabservice/utils"
	models "cabservice/model"
	"github.com/gin-gonic/gin"
)
// function to validate users request and fetching user
func ValidateUserRequestAndFetchUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("authorization") != "" {
			headerPieces := strings.Split(c.Request.Header.Get("authorization"), " ")
			fmt.Printf("\n ValidateUserRequestAndFetchUser 0 %+v", headerPieces[1])
			userId, err := utils.DecodeToken(headerPieces[1])
			if err != nil {
				fmt.Printf("\n ValidateUserRequestAndFetchUser 1 %+v", err)
				c.AbortWithStatusJSON(200, utils.GetError(utils.USER_ERROR_INVALID_SESSION))
				return
			}
			c.Set("user_id", userId)
			if user, err := db.GetUserInfoById(userId); nil == err {
				if user.IsDisabled {
					fmt.Printf("\n ValidateUserRequestAndFetchUser 2 IsDisabled")
					c.AbortWithStatusJSON(200, utils.GetError(utils.USER_ERROR_INVALID_SESSION))
					return
				}
				c.Set("user", user)
				c.Next()
			} else {
				fmt.Printf("\nget complete user by id error\n %+v", err)
				c.AbortWithStatusJSON(200, utils.GetError(utils.USER_ERROR_INVALID_SESSION))
				return
			}

		} else {
				fmt.Printf("\nRequest.Header nil\n")
			c.AbortWithStatusJSON(200, utils.GetError(utils.USER_ERROR_INVALID_SESSION))
			return
		}
	}
}

func getUserFromSession(c *gin.Context) (models.UserInfo, bool) {
	userInSession, Ok := c.Get("user")
	if Ok {
		return userInSession.(models.UserInfo), true
	}
	return models.UserInfo{}, false
}
