package transformations

import (
	"errors"
	gql "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/date"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	formattedCreatedAt := date.FormatToSqlDate(i.CreatedAt)
	formattedUpdatedAt := date.FormatToSqlDate(*i.UpdatedAt)

	o = &gql.User{
		Email:     i.Email,
		Role:	   int(i.Role),
		CreatedAt: &formattedCreatedAt,
		UpdatedAt: &formattedUpdatedAt,
	}
	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.NewUser, update bool) (o *dbm.User, err error) {
	o = &dbm.User{
		Email:     i.Email,
	}

	if i.Token == "" && !update {
		return nil, errors.New("Field [password] is required")
	}

	if i.Email == "" && !update {
		return nil, errors.New("Field [email] is required")
	}

	if i.Email != "" {
		o.Email = i.Email
	}

	return o, err
}

func DBUserProfileToGQLUserProfile(i *dbm.UserProfile) (o *gql.UserProfile, err error) {

	if i.ID == 0 {
		return nil, errors.New("unable to find user profile")
	}

	formattedDateOfBirth := date.FormatToSqlDate(i.DateOfBirth)

	o = &gql.UserProfile{
		FirstName:   &i.FirstName,
		LastName:    &i.LastName,
		DateOfBirth: &formattedDateOfBirth,
	}

	return o, err
}

