package resolvers

import (
	"context"
	"errors"
	authhandler "github.com/fedoratipper/bitkiosk/server/internal/authentication"
	"github.com/fedoratipper/bitkiosk/server/internal/authentication/session"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	logger "github.com/fedoratipper/bitkiosk/server/internal/logger"
	"github.com/fedoratipper/bitkiosk/server/internal/orm"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/auth"
	"github.com/fedoratipper/bitkiosk/server/internal/orm/models/user"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/passwordvalidator"
	stringUtil "github.com/fedoratipper/bitkiosk/server/pkg/utils/string"
	"strings"
)

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (string, error) {
	return userCreate(r, input)
}

// UpdateUser updates a record
func (r *mutationResolver) UpdateUser(ctx context.Context, id uint, input models.NewUser) (*models.User, error) {
	return userUpdate(r, input, id)
}

// DeleteUser deletes a record
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	return userDelete(r, id)
}

// Users lists records
func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*models.User, error) {
	authLevel, err := authhandler.GetAuthLevelFromSession(ctx)

	if authLevel != nil && authLevel.AuthLevel == session.AdminAuth {
		return userList(r)
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

// ## Helper functions
func userCreate(r *mutationResolver, input models.NewUser) (string, error) {
	passwordValidator := passwordvalidator.DefaultPasswordValidator(input.Token)

	errs := passwordValidator.Validate()

	if errs != nil {
		var jointError []string
		for _, err := range errs {
			jointError = append(jointError, err.Error())
		}
		return "", errors.New("Please enter enter a password which meets the requirements: " + strings.Join(jointError, ", "))
	}

	var gqlReturn string

	userDbo, err := tf.GQLInputUserToDBUser(&input, false)

	if err != nil {
		return "", err
	}

	userDbo.Role = session.UserAuth

	// Create scoped clean tx interface
	tx := r.ORM.DB.New().Begin()

	tx, err = userDbo.Create(tx)

	if err == nil {
		tokenDigest := digest.GetDigest(input.Token, uint(input.AuthMethodID))

		tx, err = (&auth.AuthenticationMatrix{UserID: userDbo.Id, AuthMethodID: input.AuthMethodID, Token: tokenDigest}).Create(tx)

		if err == nil {
			userProfileDbo := &user.UserProfile{
				UserID:      userDbo.Id,
				FirstName:   stringUtil.FormatNameString(input.FirstName),
				LastName:    stringUtil.FormatNameString(input.LastName),
			}

			tx, err = userProfileDbo.Create(tx)

			if err == nil {
				gqlReturn = "success"
			}
		}
	}

	orm.CommitOrRollBackIfErrorAndCloseSession(tx, err)

	return gqlReturn, err
}

func userUpdate(r *mutationResolver, input models.NewUser, ids ...uint) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, false)
	if err != nil {
		return nil, err
	}
	// Create scoped clean tx interface
	tx := r.ORM.DB.New().Begin()

	var errBeforeUpdate = dbo.BeforeUpdate(tx)

	if errBeforeUpdate != nil {
		orm.CommitOrRollBackIfErrorAndCloseSession(tx, err)
		return nil, errBeforeUpdate
	}

	tx = tx.Model(&dbo).Update(dbo).First(dbo) // Or update it

	gql, err := tf.DBUserToGQLUser(dbo)

	orm.CommitOrRollBackIfErrorAndCloseSession(tx, err)

	return gql, err
}

func userDelete(r *mutationResolver, id string) (bool, error) {
	return false, nil
}

func userList(r *queryResolver) ([]*models.User, error) {
	//TODO: Comeback to this.
	entity := "users"
	var res []*models.User

	db := r.ORM.DB

	var dbRecords = []user.User{}
	db.Preload("UserProfile").Find(&dbRecords)

	if dbRecords == nil {
		return nil, nil
	}

	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(&dbRec); err != nil {
			logger.Errorfn(entity, err)
		} else {
			res = append(res, rec)
		}
	}

	return res, nil
}

