package handlers

import (
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/auth"
	user2 "github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
	timeUtil "github.com/fedoratipper/bitkiosk/server/pkg/utils/time"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// Standard gin endpoint for authentication
type loginDetails struct {
	Identification string `json:"identification" binding:"required"`
	Token string `json:"token" binding:"required"`
	AuthMethodId int `json:"authMethodId" binding:"required"`
}

type authenticationDetails struct {
	TokenToStore string
	TTL          int
	RefreshToken string
}

type authResponse struct {
	Error string `json:"error"`
	TokenToStore string `json:"tokenToStore"`
}

var domain string

func init() {
	domain = utils.MustGet("DOMAIN")
}

//TODO: MOVE BUSINESS LOGIC INTERNAL/AUTHENTICATION
func AuthenticationHandler(orm *orm.ORM) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginDetails loginDetails

		err := c.ShouldBindJSON(&loginDetails); if err != nil {
			var errorMessage = ""
			if loginDetails.Identification == "" || loginDetails.Token == "" || loginDetails.AuthMethodId == 0 {
				errorMessage = "Poorly formed login request."
			}

			c.JSON(401, gin.H{"error": errorMessage})
			return
		}

		var authDetails authenticationDetails

		var storedUserAuthMatrix auth.AuthenticationMatrix
		var user user2.User

		db := orm.DB.New().Begin()

		if dbErr := db.Where("email = ?", loginDetails.Identification).Find(&user); dbErr.Value == nil || user.Id == 0 {
			setAuthResponse(c, &authDetails)
			return
		}

		authMethod := auth.GetAuthMethod(loginDetails.AuthMethodId)

		if authMethod == nil{
			setAuthResponse(c, &authDetails)
			return
		}

		if dbErr := db.Where("auth_method_id = ? AND user_id = ?", authMethod.ID, user.Id).First(&storedUserAuthMatrix); dbErr.Value == nil || storedUserAuthMatrix.Id == 0 {
			setAuthResponse(c, &authDetails)
			return
		}

		if digest.CompareDigest(loginDetails.Token, storedUserAuthMatrix.Token, authMethod.ID) {

			durationString := strconv.Itoa(authMethod.TTL)
			ttl, err := time.ParseDuration(durationString + "m")

			if err != nil {
				logger.Error("Unable to parse auth method duration.\nError:\n" + err.Error())
				setAuthResponse(c, &authDetails)
				return
			}

			authLevel := session.AuthLevel{
				AuthLevel: int(user.Role),
				UID:       int(user.Id),
			}

			sessionToken, err := session.GenerateSession(ttl, authLevel)

			if err != nil {
				setAuthResponse(c, &authDetails)
				return
			}

			authDetails := authenticationDetails{
				TokenToStore: sessionToken,
				TTL:          authMethod.TTL,
			}

			setAuthResponse(c, &authDetails)
			return
		}

		setAuthResponse(c, &authDetails)
		return
	}
}


func setAuthResponse(c *gin.Context, authDetails *authenticationDetails) {
	if authDetails.TokenToStore == "" {
		c.JSON(401, gin.H{
				"error":       "Incorrect login details",
			})
	} else {
		c.SetCookie("Authorization", authDetails.TokenToStore, timeUtil.ConvertMinutesToSeconds(authDetails.TTL), "/", domain, false, true)
		c.JSON(200, gin.H{
			"error":"",
		})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authCookie, err := c.Request.Cookie("Authorization")
		if err == nil && authCookie != nil && authCookie.Value != "" {
			err = session.DestroySession(authCookie.Value)

			if err == nil {
				c.JSON(200, gin.H{"error":""})
			} else {
				logger.Errorfn("LogoutHandler", err)
				c.JSON(200, gin.H{"error":""})
			}
		}

		c.SetCookie("Authorization", "", 0, "", "", false, false)
	}
}