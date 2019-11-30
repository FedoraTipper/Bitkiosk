package transformations

import (
	"errors"

	gql "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		Email:     i.Email,
		FirstName: i.FirstName,
		LastName:  i.LastName,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	}
	return o, err
}

// GQLInputUserToDBUser transforms [user] gql input to db model
func GQLInputUserToDBUser(i *gql.NewUser, update bool) (o *dbm.User, err error) {
	o = &dbm.User{
		Email:     i.Email,
		FirstName: i.FirstName,
		LastName:  i.LastName,
	}

	if i.Email == "" && !update {
		return nil, errors.New("Field [email] is required")
	}

	if i.Email != "" {
		o.Email = i.Email
	}

	return o, err
}

func GQLInputAuthMethodToDBAuthMethod(i *gql.NewAuthMethod, update bool) (o *dbm.AuthMethod, err error) {
	o = &dbm.AuthMethod{
		MethodId: uint(i.MethodID),
		Name: i.Name,
	}

	if i.MethodID == 0 && !update {
		return nil, errors.New("Field [methodId] is required")
	}

	if i.Name == "" && !update {
		return nil, errors.New("Field [name] is required")
	}

	return o, err
}
