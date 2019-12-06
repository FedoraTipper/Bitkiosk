package resolvers

import (
	"context"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"

)


func (r *queryResolver) Authenticate(ctx context.Context, authDetails models.LoginDetails) (*models.AuthResponse, error) {
	return authenticate(r, authDetails)
}

func (r *mutationResolver) AddAuthMethod(ctx context.Context, newMethod *models.NewAuthMethod) (bool, error) {
	return authMethodCreate(r, newMethod)
}

func authenticate(r *queryResolver, authDetails models.LoginDetails) (*models.AuthResponse, error){
	// Load authentication matrix with user id and method id
	var authResponse models.AuthResponse

	var storedUserAuthMatrix dbm.AuthenticationMatrix
	var authMethod dbm.AuthMethod
	var user dbm.User

	db := r.ORM.DB.New().Begin()

	if dbErr := db.Where("email = ?", authDetails.Identification).Find(&user); dbErr.Value == nil {
		return &authResponse, nil
	}

	if dbErr := db.Where("method_id = ?", authDetails.AuthMethodID).Find(&authMethod); dbErr.Value == nil {
		return &authResponse, nil
	}

	if dbErr := db.Where("auth_method_id = ? AND user_id = ?", authMethod.MethodId, user.ID).First(&storedUserAuthMatrix); dbErr.Value == nil {
		return &authResponse, nil
	}

	//TODO: Flesh out later
	if digest.CompareDigest(authDetails.Token, storedUserAuthMatrix.Token, authMethod.MethodId) {
		return &models.AuthResponse{
			TokenToStore: "Success",
			TTL:          123,
			RefreshToken: "KEKW",
		}, nil
	}

	return &authResponse, nil
}

func authMethodCreate(r *mutationResolver, newMethod *models.NewAuthMethod) (bool, error) {
	authMethodDbo, err := tf.GQLInputAuthMethodToDBAuthMethod(newMethod, false)

	if err != nil {
		return false, err
	}

	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()

	var errBeforeCreate = authMethodDbo.BeforeCreate(db)

	if errBeforeCreate != nil {
		return false, errBeforeCreate
	}

	db = db.Create(authMethodDbo).First(authMethodDbo) // Create the user

	if db.Error != nil {
		return false, db.Error
	}

	db = db.Commit()

	return db.Error == nil, db.Error
}