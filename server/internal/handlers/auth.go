package handlers

import (
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils"
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

type authResponse struct {
	TokenToStore string `json:"tokenToStore"`
	TTL          int    `json:"TTL"`
	RefreshToken string `json:"refreshToken"`
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

		var authDetails authResponse

		var storedUserAuthMatrix dbm.AuthenticationMatrix
		var user dbm.User

		db := orm.DB.New().Begin()

		if dbErr := db.Where("email = ?", loginDetails.Identification).Find(&user); dbErr.Value == nil || user.ID == 0 {
			setAuthResponse(c, &authDetails)
			return
		}

		authMethod := dbm.GetAuthMethod(loginDetails.AuthMethodId)

		if authMethod == nil{
			setAuthResponse(c, &authDetails)
			return
		}

		if dbErr := db.Where("auth_method_id = ? AND user_id = ?", authMethod.ID, user.ID).First(&storedUserAuthMatrix); dbErr.Value == nil || storedUserAuthMatrix.ID == 0 {
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
				UID:       int(user.ID),
			}

			sessionToken, err := session.GenerateSession(ttl, authLevel)

			if err != nil {
				setAuthResponse(c, &authDetails)
				return
			}

			authDetails := authResponse{
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


func setAuthResponse(c *gin.Context, authResponse *authResponse) {
	if authResponse.TokenToStore == "" {
		c.JSON(401, gin.H{"error": "Incorrect login details"})
	} else {
		c.SetCookie("Authorization", authResponse.TokenToStore, authResponse.TTL, "/", domain, true, true)
		c.JSON(200, gin.H{"error":""})
	}
}


type sessionDetails struct {
}

func LogoutHandler(orm *orm.ORM) gin.HandlerFunc {
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
	}
}