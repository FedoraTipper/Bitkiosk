package resolvers

import (
	"context"
	"errors"
	"github.com/fedoratipper/bitkiosk/server/internal/digest"
	"github.com/fedoratipper/bitkiosk/server/internal/handlers/authhandler"
	logger "github.com/fedoratipper/bitkiosk/server/internal/logger"

	"github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	tf "github.com/fedoratipper/bitkiosk/server/internal/gql/resolvers/transformations"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
)

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
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
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	authLevel, err := authhandler.GetAuthLevelFromContext(ctx)

	if authLevel == authhandler.UserAuth {
		return userList(r)
	} else {
		if err != nil {
			logger.Error("Unable to resolve auth level with Users resolver. \n" + err.Error())
		}
	}

	return nil, errors.New("not authenticated for query")
}

// ## Helper functions

func userCreate(r *mutationResolver, input models.NewUser) (*models.User, error) {
	userDbo, err := tf.GQLInputUserToDBUser(&input, false)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()

	var errBeforeCreate = userDbo.BeforeCreate(db)

	if errBeforeCreate != nil {
		return nil, errBeforeCreate
	}

	db = db.Create(userDbo).First(userDbo) // Create the user

	gql, err := tf.DBUserToGQLUser(userDbo)

	if err != nil {
		return nil, err
	}

	digest := digest.GetDigest(input.Token, uint(input.AuthMethodID))

	authenticationMatrixDbo := &dbm.AuthenticationMatrix{UserID:userDbo.ID, AuthMethodID: uint(input.AuthMethodID), Token:digest}

	db = db.Create(authenticationMatrixDbo).First(authenticationMatrixDbo) // Create new authentication matrix for user

	if db.Error != nil {
		return nil, db.Error
	}

	db = db.Commit()

	return gql, db.Error
}

func userUpdate(r *mutationResolver, input models.NewUser, ids ...uint) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, false)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()

	var errBeforeUpdate = dbo.BeforeUpdate(db)

	if errBeforeUpdate != nil {
		return nil, errBeforeUpdate
	}

	db = db.Model(&dbo).Update(dbo).First(dbo) // Or update it

	gql, err := tf.DBUserToGQLUser(dbo)

	if err != nil {
		return nil, err
	}
	db = db.Commit()

	return gql, db.Error
}

func userDelete(r *mutationResolver, id string) (bool, error) {
	return false, nil
}

func userList(r *queryResolver) ([]*models.User, error) {
	entity := "users"
	var res []*models.User

	db := r.ORM.DB.New()

	var dbRecords = []dbm.User{}
	db.Find(&dbRecords)

	db.Close()

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



//func userList(r *queryResolver, id *string) ([]*models.User, error) {
//	entity := "users"
//	whereID := "id = ?"
//	record := &models.User{}
//	dbRecords := []*dbm.User{}
//	db := r.ORM.DB.New()
//	if id != nil {
//		db = db.Where(whereID, *id)
//	}
//	db = db.Find(&dbRecords).Count(&record.Count)
//	for _, dbRec := range dbRecords {
//		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
//			log.Errorfn(entity, err)
//		} else {
//			record.List = append(record.List, rec)
//		}
//	}
//	return record, db.Error
//}
