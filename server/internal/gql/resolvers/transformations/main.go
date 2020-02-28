package transformations

import (
	"errors"
	gql "github.com/fedoratipper/bitkiosk/server/internal/gql/models"
	dbm "github.com/fedoratipper/bitkiosk/server/internal/orm/models"
	"github.com/fedoratipper/bitkiosk/server/pkg/utils/date"
	"github.com/jinzhu/gorm"
)

// DBUserToGQLUser transforms [user] db input to gql type
func DBUserToGQLUser(i *dbm.User) (o *gql.User, err error) {
	o = &gql.User{
		Email:     i.Email,
		Role:	   int(i.Role),
		CreatedAt: date.FormatToSqlDate(&i.CreatedAt),
		UpdatedAt: date.FormatToSqlDate(i.UpdatedAt),
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

func DBUserProfileToGQLUserProfile(up *dbm.UserProfile, u *dbm.User) (o *gql.UserProfile, err error) {

	if u.ID == 0 {
		return nil, errors.New("unable to find user")
	}

	if up.ID == 0 {
		return nil, errors.New("unable to find user profile")
	}

	o = &gql.UserProfile{
		FirstName:   up.FirstName,
		LastName:    up.LastName,
		Email: 		 &u.Email,
	}

	return o, err
}


func UpdatedDBUserProfileToGQLUserProfile(up *dbm.UserProfile, email string) (o *gql.UserProfile, err error) {

	if up.ID == 0 {
		return nil, errors.New("unable to find user profile")
	}

	o = &gql.UserProfile{
		FirstName:   up.FirstName,
		LastName:    up.LastName,
		Email: 		 &email,
	}

	return o, err
}

func GQLUpdateUserProfileToDBUserProfile(i *gql.UpdatedProfile, db *gorm.DB) (*dbm.UserProfile, error) {

	if i.Email == "" {
		return nil, errors.New("unable to find user profile")
	}

	var userProfile dbm.UserProfile

	db = db.Joins("JOIN users on users.id = user_profiles.user_id").Where("users.email like ?", i.Email).Find(&userProfile)

	if userProfile.ID == 0 {
		return nil, errors.New("unable to find user with email " + i.Email)
	}

	userProfile.FirstName = &i.FirstName
	userProfile.LastName = &i.LastName

	return &userProfile, nil
}
