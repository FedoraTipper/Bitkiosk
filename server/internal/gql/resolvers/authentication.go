package resolvers

import (
	"context"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	"github.com/fedoratipper/bitkiosk/server/internal/logger"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"strconv"
	"time"
)

func (r *queryResolver) Authenticate(ctx context.Context, authDetails models.LoginDetails) (*models.AuthResponse, error) {
	return authenticate(r, authDetails)
}

func authenticate(r *queryResolver, authDetails models.LoginDetails) (*models.AuthResponse, error) {
	// Load authentication matrix with user id and method id
	var authResponse models.AuthResponse

	var storedUserAuthMatrix dbm.AuthenticationMatrix
	var user dbm.User

	db := r.ORM.DB.New().Begin()

	if dbErr := db.Where("email = ?", authDetails.Identification).Find(&user); dbErr.Value == nil || user.ID == 0 {
		return &authResponse, nil
	}

	authMethod := dbm.GetAuthMethod(authDetails.AuthMethodID)

	if authMethod == nil{
		return &authResponse, nil
	}

	if dbErr := db.Where("auth_method_id = ? AND user_id = ?", authMethod.ID, user.ID).First(&storedUserAuthMatrix); dbErr.Value == nil || storedUserAuthMatrix.ID == 0 {
		return &authResponse, nil
	}

	//TODO: Flesh out later
	if digest.CompareDigest(authDetails.Token, storedUserAuthMatrix.Token, authMethod.ID) {

		durationString := strconv.Itoa(authMethod.TTL)
		ttl, err := time.ParseDuration(durationString + "M")

		if err != nil {
			logger.Error("Unable to parse auth method duration.\nError:\n" + err.Error())
			return &authResponse, nil
		}

		authLevel := session.AuthLevel{
			AuthLevel: int(user.Role),
			UID:       int(user.ID),
		}

		sessionToken, err := session.GenerateSession(ttl, authLevel)

		if err != nil {
			return &authResponse, nil
		}

		return &models.AuthResponse{
			TokenToStore: sessionToken,
			TTL:         authMethod.TTL,
			RefreshToken: "KEKW",
		}, nil
	}

	return &authResponse, nil
}

